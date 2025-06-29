// SPDX-License-Identifier: MIT
pragma solidity ^0.8;
import {ERC721} from "@openzeppelin/contracts/token/ERC721/ERC721.sol";
import {ERC721Burnable} from "@openzeppelin/contracts/token/ERC721/extensions/ERC721Burnable.sol";
import {ERC721Pausable} from "@openzeppelin/contracts/token/ERC721/extensions/ERC721Pausable.sol";
import {ERC721URIStorage} from "@openzeppelin/contracts/token/ERC721/extensions/ERC721URIStorage.sol";
import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";

contract Naruto is ERC721, ERC721URIStorage, ERC721Pausable, Ownable, ERC721Burnable {
    uint256 public _nextTokenId;
    //通常部署这个合约的人 就把他认为是管理员了  我们这边假装合约是火影忍者系列的  所以就默认初始化好了
    //如果是多种多样的nft  什么神奇宝贝杰尼龟，铁甲小宝，那名字跟符号就开放给他们自己输入  那么mint的方法也需要改一下
    //还得增加名字跟符号的mapping等才能实现   比如还是继承这个erc721  只不过名字跟符号  定义的范围更大点 
    //比如说  名字是动漫  Anime 符号是 A 然后里面的mapping  存的Naruto , NA 
    //`initialOwner` → 合约 owner（控制合约权限）
    // `safeMint(address to, ...)` 的 `to` → NFT token 的 owner
    constructor(address initialOwner)
        ERC721("Naruto", "NA")
        Ownable(initialOwner)
    {}
    //合约拥有者调用，暂停 NFT 合约（所有转账、铸造等操作会被暂停）
    function pause() public onlyOwner {
        _pause();
    }
    //合约拥有者调用，恢复合约（解除暂停状态）。
    function unpause() public onlyOwner {
        _unpause();
    }
    //合约拥有者调用，给指定地址 to 铸造一个新的 NFT，cid 是 IPFS 的内容哈希，自动拼接成 ipfs:// 链接作为 tokenURI。
    function safeMint(address to, string memory cid)
        public
        onlyOwner
        returns (uint256)
    {
        string memory uri = string(abi.encodePacked("ipfs://", cid));
        uint256 tokenId = _nextTokenId++;
        _safeMint(to, tokenId);
        _setTokenURI(tokenId, uri);
        return tokenId;
    }

    // The following functions are overrides required by Solidity.
    //内部重写函数，配合 ERC721Pausable 实现暂停时禁止转账。一般用户不会直接调用。
    function _update(address to, uint256 tokenId, address auth)
        internal
        override(ERC721, ERC721Pausable)
        returns (address)
    {
        return super._update(to, tokenId, auth);
    }
    //返回指定 tokenId 的元数据 URI（通常是 NFT 的图片/描述等 json 文件的链接）。
    function tokenURI(uint256 tokenId)
        public
        view
        override(ERC721, ERC721URIStorage)
        returns (string memory)
    {
        return super.tokenURI(tokenId);
    }
    //返回合约是否支持某个接口（如 ERC721、ERC165 等），用于合约兼容性检测。
    //NFT 合约自身实现 supportsInterface 是为了让外部（如 OpenSea、钱包、其他合约）能检测它是否兼容 ERC721、ERC165 等标准。
    //比如说  有一个其他的NFT合约想要调用我这个合约的safeTransferFrom时
    //他就会先通过 supportsInterface 这个接口来检查
    //这个合约是否支持 IERC721Receiver 接口（即 onERC721Received 方法），
    //以确保 NFT 不会被转到一个无法接收 NFT 的合约里
    //如果不支持，则交易会 revert，防止 NFT 被锁死。
    function supportsInterface(bytes4 interfaceId)
        public
        view
        override(ERC721, ERC721URIStorage)
        returns (bool)
    {
        return super.supportsInterface(interfaceId);
    }
}


/**
 *  1.先写 NFT 合约，确保能 mint。
    2.写拍卖合约，先实现 ETH 拍卖流程，后续加 ERC20、预言机。
    3.写工厂合约，支持批量部署和管理拍卖合约。
    4.集成 Chainlink 预言机，完成价格换算。
    5.用 OpenZeppelin 插件支持合约升级。
    6.写测试和部署脚本，逐步完善功能和测试覆盖率。
 */