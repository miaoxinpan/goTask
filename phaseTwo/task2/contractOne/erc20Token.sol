// SPDX-License-Identifier: MIT
pragma solidity ^0.8;

/***
 * ### ✅ 作业 1：ERC20 代币
任务：参考 openzeppelin-contracts/contracts/token/ERC20/IERC20.sol实现一个简单的 ERC20 代币合约。要求：
1. 合约包含以下标准 ERC20 功能：
- balanceOf：查询账户余额。
- transfer：转账。
- approve 和 transferFrom：授权和代扣转账。
2. 使用 event 记录转账和授权操作。
3. 提供 mint 函数，允许合约所有者增发代币。
提示：
- 使用 mapping 存储账户余额和授权信息。
- 使用 event 定义 Transfer 和 Approval 事件。
4. 部署到sepolia 测试网，导入到自    己的钱包
 */

//IERC20.sol 这是一个接口 ，作业估计就是实现这个接口下的方法


import {ERC20} from "@openzeppelin/contracts/token/ERC20/ERC20.sol";

contract GLDToken is ERC20 {
    constructor(uint256 initialSupply) ERC20("Gold", "GLD") {
        _mint(msg.sender, initialSupply);
    }
}
