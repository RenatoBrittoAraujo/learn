package engine

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func loadConfigs() (*Config, bool) {
	jsonFile, err := os.Open("config.json")
	if err != nil {
		fmt.Println(err)
		return nil, false
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var config *Config
	json.Unmarshal(byteValue, config)

	return config, false
}
