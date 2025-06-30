const { ethers, upgrades } = require("hardhat");


module.exports = async ({ getNamedAccounts, deployments, network }) => {
    const priceOracleDeployment = await deployments.get('PriceOracle');
    const proxyAddress = priceOracleDeployment.address;
    const PriceOracleV2 = await ethers.getContractFactory("PriceOracleV2");
    const upgraded = await upgrades.upgradeProxy(proxyAddress, PriceOracleV2);
    await upgraded.waitForDeployment();
    console.log("PriceOracle upgraded at:", await upgraded.getAddress());
}
module.exports.tags = ["PriceOracleProxyV2"];