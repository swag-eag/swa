import json
from pathlib import Path

import pytest

from .network import setup_custom_swa
from .utils import ADDRS, CONTRACTS

pytestmark = pytest.mark.slow


@pytest.fixture(scope="module")
def custom_swa(tmp_path_factory):
    path = tmp_path_factory.mktemp("swa")
    yield from setup_custom_swa(
        path, 26000, Path(__file__).parent / "configs/genesis_token_mapping.jsonnet"
    )


def test_exported_contract(custom_swa):
    "demonstrate that contract state can be deployed in genesis"
    w3 = custom_swa.w3
    abi = json.loads(CONTRACTS["TestERC20Utility"].read_text())["abi"]
    erc20 = w3.eth.contract(
        address="0x68542BD12B41F5D51D6282Ec7D91D7d0D78E4503", abi=abi
    )
    assert erc20.caller.balanceOf(ADDRS["validator"]) == 100000000000000000000000000


def test_exported_token_mapping(custom_swa):
    cli = custom_swa.cosmos_cli(0)
    rsp = cli.query_contract_by_denom(
        "gravity0x0000000000000000000000000000000000000000"
    )
    assert rsp["contract"] == "0x68542BD12B41F5D51D6282Ec7D91D7d0D78E4503"
    assert rsp["auto_contract"] == "0x68542BD12B41F5D51D6282Ec7D91D7d0D78E4503"
