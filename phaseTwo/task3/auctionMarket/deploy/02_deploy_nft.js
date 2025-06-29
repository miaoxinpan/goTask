
//部署nft合约
module.exports = async ({ getNamedAccounts, deployments }) => {
    const { firstAccount } = await getNamedAccounts();
    const { deploy, log } = deployments;


    log("The 02_deploy_nft executing...")
    // 读取 hardhat.config.js 中的自定义参数
    const feeds = network.config.priceFeeds;

    if (!feeds) throw new Error("Missing priceFeeds config for this network!");
    await deploy('Naruto', {
        contract:"Naruto",
        from: firstAccount,
        args: [firstAccount],
        log: true,
    });
    log("The 02_deploy_nft deployed successfully!")
};
module.exports.tags = ['sepolia'];