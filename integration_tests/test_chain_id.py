import json

import pytest
from pystarport import ports

from .utils import wait_for_block, wait_for_port


def test_config_client_id(swa):
    n0 = "swa_777-1-node0"
    p0 = swa.base_port(0)
    w3 = swa.w3
    swa.supervisorctl("stop", n0)
    cli = swa.cosmos_cli(0)
    dir = cli.data_dir / "config"

    def assert_chain_id(chain_id, timeout=None):
        genesis_cfg = dir / "genesis.json"
        genesis = json.loads(genesis_cfg.read_text())
        genesis["chain_id"] = f"swa_{chain_id}-1"
        genesis_cfg.write_text(json.dumps(genesis))
        swa.supervisorctl("start", n0)
        wait_for_port(ports.evmrpc_port(p0))
        assert w3.eth.chain_id == chain_id
        height = w3.eth.get_block_number() + 2
        # check CONSENSUS FAILURE
        if timeout is None:
            wait_for_block(cli, height)
        else:
            with pytest.raises(TimeoutError):
                wait_for_block(cli, height, timeout)

    assert_chain_id(776, 5)
    swa.supervisorctl("stop", n0)
    assert_chain_id(777)
