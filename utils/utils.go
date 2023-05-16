package utils

import (
	"encoding/json"
	"github.com/shopspring/decimal"
	"io"
	"math"
	"math/big"
	"math/rand"
	"net/http"
	"time"
)

// GetTaobaoTime Get Taobao time
func GetTaobaoTime() int64 {
	url := "http://api.m.taobao.com/rest/api3.do?api=mtop.common.getTimestamp"
	client := &http.Client{
		Timeout: 8,
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return time.Now().Unix()
	}
	res, err := client.Do(req)
	if err != nil {
		return time.Now().Unix()
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return time.Now().Unix()
	}
	jsonMap := map[string]interface{}{}
	err = json.Unmarshal(body, &jsonMap)
	if err != nil {
		return time.Now().Unix()
	}
	timeStamp := jsonMap["data"].(map[string]interface{})["t"].(int64)
	return timeStamp
}

// RandRouletteNums Random roulette numbers
func RandRouletteNums() []int {
	var arr []int
	for i := 1; i <= 32; i++ {
		arr = append(arr, i)
	}
	for i := len(arr) - 1; i > 0; i-- {
		randInt := rand.Intn(i)
		arr[randInt], arr[i] = arr[i], arr[randInt]
	}
	return arr
}

// ToWei Decimals to wei
func ToWei(value interface{}, decimals int) *big.Int {
	amount := decimal.NewFromFloat(0)
	switch v := value.(type) {
	case string:
		amount, _ = decimal.NewFromString(v)
	case float64:
		amount = decimal.NewFromFloat(v)
	case int64:
		amount = decimal.NewFromFloat(float64(v))
	case decimal.Decimal:
		amount = v
	case *decimal.Decimal:
		amount = *v
	}
	mul := decimal.NewFromFloat(float64(10)).Pow(decimal.NewFromFloat(float64(decimals)))
	result := amount.Mul(mul)
	wei := new(big.Int)
	wei.SetString(result.String(), 10)
	return wei
}

// ToEther Wei to ether
func ToEther(wei *big.Int) *big.Float {
	balance := new(big.Float)
	balance.SetString(wei.String())
	ether := new(big.Float).Quo(balance, big.NewFloat(math.Pow10(18)))
	return ether
}

// GetNameByGameType Get game name by game type
func GetNameByGameType(gameType int64) string {
	var gameName string
	switch gameType {
	case 0:
		gameName = "Dice"
		break
	case 1:
		gameName = "Big Wheel"
		break
	default:
		gameName = "Unkonw"
		break
	}
	return gameName
}

// GetNameByBetType Get game bet type by game type and bet type
func GetNameByBetType(gameType, betType int64) string {
	var betName string
	// Dice
	if gameType == 0 {
		switch betType {
		case 0:
			betName = "Size (2x)"
			break
		case 1:
			betName = "Red or Blue (2x)"
			break
		case 2:
			betName = "Number (6x)"
			break
		default:
			betName = "Unkonw (0x)"
			break
		}
	}
	// Big Wheel
	if gameType == 1 {
		switch betType {
		case 0:
			betName = "Red or Blue (2x) / Green (6x)"
			break
		case 1:
			betName = "Number (10x)"
			break
		default:
			betName = "Unkonw (0x)"
			break
		}
	}
	return betName
}
