pragma solidity ^0.5.11;

contract File {

    address public creator;
    address[] public owners;

    // Power Values
    // 0 - Blocked
    // 1 - Read Only
    // 2 - Read+Write
    // 3 - Read+Write+Add Owner+Change_Power_of_Owner
    // 4 - Creator
    mapping (address => uint) owner_power;
    mapping (address => bool) is_owner_exists;
    string private key;
    string private ipfs_hash;

    modifier restricted_creator() {
        require(msg.sender == creator, "only creator can access");
        _;
    }

    modifier restricted_owner_doesnot_exists(address _owner) {
        require(!is_owner_exists[_owner], "owner already exists");
        _;
    }

    modifier restricted_owner_has_power(uint power) {
        require(is_owner_exists[msg.sender], "owner doesnot exists");
        require(owner_power[msg.sender]>=power, "doesnot have sufficient power");
        _;
    }

    modifier restricted_owner_is_not_creator(address _owner) {
        require(_owner != creator, "cannot perform operation on owner");
        _;
    }

    modifier restricted_owner_exists(address _owner) {
        require(is_owner_exists[_owner], "owner doesnot exists");
        _;
    }

    constructor (address _creator) public {
        creator = _creator;
        add_owner(_creator, 4);
    }

    function add_key(string memory _key) public {
        key = _key;
    }
    
    // Owner
    
    function add_owner(address _owner, uint _power) public
            restricted_owner_doesnot_exists(_owner) restricted_owner_has_power(3) {
        owners.push(creator);
        owner_power[_owner] = _power;
        is_owner_exists[_owner] = true;
    }

    function change_owner_power(address _owner, uint _power) public
            restricted_owner_is_not_creator(_owner) restricted_owner_exists(_owner)
            restricted_owner_has_power(3) {
        owner_power[_owner] = _power;
    }

    // Owner's view

    function get_owner_power(address _owner) public restricted_owner_has_power(1) view returns(uint) {
        return owner_power[_owner];
    }

    function get_owner_list() public restricted_owner_has_power(1) view returns(address[] memory) {
        return owners;
    }




    // keys
    function update_key(string memory _key) public restricted_creator() {
        key = _key;
    }
    function get_key() public restricted_owner_has_power(1) view returns(string memory) {
        return key;
    }

    // Ipfs Hash
    function update_hash(string memory _hash) public restricted_owner_has_power(2) {
        ipfs_hash = _hash;
    }
    function get_ipfs_hash() public restricted_owner_has_power(1) view returns(string memory) {
        return ipfs_hash;
    } 

}
