package main

import (
	"cs255/logLayer/logServer"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

var default_config_path string = "config.json"

type Config_param struct {
	HttpIP              string `json:"httpIP"`
	HttpPort            string `json:"httpPort"`
	Local_redis_address string `json:"local_redis_address"`
	Log_id              int    `json:"log_id"`
	Replica_ids         []int  `json:"replica_ids"`
}

func main() {
	fmt.Println("It's my GO!!")

	var config_path string

	if len(os.Args) == 1 {
		config_path = default_config_path
	} else {
		config_path = os.Args[1]
	}

	json_bytes, err := ioutil.ReadFile(config_path)
	if err != nil {
		panic("无法打开config文件")
	}

	var config Config_param
	err = json.Unmarshal(json_bytes, &config)
	if err != nil {
		panic("无法解析config文件")
	}

	logServer.Init_ids(config.Log_id, config.Replica_ids)
	logServer.Init_local_redis(config.Local_redis_address)

	logServer.HttpServer(config.HttpIP, config.HttpPort)

}
