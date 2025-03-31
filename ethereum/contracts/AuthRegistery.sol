// SPDX-License-Identifier: MIT
pragma solidity ^0.7.0;

contract AuthRegistery {
    mapping(address => bool) private authorizedUsers;
    address public owner;

    event UserAuthorized(address indexed user);
    event UserRevoked(address indexed user);

    modifier onlyOwner() {
        require(msg.sender == owner, "Only owner can perform this action");
        _;
    }

    constructor() {
        owner = msg.sender;
    }

    function authorizeUser(address user) external onlyOwner {
        authorizedUsers[user] = true;
        emit UserAuthorized(user);
    }

    function revokeUser(address user) external onlyOwner {
        authorizedUsers[user] = false;
        emit UserRevoked(user);
    }

    function isAuthorized(address user) external view returns (bool) {
        return authorizedUsers[user];
    }
}
