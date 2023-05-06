// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

import "https://github.com/OpenZeppelin/openzeppelin-contracts/blob/master/contracts/utils/math/SafeMath.sol";

// ERRORS
error WithdrawError();
error OwnerCheckError();
error BasicError();

contract HashCasino {
    // EVENTS
    event ErrorLog(address indexed sender, string message);
    event WinnerLog(address indexed sender, uint gameType, uint betType, uint position, uint256 amount, uint multiple);
    event LoserLog(address indexed sender, uint gameType, uint betType, uint position, uint256 amount, uint multiple);
    event ReceiveLog(address indexed sender, uint value);

    struct BetInfo {
        uint gameType;
        uint betType;
        uint position;
        uint256 amount;
        address owner;
    }

    uint private prevDiceNum;
    uint private prevRouletteNum;
    address private owner;
    address[] private players;
    BetInfo[] private betInfos;
    uint[] private rouletteNums;

    constructor() {
        owner = msg.sender;
    }

    receive() external payable {
        emit ReceiveLog(msg.sender, msg.value);
    }

    modifier ownerOnly() {
        if (msg.sender != owner) {
            emit ErrorLog(msg.sender, "Caller is not owner");
            revert OwnerCheckError();
        }
        _;
    }

    modifier participateLimit() {
        if (msg.value < .01 ether) {
            emit ErrorLog(msg.sender, "Participate falied, bet amount < 0.01 ICT");
            revert BasicError();
        }
        if (msg.value > 10 ether) {
            emit ErrorLog(msg.sender, "Participate falied, bet amount > 10 ICT");
            revert BasicError();
        }
        _;
    }

    modifier depositLimit() {
        if (msg.value < .01 ether) {
            emit ErrorLog(msg.sender, "Deposit falied, bet amount < 0.01 ICT");
            revert BasicError();
        }
        if (msg.value > 1000 ether) {
            emit ErrorLog(msg.sender, "Deposit falied, bet amount > 1000 ICT");
            revert BasicError();
        }
        _;
    }

    function setRouletteNums(uint[] calldata nums) external ownerOnly {
        if (nums.length <= 0) {
            emit ErrorLog(msg.sender, "Setting Roulette Numbers failed, nums length is zero");
            revert BasicError();
        }
        rouletteNums = nums;
    }

    function participate(uint gameType, uint betType, uint position, uint256 amount) public payable participateLimit {
        bool isParticipate = getBetStatus();
        if (!isParticipate) {
            emit ErrorLog(msg.sender, "Do not place repeated bets");
            revert BasicError();
        }
        BetInfo memory betInfo = BetInfo(gameType, betType, position, amount, msg.sender);
        betInfos.push(betInfo);
        players.push(msg.sender);
    }

    function pickWinners() external ownerOnly {
        if (betInfos.length <= 0) {
            emit ErrorLog(msg.sender, "Pick winners failed, betInfos length is zero");
            revert BasicError();
        }
        // Pick dice winners
        uint randNum = uint(keccak256(abi.encodePacked(block.prevrandao, block.timestamp, block.number, block.coinbase, players)));
        uint diceNum = (randNum % 6) + 1;
        uint position = getPosition(diceNum);
        uint color = getColor(diceNum);
        // Pick roulette winners
        uint rouletteIndex = (randNum % rouletteNums.length) + 1;
        uint rouletteNum = rouletteNums[rouletteIndex];
        uint rouletteColor = getRouletteColor(rouletteNum);
        prevDiceNum = diceNum;
        prevRouletteNum = rouletteNum;
        // Payments
        for (uint i = 0; i < betInfos.length; i++) {
            BetInfo memory betInfo = betInfos[i];
            // Dice
            if (betInfo.gameType == 0) {
                // Size (2x)
                if (betInfo.betType == 0) {
                    if (betInfo.position == position) {
                        emit WinnerLog(betInfo.owner, betInfo.gameType, betInfo.betType, betInfo.position, betInfo.amount, 2);
                        (bool success,) = payable(betInfo.owner).call{value: SafeMath.mul(betInfo.amount, 2)}("");
                        if(!success){
                            emit ErrorLog(betInfo.owner, "Pick winners failed, transfer error");
                            revert BasicError();
                        }
                        continue;
                    }
                    emit LoserLog(betInfo.owner, betInfo.gameType, betInfo.betType, betInfo.position, betInfo.amount, 2);
                }
                // Red or Blue (2x)
                if (betInfo.betType == 1) {
                    if (betInfo.position == color) {
                        emit WinnerLog(betInfo.owner, betInfo.gameType, betInfo.betType, betInfo.position, betInfo.amount, 2);
                        (bool success,) = payable(betInfo.owner).call{value: SafeMath.mul(betInfo.amount, 2)}("");
                        if(!success){
                            emit ErrorLog(betInfo.owner, "Pick winners failed, transfer error");
                            revert BasicError();
                        }
                        continue;
                    }
                    emit LoserLog(betInfo.owner, betInfo.gameType, betInfo.betType, betInfo.position, betInfo.amount, 2);
                }
                // Number (6x)
                if (betInfo.betType == 2) {
                    if (betInfo.position == diceNum) {
                        emit WinnerLog(betInfo.owner, betInfo.gameType, betInfo.betType, betInfo.position, betInfo.amount, 6);
                        (bool success,) = payable(betInfo.owner).call{value: SafeMath.mul(betInfo.amount, 6)}("");
                        if(!success){
                            emit ErrorLog(betInfo.owner, "Pick winners failed, transfer error");
                            revert BasicError();
                        }
                        continue;
                    }
                    emit LoserLog(betInfo.owner, betInfo.gameType, betInfo.betType, betInfo.position, betInfo.amount, 6);
                }
            }
            // Roulette
            if (betInfo.gameType == 1) {
                // Color (2x)
                if (betInfo.betType == 0) {
                    uint multiple = 2;
                    // Green Color (6x)
                    if (rouletteColor == 2) {
                        multiple = 6;
                    }
                    if (betInfo.position == rouletteColor) {
                        emit WinnerLog(betInfo.owner, betInfo.gameType, betInfo.betType, betInfo.position, betInfo.amount, multiple);
                        (bool success,) = payable(betInfo.owner).call{value: SafeMath.mul(betInfo.amount, multiple)}("");
                        if(!success){
                            emit ErrorLog(betInfo.owner, "Pick winners failed, transfer error");
                            revert BasicError();
                        }
                        continue;
                    }
                    emit LoserLog(betInfo.owner, betInfo.gameType, betInfo.betType, betInfo.position, betInfo.amount, multiple);
                }
                // Number (10x)
                if (betInfo.betType == 1) {
                    if (betInfo.position == rouletteNum) {
                        emit WinnerLog(betInfo.owner, betInfo.gameType, betInfo.betType, betInfo.position, betInfo.amount, 10);
                        (bool success,) = payable(betInfo.owner).call{value: SafeMath.mul(betInfo.amount, 10)}("");
                        if(!success){
                            emit ErrorLog(betInfo.owner, "Pick winners failed, transfer error");
                            revert BasicError();
                        }
                        continue;
                    }
                    emit LoserLog(betInfo.owner, betInfo.gameType, betInfo.betType, betInfo.position, betInfo.amount, 10);
                }
            }
        }
        players = new address[](0);
        delete betInfos;
    }

    function deposit() public payable depositLimit {}

    function withdraw() external ownerOnly {
       (bool success,) =  payable(msg.sender).call{value: address(this).balance}("");
        if(!success){
            emit ErrorLog(msg.sender, "Withdraw failed, msg.sender call error");
            revert WithdrawError();
        }
    }

    function getBetInfos() public view returns(BetInfo[] memory) {
        return betInfos;
    }

    function getPoolAmount() public view returns(uint256) {
        uint256 amount = 0;
        for (uint i = 0; i < betInfos.length; i++) {
            BetInfo memory betInfo = betInfos[i];
            amount += betInfo.amount;
        }
        return amount;
    }

    function getPrevNum() public view returns (uint) {
        return prevDiceNum;
    }

    function getPrevRouletteNum() public view returns (uint) {
        return prevRouletteNum;
    }
    
    function getBetStatus() public view returns (bool) {
        bool isParticipate = true;
        for (uint i = 0; i < betInfos.length; i++) {
            BetInfo memory betInfo = betInfos[i];
            if (betInfo.owner == msg.sender) {
                isParticipate = false;
                break;
            }
        }
        return isParticipate;
    }

    function getRouletteNums() public view returns (uint[] memory) {
        return rouletteNums;
    }

    function getPosition(uint diceNum) private pure returns (uint) {
        if (diceNum < 0) {
            revert BasicError();
        }
        if (diceNum <= 3) {
            return 0;
        }
        if (diceNum > 3 && diceNum <= 6) {
            return 1;
        }
        return 2;
    }

    function getColor(uint diceNum) private pure returns (uint) {
        if (diceNum < 0) {
            revert BasicError();
        }
        if (diceNum == 1 || diceNum == 3|| diceNum == 5) {
            return 0;
        }
        return 1;
    }

    function getRouletteColor(uint num) private pure returns (uint) {
        uint mod = num % 2;
        // Specific number (Green)
        if (num == 4 || num == 12 || num == 19 || num == 24 || num == 28 || num == 31) {
            return 2;
        }
        // Odd number (Red)
        if (mod == 0) {
            return 0;
        }
        // Even number (Blue)
        if (mod == 1) {
            return 1;
        }
        return 3;
    }
}