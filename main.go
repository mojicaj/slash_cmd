package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", slash)

	http.ListenAndServe(":9080", nil)
}

func slash(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	r := struct {
		ResponseType string
		Text         string
	}{
		ResponseType: "in_channel",
		Text:         "slash command successful",
	}

	if err := json.NewEncoder(w).Encode(r); err != nil {
		fmt.Println(err)
	}
}
