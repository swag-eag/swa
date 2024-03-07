from pathlib import Path

import pytest

from .network import setup_custom_swa
from .utils import ADDRS, KEYS, send_transaction, w3_wait_for_block, wait_for_new_blocks

pytestmark = pytest.mark.gas


@pytest.fixture(scope="module")
def custom_swa_eq(tmp_path_factory):
    path = tmp_path_factory.mktemp("min-gas-price-eq")
    yield from setup_custom_swa(
        path, 26500, Path(__file__).parent / "configs/min_gas_price_eq.jsonnet"
    )


@pytest.fixture(scope="module")
def custom_swa(tmp_path_factory):
    path = tmp_path_factory.mktemp("min-gas-price")
    yield from setup_custom_swa(
        path, 26530, Path(__file__).parent / "configs/min_gas_price.jsonnet"
    )


@pytest.fixture(scope="module")
def custom_swa_lte(tmp_path_factory):
    path = tmp_path_factory.mktemp("min-gas-price-lte")
    yield from setup_custom_swa(
        path, 26560, Path(__file__).parent / "configs/min_gas_price_lte.jsonnet"
    )


@pytest.fixture(
    scope="module",
    params=["custom_swa_eq", "custom_swa", "custom_swa_lte"],
)
def custom_cluster(request, custom_swa_eq, custom_swa_lte, custom_swa):
    if request.param == "custom_swa_eq":
        return custom_swa_eq
    elif request.param == "custom_swa_lte":
        return custom_swa_lte
    return custom_swa


def adjust_base_fee(parent_fee, gas_limit, gas_used, params):
    "spec: https://eips.ethereum.org/EIPS/eip-1559#specification"
    change_denominator = params["base_fee_change_denominator"]
    elasticity_multiplier = params["elasticity_multiplier"]
    gas_target = gas_limit // elasticity_multiplier
    if gas_used == gas_target:
        return parent_fee
    delta = parent_fee * abs(gas_target - gas_used) // gas_target // change_denominator
    # https://github.com/crypto-org-chain/ethermint/blob/develop/x/feemarket/keeper/eip1559.go#L104
    if gas_target > gas_used:
        return max(parent_fee - delta, params["min_gas_price"])
    else:
        return parent_fee + max(delta, 1)


def get_params(cli):
    params = cli.query_params("feemarket")["params"]
    return {k: int(float(v)) for k, v in params.items()}


def test_dynamic_fee_tx(custom_cluster):
    wait_for_new_blocks(custom_cluster.cosmos_cli(), 1)
    w3 = custom_cluster.w3
    amount = 10000
    before = w3.eth.get_balance(ADDRS["community"])
    tip_price = 1
    max_price = 10000000000000 + tip_price
    tx = {
        "to": "0x0000000000000000000000000000000000000000",
        "value": amount,
        "gas": 21000,
        "maxFeePerGas": max_price,
        "maxPriorityFeePerGas": tip_price,
    }
    txreceipt = send_transaction(w3, tx, KEYS["community"])
    assert txreceipt.status == 1
    blk = w3.eth.get_block(txreceipt.blockNumber)
    assert txreceipt.effectiveGasPrice == blk.baseFeePerGas + tip_price

    fee_expected = txreceipt.gasUsed * txreceipt.effectiveGasPrice
    after = w3.eth.get_balance(ADDRS["community"])
    fee_deducted = before - after - amount
    assert fee_deducted == fee_expected

    assert blk.gasUsed == txreceipt.gasUsed  # we are the only tx in the block

    # check the next block's base fee is adjusted accordingly
    w3_wait_for_block(w3, txreceipt.blockNumber + 1)
    fee = w3.eth.get_block(txreceipt.blockNumber + 1).baseFeePerGas
    params = get_params(custom_cluster.cosmos_cli())
    assert fee == adjust_base_fee(
        blk.baseFeePerGas, blk.gasLimit, blk.gasUsed, params
    ), fee


def test_base_fee_adjustment(custom_cluster):
    """
    verify base fee adjustment of three continuous empty blocks
    """
    wait_for_new_blocks(custom_cluster.cosmos_cli(), 1)
    w3 = custom_cluster.w3
    begin = w3.eth.block_number
    w3_wait_for_block(w3, begin + 3)

    blk = w3.eth.get_block(begin)
    parent_fee = blk.baseFeePerGas
    params = get_params(custom_cluster.cosmos_cli())

    for i in range(3):
        fee = w3.eth.get_block(begin + 1 + i).baseFeePerGas
        assert fee == adjust_base_fee(parent_fee, blk.gasLimit, 0, params)
        parent_fee = fee

    call = w3.provider.make_request
    res = call("eth_feeHistory", [2, "latest", []])["result"]["baseFeePerGas"]
    # nextBaseFee should align max with minGasPrice in eth_feeHistory
    assert all(fee == hex(10000000000000) for fee in res), res
