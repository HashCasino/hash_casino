package model

import (
	"HashCasino/logger"
	"encoding/json"
	"os"
)

type Config struct {
	MainNetworkRpcUrl  string `json:"main_network_rpc_url"`
	MainNetworkChainID int64  `json:"main_network_chain_id"`
	TestNetworkRpcUrl  string `json:"test_network_rpc_url"`
	TestNetworkChainID int64  `json:"test_network_chain_id"`
	ContractAddress    string `json:"contract_address"`
	AccountPrivateKey  string `json:"account_private_key"`
	GasLimit           uint64 `json:"gas_limit"`
	TestNetwork        bool   `json:"test_network"`
}

func ReadConfig() *Config {
	var config *Config
	var defaultConfig = &Config{
		MainNetworkRpcUrl:  "wss://nodewebsoket.ic-plaza.org/ws",
		MainNetworkChainID: 9000,
		TestNetworkRpcUrl:  "https://rpctest.ic-plaza.org/",
		TestNetworkChainID: 10000,
		GasLimit:           500000,
		ContractAddress:    "YOUR CONTRACT ADDRESS",
		AccountPrivateKey:  "YOUR ACCOUNT PRIVATE KEY",
		TestNetwork:        false,
	}
	cJson, err := os.ReadFile("./config.json")
	if err != nil {
		defaultConfig.SaveConfig()
		return defaultConfig
	}
	err = json.Unmarshal(cJson, &config)
	if err != nil {
		defaultConfig.SaveConfig()
		return defaultConfig
	}
	return config
}

func (c *Config) SaveConfig() {
	cJson, err := json.Marshal(&c)
	if err != nil {
		logger.Error("Save config failed, %v", err.Error())
		return
	}
	err = os.WriteFile("./config.json", cJson, os.ModePerm)
	if err != nil {
		logger.Error("Save config failed, %v", err.Error())
		return
	}
}
