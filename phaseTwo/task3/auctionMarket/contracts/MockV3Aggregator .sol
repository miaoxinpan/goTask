// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract MockV3Aggregator {
    uint8 public decimals;
    int256 public latestAnswer;

    uint80 public roundId = 1;
    uint256 public startedAt = block.timestamp;
    uint256 public updatedAt = block.timestamp;
    uint80 public answeredInRound = 1;

    constructor(uint8 _decimals, int256 _initialAnswer) {
        decimals = _decimals;
        latestAnswer = _initialAnswer;
    }

    function updateAnswer(int256 _answer) public {
        latestAnswer = _answer;
        updatedAt = block.timestamp;
        roundId += 1;
        answeredInRound = roundId;
    }

    function latestRoundData()
        external
        view
        returns (
            uint80,
            int256,
            uint256,
            uint256,
            uint80
        )
    {
        return (
            roundId,
            latestAnswer,
            startedAt,
            updatedAt,
            answeredInRound
        );
    }
}