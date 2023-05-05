package main

import (
	"HashCasino/abi"
	"HashCasino/logger"
	"HashCasino/utils"
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"time"
)

var (
	timeDefine        = 0
	notifyChan        = make(chan bool)
	errorLogSenders   []common.Address
	errorLogChan      = make(chan *abi.StructnameErrorLog)
	winnerLogSenders  []common.Address
	winnerLogChan     = make(chan *abi.StructnameWinnerLog)
	loserLogSenders   []common.Address
	loserLogChan      = make(chan *abi.StructnameLoserLog)
	receiveLogSenders []common.Address
	receiveLogChan    = make(chan *abi.StructnameReceiveLog)
)

func main() {
	rouletteNums := utils.RandRouletteNums()
	SettingRouletteNumbers(rouletteNums)
	go ListenPickwinner()
	go ListenDappEvents()
	for {
		select {
		case <-notifyChan:
			go Pickwinner()
			go ListenPickwinner()
		case errorLog := <-errorLogChan:
			logger.Info("Received contract error event, Message: %v, Sender: %v", errorLog.Message, errorLog.Sender)
		case winnerLog := <-winnerLogChan:
			logger.Success("Received contract winner event, GameType: %v, BetType: %v, Position: %v, Amount: %v ICT, Multiple: %v, Sender: %v", winnerLog.GameType, winnerLog.BetType, winnerLog.Position, utils.ToEther(winnerLog.Amount).String(), winnerLog.Multiple, winnerLog.Sender)
		case loserLog := <-loserLogChan:
			logger.Warn("Received contract loser event, GameType: %v, BetType: %v, Position: %v, Amount: %v ICT, Multiple: %v, Sender: %v", loserLog.GameType, loserLog.BetType, loserLog.Position, utils.ToEther(loserLog.Amount).String(), loserLog.Multiple, loserLog.Sender)
		case receiveLog := <-receiveLogChan:
			logger.Success("Received contract receive event, Sender: %v, Amount: %v ICT", receiveLog.Sender, utils.ToEther(receiveLog.Value).String())
		}
	}
}

func ListenPickwinner() {
	time.Sleep(time.Second * 3)
	timestamp := utils.GetTaobaoTime()
	taobaoTime := time.Unix(timestamp, 0)
	second := taobaoTime.Second()
	signer, hashCasinoInstance, _ := utils.GetHashCasinoInstance(big.NewInt(0))
	if signer == nil || hashCasinoInstance == nil {
		logger.Error("Signer or hash casino instance is nil")
		ListenPickwinner()
		return
	}
	if second <= 30 {
		timeDefine = 29 - second
		timer := time.NewTimer(time.Second * time.Duration(timeDefine))
	End:
		for {
			select {
			case <-timer.C:
				notifyChan <- true
				break End
			}
		}
		return
	}
	if second >= 30 {
		timeDefine = 60 - second
		timer := time.NewTimer(time.Second * time.Duration(timeDefine))
		for {
			select {
			case <-timer.C:
				notifyChan <- true
				break
			}
		}
	}
}

// Pickwinner Pick winner
func Pickwinner() {
	signer, hashCasinoInstance, client := utils.GetHashCasinoInstance(big.NewInt(0))
	if signer == nil || hashCasinoInstance == nil {
		logger.Info("Signer or hash casino instance is nil")
		Pickwinner()
		return
	}
	amount, err := hashCasinoInstance.GetPoolAmount(&bind.CallOpts{})
	if err != nil {
		logger.Error("Pick winners error, %v", err.Error())
		Pickwinner()
		return
	}
	if amount.Int64() <= 0 {
		logger.Info("Abort pick winner, pool amount is zero")
		return
	}
	etherAmount := utils.ToEther(amount).String()
	tx, err := hashCasinoInstance.PickWinners(signer)
	if err != nil {
		logger.Error("Pick winners error, %v", err.Error())
		Pickwinner()
		return
	}
	logger.Success("Successful lottery, Liquidation Amount: %v ICT, TxID: %v", etherAmount, tx.Hash())
	client.Close()
}

// SettingRouletteNumbers Setting roulette numbers
func SettingRouletteNumbers(rouletteNums []int) {
	var nums []*big.Int
	for _, v := range rouletteNums {
		nums = append(nums, big.NewInt(int64(v)))
	}
	signer, hashCasinoInstance, client := utils.GetHashCasinoInstance(big.NewInt(0))
	if signer == nil || hashCasinoInstance == nil {
		logger.Info("Signer or hash casino instance is nil")
		return
	}
	tx, err := hashCasinoInstance.SetRouletteNums(signer, nums)
	if err != nil {
		logger.Error("Setting roulette numbers failed, %v", err.Error())
		return
	}
	logger.Success("Setting roulette numbers successful, TxID: %v", tx.Hash())
	client.Close()
}

// ListenDappEvents Listen dapp events notify
func ListenDappEvents() {
	signer, hashCasinoInstance, _ := utils.GetHashCasinoInstance(big.NewInt(0))
	if signer == nil || hashCasinoInstance == nil {
		logger.Info("Signer or hash casino instance is nil")
		return
	}
	_, err := hashCasinoInstance.WatchWinnerLog(&bind.WatchOpts{
		Start:   nil,
		Context: context.Background(),
	}, winnerLogChan, winnerLogSenders)
	if err != nil {
		logger.Error("Watch winner log failed, %v", err.Error())
	}
	_, err = hashCasinoInstance.WatchLoserLog(&bind.WatchOpts{
		Start:   nil,
		Context: context.Background(),
	}, loserLogChan, loserLogSenders)
	if err != nil {
		logger.Error("Watch loser log failed, %v", err.Error())
	}
	_, err = hashCasinoInstance.WatchErrorLog(&bind.WatchOpts{
		Start:   nil,
		Context: context.Background(),
	}, errorLogChan, errorLogSenders)
	if err != nil {
		logger.Error("Watch error log failed, %v", err.Error())
	}
	_, err = hashCasinoInstance.WatchReceiveLog(&bind.WatchOpts{
		Start:   nil,
		Context: context.Background(),
	}, receiveLogChan, receiveLogSenders)
	if err != nil {
		logger.Error("Watch receive log failed, %v", err.Error())
		return
	}
}
