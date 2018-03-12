package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func GetStreams(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	s, err := twitchapi.GetLivestreams(vars["query"])
	if err != nil {
		w.WriteHeader(503)
	} else {
		json.NewEncoder(w).Encode(s)
	}
}

func main() {

	//Router Functions
	router := mux.NewRouter()
	router.HandleFunc("/livestreams/{query}", GetStreams).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}
