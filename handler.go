package main

import "fmt"
import "encoding/json"
import "net/http"

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

	s.State.Open = t.Status

	saveSpaceAPIData(&s, c.JSONPath)

	fmt.Fprintf(w, "{\"success\":true}")
}

func changeSensorStatusHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	var t interface{}

	err := decoder.Decode(&t)
	if err != nil {
		http.Error(w, "{\"success\":false}", 400)
		return
	}

	sensor, ok := t.(map[string]interface{})["sensor"].(map[string]interface{})
	token, ok := t.(map[string]interface{})["token"].(string)

	if !ok {
		http.Error(w, "{\"success\":false}", 400)
		return
	}

	if !tokenOk(token, c.APITokens) {
		http.Error(w, "{\"success\":false}", 403)
	}

	// extract sensor name
	var sensorname string
	for k, _ := range sensor {
		sensorname = k
	}

	if s.Sensors == nil {
		s.Sensors = map[string]interface{}{}
	}
	s.Sensors[sensorname] = sensor[sensorname]

	saveSpaceAPIData(&s, c.JSONPath)

	fmt.Fprintf(w, "{\"success\":true}")
}
