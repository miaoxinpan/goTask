// SPDX-License-Identifier: MIT
pragma solidity ^0.8;
import "@chainlink/contracts/src/v0.8/interfaces/AggregatorV3Interface.sol";

contract PriceOracle {
    AggregatorV3Interface public priceFeed;

    /**
     * ETH/USD
     * Sepolia  测试网预言机地址：
     *          0x694AA1769357215DE4FAC081bf1f309aDC325306
     * 主网预言机地址：
     *          0x5f4ec3df9cbd43714fe2740f5e3616155c5b8419
     * 
     * ERC20/USD TODO:回家查询一下
     */

    constructor(address feedAddress) {
        priceFeed = AggregatorV3Interface(feedAddress);
    }

    // 获取最新价格（如 ETH/USD），返回值带8位小数
    function getLatestPrice() public view returns (int256) {
        (, int256 price,,,) = priceFeed.latestRoundData();
        return price;
    }
}