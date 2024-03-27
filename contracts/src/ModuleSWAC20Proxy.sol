pragma solidity ^0.6.8;

import "ds-math/math.sol";
import "./ModuleSWAC20.sol";

contract ModuleSWAC20Proxy is DSMath {
    // sha256('swa-evm')[:20]
    address constant module_address = 0x89A7EF2F08B1c018D5Cc88836249b84Dd5392905;
    ModuleSWAC20 swac20Contract;
    bool isSource;

    event __SwaSendToIbc(address indexed sender, uint256 indexed channel_id, string recipient, uint256 amount, bytes extraData);
    event __SwaSendToEvmChain(address indexed sender, address indexed recipient, uint256 indexed chain_id, uint256 amount, uint256 bridge_fee, bytes extraData);
    event __SwaCancelSendToEvmChain(address indexed sender, uint256 id);

    /**
        Instantiate a ModuleSWAC20Proxy contract. Need to set manually the swac20 contract authority to be the proxy
        like the following call:
        swac20Contract.setAuthority(DSAuthority(address(new ModuleSWAC20ProxyAuthority(address(this)))));
    **/
    constructor(address swac20Contract_, bool isSource_) public {
        swac20Contract = ModuleSWAC20(swac20Contract_);
        isSource = isSource_;
    }

    /**
        views
    **/
    function swac20() public view returns (address) {
        return address(swac20Contract);
    }

    function is_source() public view returns (bool) {
        return isSource;
    }


    /**
        Internal functions to be called by swa module.
    **/
    function mint_by_swa_module(address addr, uint amount) public {
        require(msg.sender == module_address);
        swac20Contract.mint(addr, amount);
    }

    function burn_by_swa_module(address addr, uint amount) public {
        require(msg.sender == module_address);
        swac20_burn(addr, amount);
    }

    function transfer_by_swa_module(address addr, uint amount) public {
        require(msg.sender == module_address);
        swac20Contract.move(addr, module_address, amount);
    }

    function transfer_from_swa_module(address addr, uint amount) public {
        require(msg.sender == module_address);
        swac20Contract.move(address(this), addr, amount);
    }


    /**
        Evm hooks functions
    **/

    // send to another chain through gravity bridge, require approval for the burn.
    function send_to_evm_chain(address recipient, uint amount, uint chain_id, uint bridge_fee, bytes calldata extraData) external {
        // transfer back the token to the proxy account
        if (isSource) {
            swac20Contract.move(msg.sender, address(this), add(amount, bridge_fee));
        } else {
            swac20_burn(msg.sender, add(amount, bridge_fee));
        }
        emit __SwaSendToEvmChain(msg.sender, recipient, chain_id, amount, bridge_fee, extraData);
    }

    // send an "amount" of the contract token to recipient through IBC
    function send_to_ibc(string memory recipient, uint amount, uint channel_id, bytes memory extraData) public {
        if (isSource) {
            swac20Contract.move(msg.sender, address(this), amount);
        } else {
            swac20_burn(msg.sender, amount);
        }
        emit __SwaSendToIbc(msg.sender, channel_id, recipient, amount, extraData);
    }

    // cancel a send to chain transaction considering if it hasnt been batched yet.
    function cancel_send_to_evm_chain(uint256 id) external {
        emit __SwaCancelSendToEvmChain(msg.sender, id);
    }

    /**
        Internal functions
    **/

    // burn the token on behalf of the user. requires approval
    function swac20_burn(address addr, uint amount) internal {
        swac20Contract.move(addr, address(this), amount);
        swac20Contract.burn(amount);
    }
}