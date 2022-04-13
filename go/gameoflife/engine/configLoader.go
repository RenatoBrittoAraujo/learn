package engine

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func loadConfigs() (*Config, bool) {
	jsonFile, err := os.Open("configs.json")

	if err != nil {
		fmt.Println(err)
		return nil, false
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var config Config
	unmarshallErr := json.Unmarshal(byteValue, &config)

	if unmarshallErr != nil {
		fmt.Println(unmarshallErr)
		return nil, false
	}

	return &config, false
}
