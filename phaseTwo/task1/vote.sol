// SPDX-License-Identifier: SEE LICENSE IN LICENSE
pragma solidity ^0.8;

/**
     * 1.创建一个名为Voting的合约，包含以下功能：
        一个mapping来存储候选人的得票数
        一个vote函数，允许用户投票给某个候选人
        一个getVotes函数，返回某个候选人的得票数
        一个resetVotes函数，重置所有候选人的得票数
 */
contract Vote {
    mapping(address => uint) public candidateVotes; //用来记录多少个候选人，每个候选人的得票数
    mapping(address => bool) public isCandidate; // 判断是否已是候选人
    address[] public candidates;

    constructor() {}

    //投票功能
    function ballot(address candidate) public {
        //因为默认值是0  所以用candidateVotes[candidate]==0 是不行的
        //然后又要考虑清空候选人的得票数  所以得有一个数组来维护候选人的地址
        if (!isCandidate[candidate]) {
            //不是候选人
            isCandidate[candidate] = true; //在是否候选人map里面加入进去
            candidates.push(candidate); //插入数组里面
        }
        candidateVotes[candidate] += 1;
    }

    //返回某个候选人的得票数
    function getVotes(address candidate) public view returns (uint) {
        return candidateVotes[candidate];
    }

    //重置所有候选人的得票数
    function resetVotes() public {
        //遍历数组清空票数
        for (uint i = 0; i < candidates.length; i++) {
            candidateVotes[candidates[i]] = 0;
        }
    }

    /**
    * 2.反转字符串 (Reverse String)
      题目描述：反转一个字符串。输入 "abcde"，输出 "edcba"
    */
    function reverseString(
        string memory inputs
    ) public pure returns (string memory) {
        bytes memory strBytes = bytes(inputs);
        uint left = 0;
        uint right = strBytes.length - 1;
        while (left < right) {
            // 交换
            bytes1 temp = strBytes[left];
            strBytes[left] = strBytes[right];
            strBytes[right] = temp;
            left++;
            right--;
        }
        return string(strBytes);
    }

    /**
    * 3.用 solidity 实现整数转罗马数字
    * 
      题目描述在 https://leetcode.cn/problems/integer-to-roman/description/
      I->1
      V->5
      X->10
      L->50
      C->100
      D->500
      M->1000

      1 <= intnum <= 3999
      由于输入的数字小于4000 所以基本上就4位  【2】【3】【4】【5】
      将罗马数字与对应的整数一一对应好  放在两个数组里面 匹配一个减去对应的值
    */
    function intToRoman(uint256 intnum) public pure returns (string memory) {
        string[13] memory romans = [
            "M",
            "CM",
            "D",
            "CD",
            "C",
            "XC",
            "L",
            "XL",
            "X",
            "IX",
            "V",
            "IV",
            "I"
        ];
        uint16[13] memory values = [
            1000,
            900,
            500,
            400,
            100,
            90,
            50,
            40,
            10,
            9,
            5,
            4,
            1
        ];
        string memory result = "";
        for (uint i = 0; i < 13; i++) {
            while (intnum >= values[i]) {
                result = string(abi.encodePacked(result, romans[i]));
                intnum -= values[i];
            }
        }
        return result;
    }

    /**4.用 solidity 实现罗马数字转数整数
     *  题目描述在 https://leetcode.cn/problems/roman-to-integer/description/3.
     * 得先判断这个位  是百位还是千位 不然算不出来
     *
     * 优先匹配长的罗马数字符号（如 "CM"、"IV"），
     * 每次匹配成功就累加对应的值并跳过已匹配的字符，直到字符串末尾
     */
    function romanToInt(string memory s) public pure returns (uint256) {
        string[13] memory romans = [
            "M",
            "CM",
            "D",
            "CD",
            "C",
            "XC",
            "L",
            "XL",
            "X",
            "IX",
            "V",
            "IV",
            "I"
        ];
        uint16[13] memory values = [
            1000,
            900,
            500,
            400,
            100,
            90,
            50,
            40,
            10,
            9,
            5,
            4,
            1
        ];
        bytes memory strBytes = bytes(s);
        uint256 i = 0;
        uint256 res = 0;
        while (i < strBytes.length) {
            bool matched = false;
            // 优先匹配两个字符的罗马数字
            for (uint j = 0; j < 13; j++) {
                bytes memory romanBytes = bytes(romans[j]);
                if (
                    romanBytes.length == 2 &&
                    i + 1 < strBytes.length &&
                    strBytes[i] == romanBytes[0] &&
                    strBytes[i + 1] == romanBytes[1]
                ) {
                    res += values[j];
                    i += 2;
                    matched = true;
                    break;
                }
                // 匹配单字符
                if (romanBytes.length == 1 && strBytes[i] == romanBytes[0]) {
                    res += values[j];
                    i += 1;
                    matched = true;
                    break;
                }
            }
            // 如果没有匹配到，说明输入有误，直接跳出
            if (!matched) {
                break;
            }
        }
        return res;
    }
}
/**
    5.合并两个有序数组 (Merge Sorted Array)
    题目描述：将两个有序数组合并为一个有序数组。
    6.二分查找 (Binary Search)
    题目描述：在一个有序数组中查找目标值。
 */
