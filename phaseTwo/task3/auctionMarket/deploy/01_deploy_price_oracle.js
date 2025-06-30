//用js的代码 表达一个逻辑 对这个逻辑  进行一个module.exports  hardhat 执行 导出出来的这个func
//导出一个异步函数
//部署喂价合约

//升级合约 需要导入
const { upgrades, ethers } = require("hardhat");

// deploy/00_deploy_my_contract.jss
module.exports = async ({ getNamedAccounts, deployments, network }) => {
    const { firstAccount } = await getNamedAccounts();
    // const { deploy, log } = deployments;
    //就不需要deploy 了  用的是upgrades
    const { log } = deployments;


    log("The 01_deploy_price_oracle executing...")
    // 读取 hardhat.config.js 中的自定义参数
    const feeds = network.config.priceFeeds;

    if (!feeds) throw new Error("Missing priceFeeds config for this network!");
    // await deploy('PriceOracle', {
    //     from: firstAccount,
    //     args: [feeds.ethUsd, feeds.usdcUsd, feeds.daiUsd],
    //     log: true,
    // });

    //可升级合约用以下方式部署
    const PriceOracle = await ethers.getContractFactory("PriceOracle");
    const oracle = await upgrades.deployProxy(
        PriceOracle,
        [feeds.ethUsd, feeds.usdcUsd, feeds.daiUsd],
        { initializer: "initialize", from: firstAccount }
    );
    await oracle.waitForDeployment();
    log("The 01_deploy_price_oracle successfully!")
    log("PriceOracle proxy deployed to:", await oracle.getAddress());
    await deployments.save("PriceOracle", {
        address: await oracle.getAddress(),
        abi: (await deployments.getArtifact("PriceOracle")).abi,
    });

};
module.exports.tags = ['sepolia', '12332112345678'];


