package objects

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
)

type Config struct {
	Forecast struct {
		Key string `json="key"`
	} `json="forecast"`
}

var config Config

func readConfig() {
	user, err := user.Current()
	if err != nil {
		fmt.Printf("Cannot find user directory: %s", err)
		os.Exit(1)
	}

	file, err := ioutil.ReadFile(user.HomeDir + "/.goweather/goweather.properties")
	if err != nil {
		fmt.Printf("Failed to read config file: %s", err)
		os.Exit(1)
	}
	json.Unmarshal(file, &config)
}

func GetApiKey() string {
	readConfig()
	return config.Forecast.Key
}
