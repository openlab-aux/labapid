package main

import "fmt"
import "os"
import "io/ioutil"
import "encoding/json"
import "net/http"
import "github.com/waaaaargh/gospaceapi"

var s spaceapi.SpaceAPI
var c config

func loadSpaceAPIData(filename string) (spaceapi.SpaceAPI, error) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return spaceapi.SpaceAPI{}, err
	}
	var s spaceapi.SpaceAPI
	err = json.Unmarshal(bytes, &s)
	if err != nil {
		return spaceapi.SpaceAPI{}, err
	}
	return s, nil
}

func saveSpaceAPIData(s *spaceapi.SpaceAPI, filename string) error {
	json, err := s.ToJSON()
	if err != nil {
		return err
	}
	bytes := []byte(json)
	err = ioutil.WriteFile(filename, bytes, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	var err error
	c, err = loadConfig("spaceapi.json.conf")
	if err != nil {
		fmt.Println("[!] Error loading config at 'spaceapi.json.conf': " + err.Error())
		os.Exit(1)
	} else {
		fmt.Println("[i] Config loaded successfully")
	}

	s, err = loadSpaceAPIData(c.JSONPath)
	if err != nil {
		fmt.Println("[!] Error reading SpaceAPI Data: " + err.Error())
		os.Exit(1)
	}

	http.HandleFunc("/", showSpaceAPIHandler)
	http.HandleFunc("/edit/door/", changeDoorStatusHandler)

	err = http.ListenAndServe(":5000", nil)
	if err != nil {
		fmt.Println("[!] Could not start HTTP Server: " + err.Error())
	} else {
		fmt.Println("[i] Starting HTTP Server, waiting for requests")
	}
}
