package main

import "io/ioutil"
import "encoding/json"

type config struct {
	JSONPath  string            `json:"jsonpath"`
	APITokens map[string]string `json:"apitokens"`
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
