pragma solidity ^0.5.11;

import "./File.sol";

contract FileStore {

    File[] files;

    function getAllFiles() public view returns (File[] memory) {
        return files;
    }
    function addFile() public {
        File _file = new File(msg.sender);
        files.push(_file);
    }

}
