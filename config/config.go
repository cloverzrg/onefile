package config

import (
	"encoding/json"
	"os"
	"sync"
	"time"

	"github.com/cloverzrg/onefile/logger"
)

type config struct {
	Baseurl  string `json:"baseurl"`
	OneDrive struct {
		ClientId     string   `json:"client_id"`
		ClientSecret string   `json:"client_secret"`
		RedirectUri  string   `json:"redirect_uri"`
		Scope        []string `json:"scope"`
	} `json:"onedrive"`
}

type consulConfig struct {
	Address    string `json:"address"`
	Token      string `json:"token"`
	ConfigFile string `json:"config_file"`
}

var Config config
var ConsulConfig consulConfig
var mu sync.Mutex

func BindJson(data []byte) (err error) {
	mu.Lock()
	defer mu.Unlock()
	err = json.Unmarshal(data, &Config)
	return err
}

func init() {
	if os.Getenv("NO_CONFIG") != "" {
		logger.Info("NO_CONFIG 有值，不加载配置文件")
		return
	}
	if configFile := os.Getenv("CONFIG_FILE"); configFile != "" {
		file, err := os.ReadFile(configFile)
		if err != nil {
			logger.Errorf("read config file error %s", configFile)
			return
		}
		err = json.Unmarshal(file, &Config)
		if err != nil {
			logger.Errorf("read config file error %s", configFile)
			return
		}
		return
	}
	ConsulConfig = consulConfig{
		Address:    os.Getenv("CONSUL_ADDRESS"),
		Token:      os.Getenv("CONSUL_TOKEN"),
		ConfigFile: os.Getenv("CONSUL_CONFIG_PATH"),
	}
	readConfigFromConfig(ConsulConfig.Address, ConsulConfig.Token, ConsulConfig.ConfigFile)
	go func() {
		for {
			time.Sleep(5 * time.Second)
			readConfigFromConfig(ConsulConfig.Address, ConsulConfig.Token, ConsulConfig.ConfigFile)
		}
	}()
}
