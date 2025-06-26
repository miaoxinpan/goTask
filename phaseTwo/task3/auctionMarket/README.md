# Sample Hardhat Project

This project demonstrates a basic Hardhat use case. It comes with a sample contract, a test for that contract, and a Hardhat Ignition module that deploys that contract.

Try running some of the following tasks:

```shell
npx hardhat help
npx hardhat test
REPORT_GAS=true npx hardhat test
npx hardhat node
npx hardhat ignition deploy ./ignition/modules/Lock.js
```
# 任务目标
### ✅  大作业：实现一个 NFT 拍卖市场
任务目标
1. 使用 Hardhat 框架开发一个 NFT 拍卖市场。
2. 使用 Chainlink 的 feedData 预言机功能，计算 ERC20 和以太坊到美元的价格。
3. 使用 UUPS/透明代理模式实现合约升级。
4. 使用类似于 Uniswap V2 的工厂模式管理每场拍卖。

任务步骤
1. 项目初始化
1. 使用 Hardhat 初始化项目：
npx hardhat init
2. 安装必要的依赖：
     npm install @openzeppelin/contracts @chainlink/contracts @nomiclabs/hardhat-ethers hardhat-deploy
2. 实现 NFT 拍卖市场
1. NFT 合约：
  - 使用 ERC721 标准实现一个 NFT 合约。
  - 支持 NFT 的铸造和转移。
2. 拍卖合约：
  - 实现一个拍卖合约，支持以下功能：
  - 创建拍卖：允许用户将 NFT 上架拍卖。
  - 出价：允许用户以 ERC20 或以太坊出价。
  - 结束拍卖：拍卖结束后，NFT 转移给出价最高者，资金转移给卖家。
3. 工厂模式：
  - 使用类似于 Uniswap V2 的工厂模式，管理每场拍卖。
  - 工厂合约负责创建和管理拍卖合约实例。
4. 集成 Chainlink 预言机
5. 价格计算：
  - 使用 Chainlink 的 feedData 预言机，获取 ERC20 和以太坊到美元的价格。
  - 在拍卖合约中，将出价金额转换为美元，方便用户比较。
6. 跨链拍卖：
  - 使用 Chainlink 的 CCIP 功能，实现 NFT 跨链拍卖。
  - 允许用户在不同链上参与拍卖。
7. 合约升级
  1. UUPS/透明代理：
  - 使用 UUPS 或透明代理模式实现合约升级。
  - 确保拍卖合约和工厂合约可以安全升级。
8. 测试与部署
  1. 测试：
  - 编写单元测试和集成测试，覆盖所有功能。
  2. 部署：
  - 使用 Hardhat 部署脚本，将合约部署到测试网（如 Goerli 或 Sepolia）。

任务要求
1. 代码质量：
  - 代码清晰、规范，符合 Solidity 最佳实践。
1. 功能完整性：
  - 实现所有要求的功能，包括 NFT 拍卖、价格计算和合约升级。
1. 测试覆盖率：
  - 编写全面的测试，覆盖所有功能。
1. 文档：
  - 提供详细的文档，包括项目结构、功能说明和部署步骤。

提交内容
1. 代码：提交完整的 Hardhat 项目代码。
2. 测试报告：提交测试报告，包括测试覆盖率和测试结果。
3. 部署地址：提交部署到测试网的合约地址。
4. 文档：提交项目文档，包括功能说明和部署步骤。

额外挑战（可选）
1. 动态手续费：根据拍卖金额动态调整手续费。