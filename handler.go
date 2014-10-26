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

	if time.Now().Sub(runtime.lastSphincterCall).Minutes() > float64(c.SphincterTimeout) {
		s.State.Open = gospaceapi.Unknown
	}

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

	runtime.lastSphincterCall = time.Now()

	if t.Status != runtime.lastDoorState || runtime.init {
		if t.Status {
			s.State.Open = gospaceapi.True
		} else {
			s.State.Open = gospaceapi.False
		}
		s.State.Lastchange = time.Now().Unix()
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
