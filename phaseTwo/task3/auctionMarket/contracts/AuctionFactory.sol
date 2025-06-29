// SPDX-License-Identifier: MIT
pragma solidity ^0.8;

import "./auction.sol"; // 你的拍卖合约路径

contract AuctionFactory {
    //拍卖合约地址的数组
    address[] public allAuctions;

    event AuctionCreated(address indexed auction, address indexed creator);

    // 创建新的拍卖合约
    function createAuction(
        address nftAddress,
        uint256 tokenId,
        uint256 startPrice,
        uint256 duration,
        address usdc,
        address dai,
        address priceOracle
    ) external returns (address) {
        // 部署新拍卖合约
        Auction auction = new Auction(
            msg.sender,
            nftAddress,
            tokenId,
            startPrice,
            duration,
            usdc,
            dai,
            priceOracle
        );
        allAuctions.push(address(auction));
        emit AuctionCreated(address(auction), msg.sender);
        return address(auction);
    }

    // 获取所有拍卖合约地址
    function getAllAuctions() external view returns (address[] memory) {
        return allAuctions;
    }

    // 获取指定索引的拍卖合约地址
    function getAuction(uint256 index) external view returns (address) {
        require(index < allAuctions.length, "Index out of range");
        return allAuctions[index];
    }

    // 获取拍卖合约总数
    function getAuctionCount() external view returns (uint256) {
        return allAuctions.length;
    }
    //批量创建拍卖合约
}
