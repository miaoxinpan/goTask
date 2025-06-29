//用js的代码 表达一个逻辑 对这个逻辑  进行一个module.exports  hardhat 执行 导出出来的这个func
//导出一个异步函数
//部署喂价合约


// deploy/00_deploy_my_contract.jss
module.exports = async ({ getNamedAccounts, deployments, network }) => {
    const { firstAccount } = await getNamedAccounts();
    const { deploy, log } = deployments;


    log("The 01_deploy_price_oracle executing...")
    // 读取 hardhat.config.js 中的自定义参数
    const feeds = network.config.priceFeeds;

    if (!feeds) throw new Error("Missing priceFeeds config for this network!");
    await deploy('PriceOracle', {
        from: firstAccount,
        args: [feeds.ethUsd, feeds.usdcUsd, feeds.daiUsd],
        log: true,
    });
    log("The 01_deploy_price_oracle successfully!")
};
module.exports.tags = ['sepolia'];