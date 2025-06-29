// SPDX-License-Identifier: MIT
pragma solidity ^0.8;

import "@chainlink/contracts/src/v0.8/shared/interfaces/AggregatorV3Interface.sol";
import "hardhat/console.sol"; // 引入 Hardhat 控制台


contract PriceOracle {
    AggregatorV3Interface public ethUsdPriceFeed;
    AggregatorV3Interface public usdcUsdPriceFeed;
    AggregatorV3Interface public daiUsdPriceFeed;
    //AggregatorV3Interface public usdtUsdPriceFeed;

    /**
     * ETH/USD
     * Sepolia  测试网预言机地址：
     *          0x694AA1769357215DE4FAC081bf1f309aDC325306
     * 主网预言机地址：
     *          0x5f4ec3df9cbd43714fe2740f5e3616155c5b8419
     * ERC20 常用的：USDC、DAI、USDT
     * USDC/USD 地址：0x0A6513e40db6EB1b165753AD52E80663aeA50545
     * DAI/USD 地址：0x0d79df66BE487753B02D015Fb622DED7f0E9798d
     * USDT/USD 地址：0xC16679AAd21eB7D1d45aB1FA80C05e81eA1cD02f
     */

    constructor(address ethUsdFeed,address usdcUsdFeed,address daiUsdFeed) {//,address usdtUsdFeed
        ethUsdPriceFeed = AggregatorV3Interface(ethUsdFeed);
        usdcUsdPriceFeed = AggregatorV3Interface(usdcUsdFeed);
        daiUsdPriceFeed = AggregatorV3Interface(daiUsdFeed);
    }

    // 获取最新价格（如 ETH/USD），返回值带8位小数
    //ETH/USD
    function getLatestEthUsdPrice() public view returns (int256) {
        (, int256 price, , , ) = ethUsdPriceFeed.latestRoundData();
        return price;
    }
    //USDC/USD
    function getLatestUsdcUsdPrice() public view returns (int256) {
        (, int256 price, , , ) = usdcUsdPriceFeed.latestRoundData();
        return price;
    }
    //DAI/USD
    function getLatestDaiUsdPrice() public view returns (int256) {
        (, int256 price, , , ) = daiUsdPriceFeed.latestRoundData();
        return price;
    }
}
