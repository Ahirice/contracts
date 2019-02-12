pragma solidity ^0.4.25;

import "./externals/SafeMath.sol";
import "./externals/ERC20.sol";
import "./internals/ownable.sol";
import "./internals/claimable.sol";

/// @title ILicence interface describes methods for loading a TokenCard and updating licence amount.
interface ILicence {
    function load(address, uint) external payable;
    function updateLicenceAmount(uint) external;
}

/// @title Licence loads the TokenCard and transfers the licence amout to the token holder contract.
contract Licence is Claimable, Ownable {

    using SafeMath for uint256;

    /*******************/
    /*     Events     */
    /*****************/

    event Received(address _from, uint _amount);

    event UpdatedLicenceDAO(address _sender, address _newDAO);
    event UpdatedCryptoFloat(address _sender, address _newFloat);
    event UpdatedTokenHolder(address _sender, address _newHolder);
    event UpdatedTKNContractAddress(address sender,address _newTKN);
    event UpdatedLicenceAmount(address _sender, uint _newAmount);

    event TransferredToTokenHolder(address _from, address _to, address _asset, uint _amount);
    event TransferredToCryptoFloat(address _from, address _to, address _asset, uint _amount);

    event LogTokenTransfer(address _asset, address _to, uint _amount);

    /// @dev This is 100% scaled up by a factor of 10 to give us an extra 1 decimal place of precision
    uint constant public MAX_AMOUNT_SCALE = 1000;

    address private _TKNContractAddress = 0xaAAf91D9b90dF800Df4F55c205fd6989c977E73a; // solium-disable-line uppercase

    address private _cryptoFloat;
    address private _tokenHolder;
    address private _licenceDAO;

    bool private _lockedCryptoFloat;
    bool private _lockedTokenHolder;
    bool private _lockedLicenceDAO;
    bool private _lockedTKNContractAddress;

    /// @dev This is the _licenceAmountScaled by a factor of 10
    /// i.e. 1% is 10 _licenceAmountScaled, 0.1% is 1 _licenceAmountScaled
    uint private _licenceAmountScaled;

    /// @dev Reverts if called by any address other than the DAO contract.
    modifier onlyDAO() {
        require(msg.sender == _licenceDAO);
        _;
    }

    /// @dev Constructor initializes the card licence contract.
    /// @param _owner is the owner account of the wallet contract.
    /// @param _transferable indicates whether the contract ownership can be transferred.
    /// @param _licence is the initial card licence amount. this number is scaled 10 = 1%, 9 = 0.9%
    /// @param _float is the address of the multi-sig cryptocurrency float contract.
    /// @param _holder is the address of the token holder contract
    constructor(address _owner, bool _transferable, uint _licence, address _float, address _holder, address _tknAddress) Ownable(_owner, _transferable) public {
        _licenceAmountScaled = _licence;
        _cryptoFloat = _float;
        _tokenHolder = _holder;
        if (_tknAddress != address(0)){
          _TKNContractAddress = _tknAddress;
        }
    }

    /// @dev Ether can be deposited from any source, so this contract should be payable by anyone.
    function() public payable {
        require(msg.data.length == 0);
        emit Received(msg.sender, msg.value);
    }

    /// @return the scaled licence amount, used to calculate the split when loading.
    function licenceAmountScaled() external view returns (uint) {
        return _licenceAmountScaled;
    }

    /// @return the address of the multi-sig cryptocurrency float contract.
    function cryptoFloat() external view returns (address) {
        return _cryptoFloat;
    }

    /// @return the address of the token holder contract.
    function tokenHolder() external view returns (address) {
        return _tokenHolder;
    }

    /// @return the address of the DAO contract.
    function licenceDAO() external view returns (address) {
        return _licenceDAO;
    }

    /// @return the address of the TKN contract.
    function TKNContractAddress() external view returns (address) {
        return _TKNContractAddress;
    }

    /// @dev This locks the cryptoFloat addres
    function lockFloat() external onlyOwner {
        _lockedCryptoFloat = true;
    }

    /// @dev This locks the TokenHolder address
    function lockHolder() external onlyOwner {
        _lockedTokenHolder = true;
    }

    /// @dev This locks the DAO address
    function lockLicenceDAO() external onlyOwner {
        _lockedLicenceDAO = true;
    }

    /// @dev This locks the TKN address
    function lockTKNContractAddress() external onlyOwner {
        _lockedTKNContractAddress = true;
    }

    /// @dev returns whether or not the CryptoFloat address is locked
    function floatLocked() public view returns (bool) {
        return _lockedCryptoFloat;
    }

    /// @dev returns whether or not the TokenHolder address is locked
    function holderLocked() public view returns (bool) {
        return _lockedTokenHolder;
    }

    /// @dev returns whether or not the Licence DAO address is locked
    function licenceDAOLocked() public view returns (bool) {
        return _lockedLicenceDAO;
    }

    /// @dev returns whether or not the TKN address is locked
    function TKNContractAddressLocked() public view returns (bool) {
        return _lockedTKNContractAddress;
    }

    /// @dev Updates the address of the cyptoFloat.
    /// @param _newFloat This is the new address for the CryptoFloat
    function updateFloat(address _newFloat) external onlyOwner {
        require(!floatLocked(), "float is locked");
        _cryptoFloat = _newFloat;
        emit UpdatedCryptoFloat(msg.sender, _newFloat);
    }

    /// @dev Updates the address of the Holder contract.
    /// @param _newHolder This is the new address for the TokenHolder
    function updateHolder(address _newHolder) external onlyOwner {
        require(!holderLocked(), "holder contract is locked");
        _tokenHolder = _newHolder;
        emit UpdatedTokenHolder(msg.sender, _newHolder);
    }

    /// @dev Updates the address of the DAO contract.
    /// @param _newDAO This is the new address for the Licence DAO
    function updateLicenceDAO(address _newDAO) external onlyOwner {
        require(!licenceDAOLocked(), "DAO is locked");
        _licenceDAO = _newDAO;
        emit UpdatedLicenceDAO(msg.sender, _newDAO);
    }

    /// @dev Updates the address of the TKN contract.
    /// @param _newTKN This is the new address for the TKN contract
    function updateTKNContractAddress(address _newTKN) external onlyOwner {
        require(!TKNContractAddressLocked(), "TKN is locked");
        _TKNContractAddress = _newTKN;
        emit UpdatedTKNContractAddress(msg.sender, _newTKN);
    }

    /// @dev Updates the TKN licence amount
    /// @param _newAmount is a number between 1 and MAX_AMOUNT_SCALE
    function updateLicenceAmount(uint _newAmount) external onlyDAO {
        require(1 <= _newAmount && _newAmount <= MAX_AMOUNT_SCALE, "licence amount out of range");
        _licenceAmountScaled = _newAmount;
        emit UpdatedLicenceAmount(msg.sender, _newAmount);
    }

    /// @dev Load the holder and float contracts based on the licence amount and asset amount.
    /// @param _asset is the address of an ERC20 token or 0x0 for ether.
    /// @param _amount is the amount of assets to be transferred including the licence amount.
    // TODO(mythril) Is this an issue?: Warning: This contract executes a message call to an address provided as a function argument. Generally, it is not recommended to call user-supplied addresses using Solidity's call() construct. Note that attackers might leverage reentrancy attacks to exploit race conditions or manipulate this contract's state.
    function load(address _asset, uint _amount) external payable {
        uint loadAmount = _amount;
        // If TKN then no licence to be paid
        if (_asset == _TKNContractAddress) {
            require(ERC20(_asset).transferFrom(msg.sender, _cryptoFloat, loadAmount), "TKN transfer from external account was unsuccessful");
        } else {
            loadAmount = _amount.mul(MAX_AMOUNT_SCALE).div(_licenceAmountScaled + MAX_AMOUNT_SCALE);
            uint licenceAmount = _amount.sub(loadAmount);

            if (_asset != address(0)) {
                require(ERC20(_asset).transferFrom(msg.sender, _tokenHolder, licenceAmount), "ERC20 licenceAmount transfer from external account was unsuccessful");
                require(ERC20(_asset).transferFrom(msg.sender, _cryptoFloat, loadAmount), "ERC20 token transfer from external account was unsuccessful");
            } else {
                require(msg.value == _amount, "ETH sent is not equal to amount");
                _tokenHolder.transfer(licenceAmount);
                _cryptoFloat.transfer(loadAmount);
            }

            emit TransferredToTokenHolder(msg.sender, _tokenHolder, _asset, licenceAmount);
        }

        emit TransferredToCryptoFloat(msg.sender, _cryptoFloat, _asset, loadAmount);
    }

    //// @dev Withdraw tokens from the smart contract to the specified account.
    function claim(address _to, address _asset, uint _amount) external onlyOwner {
        _claim(_to, _asset, _amount);
    }
}
