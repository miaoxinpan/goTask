
//部署AuctionFactory合约
module.exports = async ({ getNamedAccounts, deployments }) => {
    const { firstAccount } = await getNamedAccounts();
    const { deploy, log } = deployments;
    log("The 04_deploy_auction_factory executing...")
    await deploy('AuctionFactory', {
        from: firstAccount,
        log: true,
    });
    log("The 04_deploy_auction_factory successfully!")
};
module.exports.tags = ['sepolia'];