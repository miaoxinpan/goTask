// SPDX-License-Identifier: MIT
pragma solidity ^0.8;

import "@openzeppelin/contracts/token/ERC721/IERC721.sol";

contract Auction {
    address public seller;//出售人地址
    address public nftAddress;//nft物品地址
    uint256 public tokenId;//nft里面维护了一个tokenId，这个就是那个id
    uint256 public startPrice;//初始价格
    uint256 public endTime;//结束时间  用的是时间戳
    address public highestBidder;//最高出价人的地址
    uint256 public highestBid;//最高出价价格
    bool public ended;//是否结束
    //提现名单
    mapping(address => uint256) public pendingReturns;

    //记录日志
    //创建拍卖
    event AuctionStarted(address indexed seller, address indexed nft, uint256 tokenId, uint256 startPrice, uint256 endTime);
    //出价日志
    event Bid(address indexed bidder, uint256 amount);
    //提现日志
    event Withdraw(address indexed bidder, uint256 amount);
    //拍卖结束
    event AuctionEnded(address winner, uint256 amount);

    constructor(
        address _seller,
        address _nftAddress,
        uint256 _tokenId,
        uint256 _startPrice,
        uint256 _duration//持续时间 时间戳
    ) {
        seller=_seller;
        nftAddress=_nftAddress;
        tokenId=_tokenId;
        startPrice=_startPrice;
        endTime=block.timestamp + _duration;
        emit AuctionStarted(seller, nftAddress, tokenId, startPrice, endTime);
    }
    //拍卖出价
    function bid() external payable {
        require(block.timestamp < endTime, "Auction ended");
        //如果出价比当前最高价，并且出价比开始价高才进行下一步
        require(msg.value > highestBid && msg.value >= startPrice, "Bid too low");\
        
        //把最开始的那个0排除了  
        if (highestBid != 0) {
            // 把前一个价高者加到退钱名单里面
            pendingReturns[highestBidder] += highestBid;
        }
        highestBidder=msg.sender;
        highestBid=msg.value;
        emit Bid(msg.sender, msg.value);
    }
    /**
     * 提现方法
     * 1. 防止重入攻击（Reentrancy Attack） 让用户自己主动提取余额，合约只在用户调用时转账，极大降低重入风险
     * 2. 提高安全性和可控性    用户只能提取属于自己的余额，合约不会主动给用户转账，避免意外或恶意操作
     * 3. 提升用户体验  用户可以随时提现，不受拍卖流程影响
     */
    function withdraw() external {
        uint256 amount = pendingReturns[msg.sender];
        require(amount > 0, "No funds");
        pendingReturns[msg.sender] = 0;
        payable(msg.sender).transfer(amount);
        emit Withdraw(msg.sender,amount);
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
            // 转移NFT给最高出价者
            //调用 这个 nftAddress NFT合约的 safeTransferFrom 方法
            //IERC721 是 ERC721 标准接口，告诉编译器这个地址上有 safeTransferFrom 这个方法
            //这会在链上执行 nftAddress 合约的 safeTransferFrom(seller, highestBidder, tokenId)，完成 NFT 的所有权转移
            IERC721(nftAddress).safeTransferFrom(seller,highestBidder,tokenId);
            // 转账ETH给卖家
            //Solidity 语言内置的ETH 转账语法
            //payable(seller)：把 seller 地址转为可接收 ETH 的 payable 地址类型
            //.transfer(highestBid)：向该地址发送 highestBid 数量的 ETH。
            payable(seller).transfer(highestBid);
        } else {
            // 没人出价，NFT归还卖家
            IERC721(nftAddress).safeTransferFrom(address(this), seller, tokenId);
        }
        //记录结束拍卖结束日志
        emit AuctionEnded(highestBidder, highestBid);
    }
}