package main

import "fmt"
import "net/http"

func showSpaceAPIHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	json, _ := s.ToJSON()

	fmt.Fprintf(w, json)
}
