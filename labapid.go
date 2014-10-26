package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"

	"./logging"
	"github.com/waaaaargh/gospaceapi"
)

var s spaceapi.SpaceAPI
var c config
var l *logging.Logger

// Global runtime data
var runtime struct {
	lastDoorState bool
	init          bool
}

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
	l = logging.New(logging.INFO, "[labapid] [%s] %s\n", os.Stdout)

	runtime.init = true

	var err error
	c, err = loadConfig("spaceapi.json.conf")
	if err != nil {
		l.Critical("[!] Error loading config at 'spaceapi.json.conf': " + err.Error())
		os.Exit(1)
	} else {
		l.Debug("Config loaded successfully")
	}

	s, err = loadSpaceAPIData(c.JSONPath)
	if err != nil {
		l.Critical("[!] Error reading SpaceAPI Data: " + err.Error())
		os.Exit(1)
	}

	http.HandleFunc("/", showSpaceAPIHandler)
	http.HandleFunc("/edit/door/", changeDoorStatusHandler)
	http.HandleFunc("/edit/sensor/", changeSensorStatusHandler)

	err = http.ListenAndServe(c.ListenAddress, nil)
	if err != nil {
		l.Critical("[!] Could not start HTTP Server: " + err.Error())
	} else {
		l.Info("[i] Starting HTTP Server, waiting for requests")
	}
}
