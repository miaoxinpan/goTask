// SPDX-License-Identifier: MIT
pragma solidity ^0.8;

//升级合约依赖
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@chainlink/contracts/src/v0.8/shared/interfaces/AggregatorV3Interface.sol";

import "hardhat/console.sol"; // 引入 Hardhat 控制台


contract PriceOracle is Initializable{
    AggregatorV3Interface public ethUsdPriceFeed;
    AggregatorV3Interface public usdcUsdPriceFeed;
    AggregatorV3Interface public daiUsdPriceFeed;
    //AggregatorV3Interface public usdtUsdPriceFeed;
/**
 * This Contract Similar Matches the deployed Bytecode at 0x1193F56f2dc46BB4b1BBb148f07587250766a5A7 , 
 * additional source code verification (with the current user credential) is temporarily unavailable.
 * 
 * This contract may be a proxy contract. Click on More Options and select Is this a proxy?
 *  to confirm and enable the "Read as Proxy" & "Write as Proxy" tabs.
 */
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

    // 初始化函数（替代构造函数）
    //function initialize(uint256 _initValue) public initializer {
    function initialize(address ethUsdFeed,address usdcUsdFeed,address daiUsdFeed) public initializer {//,address usdtUsdFeed
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
