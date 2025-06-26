// SPDX-License-Identifier: SEE LICENSE IN LICENSE
pragma solidity ^0.8.0;

/**
 * ### ✅ 作业3：编写一个讨饭合约
任务目标
1. 使用 Solidity 编写一个合约，允许用户向合约地址发送以太币。
2. 记录每个捐赠者的地址和捐赠金额。
3. 允许合约所有者提取所有捐赠的资金。

任务步骤
1. 编写合约
  - 创建一个名为 BeggingContract 的合约。
  - 合约应包含以下功能：
  - 一个 mapping 来记录每个捐赠者的捐赠金额。
  - 一个 donate 函数，允许用户向合约发送以太币，并记录捐赠信息。
  - 一个 withdraw 函数，允许合约所有者提取所有资金。
  - 一个 getDonation 函数，允许查询某个地址的捐赠金额。
  - 使用 payable 修饰符和 address.transfer 实现支付和提款。
2. 部署合约
  - 在 Remix IDE 中编译合约。
  - 部署合约到 Goerli 或 Sepolia 测试网。
3. 测试合约
  - 使用 MetaMask 向合约发送以太币，测试 donate 功能。
  - 调用 withdraw 函数，测试合约所有者是否可以提取资金。
  - 调用 getDonation 函数，查询某个地址的捐赠金额。

任务要求
1. 合约代码：
  - 使用 mapping 记录捐赠者的地址和金额。
  - 使用 payable 修饰符实现 donate 和 withdraw 函数。
  - 使用 onlyOwner 修饰符限制 withdraw 函数只能由合约所有者调用。
2. 测试网部署：
  - 合约必须部署到 Goerli 或 Sepolia 测试网。
3. 功能测试：
  - 确保 donate、withdraw 和 getDonation 函数正常工作。

提交内容
1. 合约代码：提交 Solidity 合约文件（如 BeggingContract.sol）。
2. 合约地址：提交部署到测试网的合约地址。
3. 测试截图：提交在 Remix 或 Etherscan 上测试合约的截图。

额外挑战（可选）
1. 捐赠事件：添加 Donation 事件，记录每次捐赠的地址和金额。
2. 捐赠排行榜：实现一个功能，显示捐赠金额最多的前 3 个地址。
3. 时间限制：添加一个时间限制，只有在特定时间段内才能捐赠。
 */

contract BeggingContract {
    //记录捐赠者的地址和金额
    mapping(address donator => uint256 amount) private _donators;
    address public _owner; //合约拥有者
    //实现一个功能，显示捐赠金额最多的前 3 个地址。 需要额外维护一个donators 数组
    address[] public _donatorList;
    constructor() {
        _owner = msg.sender;
    }
    /**
     * 时间限制：添加一个时间限制，只有在特定时间段内才能捐赠。
     * 实现思路：比如说7天的捐赠时间  传入开始时间 或者只能工作日捐赠 
     * 或者只能白天捐赠  在donate前面加入require()判断时间是否符合要求
     */
    // 修饰器：仅所有者可调用
    modifier onlyOwner() {
        require(msg.sender == _owner, "Only owner can call this");
        _;
    }
    //记录捐赠的地址和金额
    event Donation(address,uint256);
    /**
     *  一个 donate 函数，允许用户向合约发送以太币，并记录捐赠信息。
     * 一个 withdraw 函数，允许合约所有者提取所有资金。
     * 一个 getDonation 函数，允许查询某个地址的捐赠金额
     */
    function donate() external payable {
        require(msg.value > 0, "Must send ETH");
        //记录捐赠信息
        //如果金额为0   则代表首次捐赠  往list里面插入
        if(_donators[msg.sender]==0){
          _donatorList.push(msg.sender);
        }
        _donators[msg.sender] += msg.value;
        emit Donation(msg.sender,msg.value);
    }

    function withdraw() external onlyOwner {
        //拿到当前合约的资金
        uint256 balance = address(this).balance;
        require(balance > 0, "No funds to withdraw");
        //向拥有者转去所有的金额
        payable(_owner).transfer(balance);
    }

    function getDonation(address donator) public view returns (uint256) {
        return _donators[donator];
    }

    function top3Donators() public  returns (address[3] memory ,uint256[3] memory ) {
      //冒泡排序 取前3个
      uint loopCount = 3;
      if (_donatorList.length < 3) {
          loopCount = _donatorList.length;
      }
      address[3] memory addrs;
      uint256[3] memory amounts;
      for (uint i = 0; i < loopCount; i++) {
        uint256 tag=0;
        uint256 index=i;
        for(uint j = i; j < _donatorList.length; j++){
          uint256 amount=_donators[_donatorList[j]];
          if(amount>tag){
            tag=amount;
            index=j;
          }
        }
        //最大的金额与地址
        addrs[i]=_donatorList[index];
        amounts[i]=tag;
        if(i != index) { // 避免自己和自己交换
            //并且将_donatorList 第一个跟最大的那个金额互换位置
          address tempAddr=_donatorList[i];//把原先的位置拿出来
          _donatorList[i]=_donatorList[index];//把最大的放过去
          _donatorList[index]=tempAddr;//把原位置的放在最大的位置
        }
      }
      return (addrs,amounts);
    }


    /**
     * ai推荐实现插入法
     * 假设数组地址所对应的值为[9,7,5,3,10,15,16,17]
     * 插入法只需要遍历一遍数组
     * 第一次 [9,0,0]  把9放进去
     * 第二次 [9,7,0]=>[9,7,0]=>[9,7,5]
     * 当10进来的时候 [9,7,5]=>[9,7,7]=>[9,9,7]=>[10,9,7] 然后break
     * 当15进来的时候 [10,9,7]=>[10,9,9]=>[10,10,9]=>[15,10,9] 然后break
     */
    function top3Donators2() public view returns (address[3] memory addrs, uint256[3] memory amount) {
      address[3] memory tops;
      uint256[3] memory topAmounts;

    for (uint i = 0; i < _donatorList.length; i++) {
        uint256 donateAmount = _donators[_donatorList[i]];
        for (uint j = 0; j < 3; j++) {
            if (donateAmount > topAmounts[j]) {
                // 后移
                for (uint k = 2; k > j; k--) {
                    topAmounts[k] = topAmounts[k-1];
                    tops[k] = tops[k-1];
                }
                topAmounts[j] = donateAmount;
                tops[j] = _donatorList[i];
                break;
            }
        }
      }
      return (tops, topAmounts);
    }
}
