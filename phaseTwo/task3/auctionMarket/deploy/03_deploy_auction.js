const { ethers } = require("hardhat");
require('dotenv').config();
//部署拍卖合约
module.exports = async ({ getNamedAccounts, deployments }) => {
    const { firstAccount } = await getNamedAccounts();
    const { deploy, log } = deployments;


    log("The 03_deploy_auction executing...")
    /**
     * address _seller,
        address _nftAddress,
        uint256 _tokenId,
        uint256 _startPrice,
        uint256 _duration,//持续时间 时间戳
        address _usdc,
        address _dai,
        address _usdt,
        address _priceOracle
     * */
    //第一个参数  卖家
    const seller=firstAccount;
    //这个时候我的nft其实还没有mint  所以直接在这里mint一个  搞完了以后再部署这个拍卖合约
    const nftDeployment = await deployments.get('Naruto'); // 'Naruto' 是你的 NFT 合约名
    const nftAddress = nftDeployment.address;//nft的合约地址 因为实际上修改的数据是这个合约内部的tokenId的归属

    const nftInstance = await ethers.getContractAt('Naruto', nftAddress);
    //两个参数 一个to  给谁  一个cid bafkreif3sn52jsphe7kdsqwwm37pbb277m2n36czhqz4lyw6bgrndnp6b4
    const cid = process.env.NFT_CID;

    const nextTokenId = await nftInstance._nextTokenId(); // mint 前读取
    await nftInstance.safeMint(firstAccount, cid);        // mint
    const tokenId = typeof nextTokenId.toNumber === "function"
        ? nextTokenId.toNumber()
        : parseInt(nextTokenId);
    //这边返回的是一个 返回的是一个 TransactionResponse（交易对象） 不是safeMint 返回的一个id
    //const tokenId= await nftInstance.safeMint(firstAccount, cid); 

    const startPrice=100
    const duration = 3600; // 1小时
    

    
    const usdc = process.env.USDC_ADDR;
    const dai = process.env.DAI_ADDR;
    //const usdt = process.env.USDT_ADDR;
    //从hardhat中获取到nft的合约
    
    const priceOracleDeployment = await deployments.get('PriceOracle');
    const priceOracle = priceOracleDeployment.address;
    

    await deploy('Auction', {
        contract:"Auction",
        from: firstAccount,
        args: [seller,nftAddress,tokenId,startPrice,duration,usdc,dai,priceOracle],
        log: true,
    });
    log("The 03_deploy_auction deployed successfully!")
};
module.exports.tags = ['sepolia'];