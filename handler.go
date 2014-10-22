package main

import (
	"fmt"
	"time"

	"encoding/json"
	"net/http"
)

func showSpaceAPIHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	json, _ := s.ToJSON()

	fmt.Fprintf(w, json)
}

func changeDoorStatusHandler(w http.ResponseWriter, r *http.Request) {
	type doorstatus struct {
		Token  string
		Status bool
	}

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	t := doorstatus{}

	err := decoder.Decode(&t)
	if err != nil {
		http.Error(w, "{\"success\":false, \"reason\": \""+err.Error()+"\"}", 400)
		return
	}

	if !tokenOk(t.Token, c.APITokens) {
		http.Error(w, "{\"success\":false}", 403)
		return
	}

	if t.Status != runtime.lastDoorState || runtime.init {
		s.State.Open = t.Status
		s.State.Lastchange = int32(time.Now().Unix())
		saveSpaceAPIData(&s, c.JSONPath)
		runtime.lastDoorState = t.Status
		runtime.init = false
	}

	fmt.Fprintf(w, "{\"success\":true}")
}

func changeSensorStatusHandler(w http.ResponseWriter, r *http.Request) {
	type sensordata struct {
		Token  string
		Sensor map[string]interface{}
	}

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	t := sensordata{}

	err := decoder.Decode(&t)

	if err != nil {
		http.Error(w, "{\"success\":false}", 400)
		l.Error(err.Error())
		return
	}

	if !tokenOk(t.Token, c.APITokens) {
		http.Error(w, "{\"success\":false}", 403)
	}

	// extract sensor name
	if len(t.Sensor) != 1 {
		http.Error(w, "{\"success\":false}", 400)
		l.Error(err.Error())
		return
	}
	var sensorname string
	for k, _ := range t.Sensor {
		sensorname = k
	}

	if s.Sensors == nil {
		s.Sensors = map[string]interface{}{}
	}
	s.Sensors[sensorname] = t.Sensor[sensorname]

	saveSpaceAPIData(&s, c.JSONPath)

	fmt.Fprintf(w, "{\"success\":true}")
}
