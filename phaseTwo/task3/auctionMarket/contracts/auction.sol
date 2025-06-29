// SPDX-License-Identifier: MIT
pragma solidity ^0.8;

import "@openzeppelin/contracts/token/ERC721/IERC721.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "hardhat/console.sol"; // 引入 Hardhat 控制台


//把喂价的接口引入进来
interface IPriceOracle {
    function getLatestEthUsdPrice() external view returns (int256);
    function getLatestUsdcUsdPrice() external view returns (int256);
    function getLatestDaiUsdPrice() external view returns (int256);
}


contract Auction {
    //预言机合约的地址 不然咋调用价格接口
    address public priceOracle;
    address public eth;//币种地址   
    address public usdc;
    address public dai;



    address public seller;//出售人地址
    address public nftAddress;//nft物品地址
    uint256 public tokenId;//nft里面维护了一个tokenId，这个就是那个id
    uint256 public startPrice;//初始价格
    uint256 public endTime;//结束时间  用的是时间戳
    address public highestBidder;//最高出价人的地址
    address public highestBidToken;    // 最高出价币种  (ETH=address(0))
    uint256 public highestBidAmount;   // 最高出价币种的数量
    uint256 public highestBidUsd;      // 最高出价的美元价值（8位小数）
    bool public ended;//是否结束




    //提现名单
    //mapping(address => uint256) public pendingReturns;
    //因为不仅仅是eth 还有其他的erc20的代币需要退回 所以这里单纯记录地址已经不够了  
    //用嵌套map的方式 重新设计一个数据结构来存储可以体现的名单
    //pendingReturns[用户地址][币种地址] = 金额
    mapping(address userAddr=> mapping(address coinAddr=>uint256 amount)) public pendingReturns;

    //记录日志
    //创建拍卖
    event AuctionStarted(address indexed seller, address indexed nft, uint256 tokenId, uint256 startPrice, uint256 endTime);
    //出价日志
    //event NewBid(address indexed bidder, uint256 amount);
    //现在多币种拍卖，所以记录的话 把这些信息也一起记录
    event NewBid(address indexed bidder,address indexed coinAddr, uint256 amount, uint256 usdValue);
    //提现日志
    event Withdraw(address indexed bidder,address indexed coinAddr, uint256 amount);
    //拍卖结束
    event AuctionEnded(address winner,address indexed coinAddr, uint256 amount);

    constructor(
        address _seller,
        address _nftAddress,
        uint256 _tokenId,
        uint256 _startPrice,
        uint256 _duration,//持续时间 时间戳
        address _usdc,
        address _dai,
        address _priceOracle
    ) {
        seller=_seller;
        nftAddress=_nftAddress;
        tokenId=_tokenId;
        startPrice=_startPrice;
        endTime=block.timestamp + _duration;
        usdc=_usdc;
        dai=_dai;
        priceOracle=_priceOracle;
        eth=address(0);
        emit AuctionStarted(seller, nftAddress, tokenId, startPrice, endTime);
    }
    //拍卖出价
    //改为用美元衡量价值的话 就得传入他所花费的币种地址
    function bid(address coinAddr,uint256  amount) external payable {
        require(block.timestamp < endTime, "Auction ended");
        //如果出价比当前最高价，并且出价比开始价高才进行下一步
        //require(msg.value > highestBid && msg.value >= startPrice, "Bid too low");
        
        //先拿到他实际对应的价格
        uint256 usdValue =getUsdValue(coinAddr, amount);
        //再根据这个价值去比较
        require(usdValue > highestBidUsd , "Bid too low");

        // 实际转账
        if (coinAddr == address(0)) {
            require(msg.value == amount, "ETH amount mismatch");
        } else {
            IERC20(coinAddr).transferFrom(msg.sender, address(this), amount);
        }
        //把最开始的那个0排除了  
        if (highestBidUsd != 0) {
            // 把前一个价高者加到退钱名单里面 还得记录币种地址 和 钱币数量
            pendingReturns[highestBidder][highestBidToken]+= highestBidAmount;
        }
        highestBidder=msg.sender;
        highestBidToken=coinAddr;
        highestBidAmount=amount;
        highestBidUsd=usdValue;
        emit NewBid(msg.sender,coinAddr,amount,usdValue);
    }
    /**
     * 提现方法
     * 1. 防止重入攻击（Reentrancy Attack） 让用户自己主动提取余额，合约只在用户调用时转账，极大降低重入风险
     * 2. 提高安全性和可控性    用户只能提取属于自己的余额，合约不会主动给用户转账，避免意外或恶意操作
     * 3. 提升用户体验  用户可以随时提现，不受拍卖流程影响
     */
    function withdraw(address coinAddr) external {
        uint256 amount = pendingReturns[msg.sender][coinAddr];
        require(amount > 0, "No funds");
        pendingReturns[msg.sender][coinAddr] = 0;
        //分别处理不同币种的提现
        if (coinAddr == address(0)) {
            payable(msg.sender).transfer(amount);
        } else {
            IERC20(coinAddr).transfer(msg.sender, amount);
        }
        emit Withdraw(msg.sender,coinAddr,amount);
    }

    //结束拍卖  转移NFT和资金
    function endAuction() external {
        //判断时间是否到了
        require(block.timestamp>=endTime,"Auction not yet ended");
        //还得再判断结束标志是否结束 防止多次暂停
        require(!ended, "Auction already ended");
        
        //把标志位改成true
        ended = true;

        //地址不是空的
        if (highestBidder != address(0)) {
            console.log("highestBidder",highestBidder);
            // 转移NFT给最高出价者
            //调用 这个 nftAddress NFT合约的 safeTransferFrom 方法
            //IERC721 是 ERC721 标准接口，告诉编译器这个地址上有 safeTransferFrom 这个方法
            //这会在链上执行 nftAddress 合约的 safeTransferFrom(seller, highestBidder, tokenId)，完成 NFT 的所有权转移
            IERC721(nftAddress).safeTransferFrom(seller,highestBidder,tokenId);
            console.log(unicode"什么情况啊  疯狂报错！！！");

            // 转账ETH给卖家
            //Solidity 语言内置的ETH 转账语法
            //payable(seller)：把 seller 地址转为可接收 ETH 的 payable 地址类型
            //.transfer(highestBid)：向该地址发送 highestBid 数量的 ETH。
            //payable(seller).transfer(highestBid);
            //多币种拍卖  这边转账方式修改一下  兼容erc20的常用代币与eth
            if(highestBidToken==address(0)){//eth
                payable(seller).transfer(highestBidAmount);
            }else{
                console.log(unicode"不会吧！！！");
                IERC20(highestBidToken).transfer(seller, highestBidAmount);
            }
        } else {
            // 没人出价，NFT归还卖家
            
            IERC721(nftAddress).safeTransferFrom(address(this), seller, tokenId);
        }
        //记录结束拍卖结束日志
        emit AuctionEnded(highestBidder,highestBidToken, highestBidAmount);
    }

    //根据币种换算成美元价值
    function getUsdValue(address coinAddr, uint256 amount) public view returns (uint256) {
        if (coinAddr == address(0)) {
            int256 price = IPriceOracle(priceOracle).getLatestEthUsdPrice();
            return amount * uint256(price) / 1e18;
        } else if (coinAddr == usdc) {
            int256 price = IPriceOracle(priceOracle).getLatestUsdcUsdPrice();
            return amount * uint256(price) / 1e6;
        } else if (coinAddr == dai) {
            int256 price = IPriceOracle(priceOracle).getLatestDaiUsdPrice();
            return amount * uint256(price) / 1e18;
        }
        revert("Unsupported token");
    }

}