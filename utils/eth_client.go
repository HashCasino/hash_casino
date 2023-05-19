package utils

import (
	"HashCasino/abi"
	"HashCasino/model"
	"context"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"
	"math/big"
	"time"
)

// GetEthClient Get ethereum rpc client
func GetEthClient() *ethclient.Client {
	config := model.ReadConfig()
	rpcUrl := config.MainNetworkRpcUrl
	if config.TestNetwork {
		rpcUrl = config.TestNetworkRpcUrl
	}
	client, err := ethclient.Dial(rpcUrl)
	if err != nil {
		log.Errorf("Get ethereum rpc client error, %v", err.Error())
		time.Sleep(time.Second * 3)
		return GetEthClient()
	}
	return client
}

// GetHashCasinoInstance Get hashcasino contarct abi instance
func GetHashCasinoInstance(value *big.Int) (*bind.TransactOpts, *abi.Structname, *ethclient.Client) {
	config := model.ReadConfig()
	signer, client := createSigner(value)
	if signer == nil || client == nil {
		log.Error("Signer or client is nil")
		return GetHashCasinoInstance(value)
	}
	instance, err := abi.NewStructname(common.HexToAddress(config.ContractAddress), client)
	if err != nil {
		log.Errorf("Abi new instance failed, %v", err.Error())
		return GetHashCasinoInstance(value)
	}
	return signer, instance, client
}

// createSigner Create signer
func createSigner(value *big.Int) (*bind.TransactOpts, *ethclient.Client) {
	client := GetEthClient()
	config := model.ReadConfig()
	chainID := config.MainNetworkChainID
	if config.TestNetwork {
		chainID = config.TestNetworkChainID
	}
	privateKey, err := crypto.HexToECDSA(config.AccountPrivateKey)
	if err != nil {
		log.Errorf("Crypto hex to ecdsa error, %v", err.Error())
		time.Sleep(time.Second * 1)
		return createSigner(value)
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Error("Error casting public key to ECDSA")
		time.Sleep(time.Second * 1)
		return createSigner(value)
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	if client == nil {
		log.Error("Get ethereum rpc client error, client is nil")
		time.Sleep(time.Second * 1)
		return createSigner(value)
	}
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Errorf("Get nonce error, %v", err.Error())
		time.Sleep(time.Second * 1)
		return createSigner(value)
	}
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Error("Get ethereum rpc client error, client is nil")
		time.Sleep(time.Second * 1)
		return createSigner(value)
	}
	signer, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(chainID))
	if err != nil {
		log.Errorf("Create signer error, %v", err.Error())
		time.Sleep(time.Second * 1)
		return createSigner(value)
	}
	signer.Nonce = big.NewInt(int64(nonce))
	signer.Value = value
	signer.GasLimit = config.GasLimit
	signer.GasPrice = gasPrice
	return signer, client
}
