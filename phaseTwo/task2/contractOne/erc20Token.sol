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
4. 部署到sepolia 测试网，导入到自己的钱包
 */

//IERC20.sol 这是一个接口 ，作业估计就是实现这个接口下的方法

//import {ERC20} from "@openzeppelin/contracts/token/ERC20/ERC20.sol";

contract GLDToken {
    mapping(address account => uint256) private _balances; //余额
    //嵌套映射存储授权额度（A允许B使用的代币数）
    mapping(address => mapping(address => uint256)) private _allowances;

    uint256 private _totalSupply; //总供给代币数量
    string private _name; //代币名称
    string private _symbol; //代币符号
    uint8 public decimals = 18; //展示小数点后多少位
    address public owner; //合约拥有者
    event Approval(
        address indexed owner,
        address indexed spender,
        uint256 value
    ); //  授权事件   事件需要emit关键字触发
    event Transfer(address indexed from, address indexed to, uint256 value); //转账事件    事件需要emit关键字触发

    constructor(
        string memory name_,
        string memory symbol_,
        uint256 initialSupply
    ) {
        _name = name_; //名称 例：Gold
        _symbol = symbol_; //符号 例：GLD
        _totalSupply = initialSupply; //初始化代币数量
        owner = msg.sender;
        _balances[owner] = initialSupply; //把代币都放在合约拥有者身上
    }

    //查询账户余额
    function balanceOf(address account) public view returns (uint256) {
        return _balances[account];
    }

    //转账 当前合约调用者转给某个人 所以只需要2个参数 目标地址跟金额 返回成功与否
    function transfer(address to, uint256 amount) public returns (bool) {
        //先看看账户上有没有这么多钱
        require(_balances[msg.sender] >= amount, "Insufficient balance");
        //开转 先把自己的扣了
        _balances[msg.sender] -= amount;
        //给to加上
        _balances[to] += amount;
        //emit 记录转账操作
        emit Transfer(msg.sender, to, amount);
        return true;
    }

    //授权 当前合约调用者授权给某个账户多少代币 所以也是两个参数 返回成功与否
    function approve(address spender, uint256 amount) public returns (bool) {
        //先看看有没有这么多钱给他授权
        require(_balances[msg.sender] >= amount, "Insufficient balance");
        //mapping(address => mapping(address => uint256)) private _allowances;
        _allowances[msg.sender][spender] = amount;
        emit Approval(msg.sender, spender, amount);
        return true;
    }

    //代扣转账 参数 实际扣钱的那个账户跟金额
    function transferFrom(
        address from,
        address to,
        uint256 amount
    ) public returns (bool) {
        //先看看账户上有没有这么多钱给他扣
        require(_balances[from] >= amount, "Insufficient balance");
        //再看看是不是给他授权了这么多钱
        require(
            _allowances[from][msg.sender] >= amount,
            "Insufficient balance"
        );
        //都满足了  再扣钱 先扣授权账户里面的金额
        _allowances[from][msg.sender] -= amount;
        //再扣实际账户的
        _balances[from] -= amount;
        //再给转账的人的钱加起来
        _balances[to] += amount;
        //记录转账日志
        emit Transfer(from, to, amount);
        return true;
    }

    //增发代币
    function mint(uint256 amount) public {
        require(msg.sender == owner, "Only owner can mint");
        _totalSupply += amount; //总供给代币数量增加
        _balances[owner] += amount; //合约拥有者代币数量增加
        emit Transfer(address(0), owner, amount);
    }

    //建议实现 allowance 查询函数  看看授权了多少钱
    function allowance(
        address _owner,
        address spender
    ) public view returns (uint256) {
        return _allowances[_owner][spender];
    }
}
