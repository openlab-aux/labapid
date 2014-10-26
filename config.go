package main

import (
	"encoding/json"
	"io/ioutil"
)

type config struct {
	JSONPath      string            `json:"jsonpath"`
	APITokens     map[string]string `json:"apitokens"`
	ListenAddress string            `json:"listenaddress"`
}

func loadConfig(filename string) (config, error) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return config{}, err
	}
	var c config
	err = json.Unmarshal(bytes, &c)
	if err != nil {
		return config{}, err
	}
	return c, nil
}
