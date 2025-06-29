const { expect } = require("chai");
const { ethers, network } = require("hardhat");
require('dotenv').config();

//首先得先确定测试 AuctionMarket 的全流程
/**
 * 1.创建nft
 * 2.创建拍卖
 * 3.不同用户参与竞拍
 * 5.结束竞拍
 * 6.没买到的用户提现
 */
//定义好全局变量
let nft, auctionFactory, priceOracle;
let owner, bidder1, bidder2, bidder3, bidder4;
let tokenId, auctionAddr;
let usdc, dai;
let USDC_ADDR = process.env.USDC_ADDR?.trim();
let DAI_ADDR = process.env.DAI_ADDR?.trim();



describe("AuctionMarket 全流程", async function () {
    //给全局变量赋值  让他在下面的测试用例中能够使用
    this.beforeEach(async function () {
        //先拿到3个地址当作  拥有者  竞拍者1    竞拍者2
        [owner, bidder1, bidder2, bidder3, bidder4] = await ethers.getSigners();


        // 部署 Mock ERC20 代币
        const MockERC20 = await ethers.getContractFactory("MockERC20");
        usdc = await MockERC20.deploy("USDC", "USDC", 18);
        dai = await MockERC20.deploy("DAI", "DAI", 18);
        //给每个账户先转钱
        await usdc.transfer(bidder2.address, ethers.parseUnits("10000", 18));
        await usdc.transfer(bidder4.address, ethers.parseUnits("10000", 18));
        await dai.transfer(bidder3.address, ethers.parseUnits("10000", 18));


        //部署到sepolia测试网的时候获取配置文件中priceFeeds
        //feeds = network.config.priceFeeds;
        //部署语言机
        //priceOracle = await (await ethers.getContractFactory("PriceOracle")).deploy(
        //    feeds.ethUsd,
        //   feeds.usdcUsd,
        //    feeds.daiUsd
        //)
        //在本地跑 自己mock的
         // 部署 MockV3Aggregator 预言机
        const MockV3Aggregator = await ethers.getContractFactory("MockV3Aggregator");
        const mockEthFeed = await MockV3Aggregator.deploy(8, ethers.parseUnits("2000", 8)); // 2000 USD
        const mockUsdcFeed = await MockV3Aggregator.deploy(8, ethers.parseUnits("1", 8));   // 1 USD
        const mockDaiFeed = await MockV3Aggregator.deploy(8, ethers.parseUnits("1", 8));    // 1 USD
        // 部署 PriceOracle 用 mock 预言机地址
        priceOracle = await (await ethers.getContractFactory("PriceOracle")).deploy(
            mockEthFeed.target,
            mockUsdcFeed.target,
            mockDaiFeed.target
        );  

        // 赋值合约地址变量
        USDC_ADDR = usdc.target;
        DAI_ADDR = dai.target;

        //再部署nft
        nft = await (await ethers.getContractFactory("Naruto")).deploy(owner.address);
        //mint nft
        const mintTx = await nft.safeMint(owner.address, "testcid");
        //等待6个交易块
        await mintTx.wait();

        //tokenId = (await nft._nextTokenId()) - 1; // 获取刚mint的tokenId
        tokenId = Number(await nft._nextTokenId()) - 1;
        //部署拍卖工厂
        //没有构造方法 所以不用传参数
        const AuctionFactory = await ethers.getContractFactory("AuctionFactory");
        auctionFactory = await AuctionFactory.deploy();
        await auctionFactory.waitForDeployment();
        console.log(`auctionFactory address:,${auctionFactory.target}`);
        
        //授权nft给拍卖合约
        await nft.connect(owner).approve(auctionFactory.target, tokenId);
    })

    //前置工作完成以后  开始做测试
    it("创建拍卖-竞拍-结束-体现", async function () {
        //1.创建拍卖
        //通过auctionFactory 工厂来创建一个auction
        /**
         * address _seller,
            address _nftAddress,
            uint256 _tokenId,
            uint256 _startPrice,
            uint256 _duration,//持续时间 时间戳
            address _usdc,
            address _dai,
            address _priceOracle
         * 
         */
        
        console.log("!!!");
        const tx = await auctionFactory.createAuction(
            nft.target, tokenId, 100, 300, USDC_ADDR, DAI_ADDR, priceOracle.target
        );

        console.log("???");
        /**
         * 等待交易 tx 被打包进区块并返回交易回执（receipt），里面包含了事件、gas等信息。
         */
        const receipt = await tx.wait();

        console.log("receipt====",receipt);

        const iface = auctionFactory.interface;
        const log = receipt.logs.find(
            (l) => l.topics[0] === iface.getEvent("AuctionCreated").topicHash
        );
        //在所有事件中查找名为 "AuctionCreated" 的事件对象
        //const event = receipt.events.find(e => e.event === "AuctionCreated");

        console.log("event?????");
        const event = iface.parseLog(log);
        auctionAddr = event.args.auction;
        //从事件参数中提取出新创建的拍卖合约地址（通常事件里会有 auction 字段）。
        auctionAddr = event.args.auction;//拍卖合约的地址就拿到了


        console.log("auctionAddr:", auctionAddr);


        //拿到地址创建实例
        auction = await ethers.getContractAt("Auction", auctionAddr);
        await nft.connect(owner).approve(auctionAddr, tokenId);

        //在出价前 授权扣钱
        await usdc.connect(bidder2).approve(auctionAddr, ethers.parseUnits("3500", 18));
        await usdc.connect(bidder4).approve(auctionAddr, ethers.parseUnits("3800", 18));
        await dai.connect(bidder3).approve(auctionAddr, ethers.parseUnits("3300", 18));
        //2.参与竞拍(用以太坊的地址  传入1个以太币)

        // console.log("1",ethers.ZeroAddress);

        // await auction.connect(bidder1).bid(ethers.ZeroAddress, ethers.parseEther("1"));
        // expect(await auction.highestBidder()).to.equal(bidder1.address)
        // expect(await auction.highestBidToken()).to.equal(ethers.ZeroAddress)
        // expect(await auction.highestBidAmount()).to.equal(ethers.parseEther("1"))

        console.log("2");
        await auction.connect(bidder2).bid(USDC_ADDR, ethers.parseUnits("3500", 18));
        console.log("2.5");
        expect(await auction.highestBidder()).to.equal(bidder2.address)
        expect(await auction.highestBidToken()).to.equal(USDC_ADDR)
        expect(await auction.highestBidAmount()).to.equal(ethers.parseUnits("3500", 18))

        console.log("3");
        //await auction.connect(bidder3).bid(DAI_ADDR, ethers.parseUnits("3300", 18));

        //expect(await auction.highestBidder()).to.equal(bidder2.address)
        //expect(await auction.highestBidToken()).to.equal(USDC_ADDR)
        //expect(await auction.highestBidAmount()).to.equal(ethers.parseUnits("3500", 18))


        console.log("4");
        await auction.connect(bidder4).bid(USDC_ADDR, ethers.parseUnits("3800", 18));

        expect(await auction.highestBidder()).to.equal(bidder4.address)
        expect(await auction.highestBidToken()).to.equal(USDC_ADDR)
        expect(await auction.highestBidAmount()).to.equal(ethers.parseUnits("3800", 18))
        //3.结束竞拍

        // 4. 快进时间到拍卖结束
        await hre.network.provider.send("evm_increaseTime", [3600]);
        await hre.network.provider.send("evm_mine");
        console.log("4.5");


        const block = await ethers.provider.getBlock("latest");
        console.log("当前区块时间:", block.timestamp);
        console.log("拍卖合约结束时间:", await auction.endTime());
        

        
        //结束竞拍
        await auction.endAuction();
        console.log("5");
        //4.竞拍者提现
        await auction.connect(bidder2).withdraw(USDC_ADDR);
        console.log("6");
        expect(await auction.pendingReturns(bidder2.address, USDC_ADDR)).to.equal(0);
        console.log("7");
        //断言nft的归属
        expect(await nft.ownerOf(tokenId)).to.equal(bidder4.address);

    })
})

