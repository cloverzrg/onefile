package config

import (
	"encoding/json"
	"github.com/cloverzrg/onefile/logger"
	"github.com/hashicorp/consul/api"
)

var consulClient *api.Client

func readConfigFromConfig(configKey string) (err error) {
	defer func() {
		if err != nil {
			logger.Error(err)
		}
	}()

	// 此处回读取环境变量
	cfg := api.DefaultConfig()
	consulClient, err = api.NewClient(cfg)
	if err != nil {
		return err
	}
	kvClient := consulClient.KV()
	kvPair, _, err := kvClient.Get(configKey, nil)
	if err != nil {
		return err
	}

	if kvPair == nil {
		UpdateConfig()
		logger.Panicf("不存在配置文件：%s,已添加到consul，请到consul补充配置信息", configKey)
	}

	err = BindJson(kvPair.Value)
	if err != nil {
		return err
	}
	return err
}

func UpdateConfig() (err error) {
	defer func() {
		if err != nil {
			logger.Error(err)
		}
	}()
	jsonbyte, err := json.MarshalIndent(Config, "", "\t")
	if err != nil {
		return err
	}
	kvPair := api.KVPair{
		Key:   consulConfigFile,
		Value: jsonbyte,
	}
	kvClient := consulClient.KV()
	_, err = kvClient.Put(&kvPair, nil)
	if err != nil {
		return err
	}
	return err
}
