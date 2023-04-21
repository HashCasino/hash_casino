package utils

import (
	"HashCasino/abi"
	"HashCasino/logger"
	"HashCasino/model"
	"context"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
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
		logger.Error("Get ethereum rpc client error, %v", err.Error())
		return nil
	}
	return client
}

// GetHashCasinoInstance Get hashcasino contarct abi instance
func GetHashCasinoInstance(value *big.Int) (*bind.TransactOpts, *abi.Structname, *ethclient.Client) {
	config := model.ReadConfig()
	signer, client := createSigner(value)
	if signer == nil || client == nil {
		logger.Error("Signer or client is nil")
		return nil, nil, nil
	}
	instance, err := abi.NewStructname(common.HexToAddress(config.ContractAddress), client)
	if err != nil {
		logger.Error("Abi new instance failed, %v", err.Error())
		return nil, nil, nil
	}
	return signer, instance, client
}

// createSigner Create signer
func createSigner(value *big.Int) (*bind.TransactOpts, *ethclient.Client) {
	config := model.ReadConfig()
	chainID := config.MainNetworkChainID
	if config.TestNetwork {
		chainID = config.TestNetworkChainID
	}
	privateKey, err := crypto.HexToECDSA(config.AccountPrivateKey)
	if err != nil {
		logger.Error("Crypto hex to ecdsa error, %v", err.Error())
		return nil, nil
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		logger.Error("Error casting public key to ECDSA")
		return nil, nil
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	client := GetEthClient()
	if client == nil {
		logger.Error("Get ethereum rpc client error, client is nil")
		return nil, nil
	}
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		logger.Error("Get nonce error, %v", err.Error())
		return nil, nil
	}
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		logger.Error("Get ethereum rpc client error, client is nil")
		return nil, nil
	}
	signer, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(chainID))
	if err != nil {
		logger.Error("Create signer error, %v", err.Error())
		return nil, nil
	}
	signer.Nonce = big.NewInt(int64(nonce))
	signer.Value = value
	signer.GasLimit = config.GasLimit
	signer.GasPrice = gasPrice
	return signer, client
}
