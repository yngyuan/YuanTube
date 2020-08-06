package config

import (
"encoding/json"
"os"
)

type Configuration struct {
	LBAddr string `json:"lb_addr"`
	OssAddr string `json:"oss_addr"`
}

var configuration *Configuration

func init() {
	file, _ := os.Open("/Users/youngyuan/go/src/YuanTube/bin/conf.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration = &Configuration{}

	err := decoder.Decode(configuration)
	if err != nil {
		panic(err)
	}
}

func GetLbAddr() string {
	return configuration.LBAddr
}

func GetOssAddr() string {
	return configuration.OssAddr
}


