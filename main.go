package main

import (
	"HashCasino/abi"
	"HashCasino/utils"
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	log "github.com/sirupsen/logrus"
	"io"
	"math/big"
	"os"
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

func init() {
	// Log rotate setting
	writer, err := rotatelogs.New("logs/%Y-%m/%Y-%m-%d.log")
	if err != nil {
		log.Error(err)
		return
	}
	// Log config
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})
	log.SetOutput(io.MultiWriter(writer, os.Stdout))
	log.SetReportCaller(true)
	log.SetLevel(log.TraceLevel)
}

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
			log.Errorf("Received contract error event, Message: %v, Sender: %v", errorLog.Message, errorLog.Sender)
		case winnerLog := <-winnerLogChan:
			log.Infof("Received contract winner event, GameName: %v, BetType: %v, Position: %v, Amount: %v ICT, Multiple: %vx, Sender: %v, Winner Amount: %v ICT", utils.GetNameByGameType(winnerLog.GameType.Int64()), utils.GetNameByBetType(winnerLog.GameType.Int64(), winnerLog.BetType.Int64()), winnerLog.Position, utils.ToEther(winnerLog.Amount).String(), winnerLog.Multiple, winnerLog.Sender, utils.ToEther(big.NewInt(winnerLog.Amount.Int64()*winnerLog.Multiple.Int64())).String())
		case loserLog := <-loserLogChan:
			log.Warnf("Received contract loser event, GameName: %v, BetType: %v, Position: %v, Amount: %v ICT, Multiple: %vx, Sender: %v", utils.GetNameByGameType(loserLog.GameType.Int64()), utils.GetNameByBetType(loserLog.GameType.Int64(), loserLog.BetType.Int64()), loserLog.Position, utils.ToEther(loserLog.Amount).String(), loserLog.Multiple, loserLog.Sender)
		case receiveLog := <-receiveLogChan:
			log.Infof("Received contract receive event, Sender: %v, Amount: %v ICT", receiveLog.Sender, utils.ToEther(receiveLog.Value).String())
		}
	}
}

func ListenPickwinner() {
	time.Sleep(time.Second * 1)
	timestamp := utils.GetTaobaoTime()
	taobaoTime := time.Unix(timestamp, 0)
	second := taobaoTime.Second()
	signer, hashCasinoInstance, _ := utils.GetHashCasinoInstance(big.NewInt(0))
	if signer == nil || hashCasinoInstance == nil {
		log.Error("Signer or hash casino instance is nil")
		ListenPickwinner()
		return
	}
	if second <= 30 {
		timeDefine = 30 - second
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
		log.Info("Signer or hash casino instance is nil")
		Pickwinner()
		return
	}
	amount, err := hashCasinoInstance.GetPoolAmount(&bind.CallOpts{})
	if err != nil {
		log.Errorf("Pick winners error, %v", err.Error())
		Pickwinner()
		return
	}
	if amount.Int64() <= 0 {
		log.Info("Abort pick winner, pool amount is zero")
		return
	}
	etherAmount := utils.ToEther(amount).String()
	tx, err := hashCasinoInstance.PickWinners(signer)
	if err != nil {
		log.Errorf("Pick winners error, %v", err.Error())
		Pickwinner()
		return
	}
	log.Infof("Successful lottery, Liquidation Amount: %v ICT, TxID: %v", etherAmount, tx.Hash())
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
		log.Info("Signer or hash casino instance is nil")
		return
	}
	tx, err := hashCasinoInstance.SetRouletteNums(signer, nums)
	if err != nil {
		log.Errorf("Setting roulette numbers failed, %v", err.Error())
		return
	}
	log.Infof("Setting roulette numbers successful, TxID: %v", tx.Hash())
	client.Close()
}

// ListenDappEvents Listen dapp events notify
func ListenDappEvents() {
	signer, hashCasinoInstance, _ := utils.GetHashCasinoInstance(big.NewInt(0))
	if signer == nil || hashCasinoInstance == nil {
		log.Info("Signer or hash casino instance is nil")
		return
	}
	_, err := hashCasinoInstance.WatchWinnerLog(&bind.WatchOpts{
		Start:   nil,
		Context: context.Background(),
	}, winnerLogChan, winnerLogSenders)
	if err != nil {
		log.Errorf("Watch winner log failed, %v", err.Error())
	}
	_, err = hashCasinoInstance.WatchLoserLog(&bind.WatchOpts{
		Start:   nil,
		Context: context.Background(),
	}, loserLogChan, loserLogSenders)
	if err != nil {
		log.Errorf("Watch loser log failed, %v", err.Error())
	}
	_, err = hashCasinoInstance.WatchErrorLog(&bind.WatchOpts{
		Start:   nil,
		Context: context.Background(),
	}, errorLogChan, errorLogSenders)
	if err != nil {
		log.Errorf("Watch error log failed, %v", err.Error())
	}
	_, err = hashCasinoInstance.WatchReceiveLog(&bind.WatchOpts{
		Start:   nil,
		Context: context.Background(),
	}, receiveLogChan, receiveLogSenders)
	if err != nil {
		log.Errorf("Watch receive log failed, %v", err.Error())
		return
	}
}
