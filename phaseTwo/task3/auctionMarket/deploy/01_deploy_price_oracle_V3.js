const { ethers, upgrades } = require("hardhat");


module.exports = async ({ getNamedAccounts, deployments, network }) => {
    const priceOracleDeployment = await deployments.get('PriceOracle');
    const proxyAddress = priceOracleDeployment.address;
    const PriceOracleV3 = await ethers.getContractFactory("PriceOracleV3");
    const upgraded = await upgrades.upgradeProxy(proxyAddress, PriceOracleV3);
    await upgraded.waitForDeployment();
    console.log("PriceOracle upgraded at:", await upgraded.getAddress());
}
module.exports.tags = ["PriceOracleProxyV3"];