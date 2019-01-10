package config

import (
	"fmt"
	"io/ioutil"
	"library-demo/helper"
	"log"
	"os"
)

type Config struct {
	ServiceName		string	`json:"service_name"`
	Database struct {
		User     string `json:"user"`
		Password string `json:"password"`
		Database string `json:"database"`
		Address  string `json:"address"`
	} `json:"database"`
}

func SetupConfig() Config {
	files, err := ioutil.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		fmt.Println(f.Name())
	}
	var conf Config

	// Đọc file config.local.json
	configFile, err := os.Open("config.local.json")
	if err != nil {
		log.Println(err)
		// Nếu không có file config.local.json thì đọc file config.default.json
		configFile, err = os.Open("config.default.json")
		if err != nil {
			panic(err)
		}
		defer configFile.Close()
	}
	defer configFile.Close()

	// Parse dữ liệu JSON và bind vào Controller
	err = helper.DecodeDataFromJsonFile(configFile, &conf)
	if err != nil {
		log.Println("Không đọc được file config.")
		panic(err)
	}

	return conf
}
