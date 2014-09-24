package main

import "fmt"
import "encoding/json"
import "net/http"

//import "io/ioutil"

func showSpaceAPIHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	json, _ := s.ToJSON()

	fmt.Fprintf(w, json)
}

func changeDoorStatusHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	var t interface{}

	err := decoder.Decode(&t)
	if err != nil {
		http.Error(w, "{\"success\":false}", 400)
		return
	}

	_, ok := t.(map[string]interface{})["status"].(string)
	token, ok := t.(map[string]interface{})["token"].(string)

	if !ok {
		http.Error(w, "{\"success\":false}", 400)
		return
	}

	if !tokenOk(token, c.APITokens) {
		http.Error(w, "{\"success\":false}", 403)
	}

	s.State.Open = false

	fmt.Fprintf(w, "{\"success\":true}")

}
