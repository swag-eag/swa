from pathlib import Path

import pytest

from .network import setup_custom_swa
from .utils import submit_any_proposal

pytestmark = pytest.mark.gov


@pytest.fixture(scope="module")
def custom_swa(tmp_path_factory):
    path = tmp_path_factory.mktemp("swa")
    yield from setup_custom_swa(
        path, 26400, Path(__file__).parent / "configs/broadcast.jsonnet"
    )


def test_submit_any_proposal(custom_swa, tmp_path):
    submit_any_proposal(custom_swa, tmp_path)
