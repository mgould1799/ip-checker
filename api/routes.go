package api 

import (
	"encoding/json"
	"net/http"
)

func healthz(w http.ResponseWriter, r *http.Request) {
	// set content type
	w.Header().Set("Content-Type", "application/json")
	// specify status code 
	w.WriteHeader(http.StatusOK)
	  
	// return a json letting know its working
	response := map[string]bool {"working": true}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
	 return
	}
	w.Write(jsonResponse)
}

func (s Server) checkIp(w http.ResponseWriter, r *http.Request) {
	// set content type
	w.Header().Set("Content-Type", "application/json")
	// specify status code 
	w.WriteHeader(http.StatusOK)
	  
	// return a json letting know its working
	response := map[string]bool {"working": true}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
	 return
	}
	w.Write(jsonResponse)
}

