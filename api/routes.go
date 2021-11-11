package api 

import (
	"encoding/json"
	"net/http"
	log "github.com/sirupsen/logrus"
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
		log.Errorf("error writing json - %s", err.Error())
	 	return
	}
	w.Write(jsonResponse)
}

func (s Server) checkIpInWhiteList(w http.ResponseWriter, r *http.Request) {
	// s.ipCheck.GetIP('')

	// request body and put it into a struct 
	defer r.Body.Close()
	decoder := json.NewDecoder(r.Body)
	requestBody := struct {
		Ip string `json:"ip"`
		WhiteListedCountries []string `json:"whiteListedCountries"`
	}{}
    err := decoder.Decode(&requestBody)
    if err != nil {
        log.Error("error decoding struct - %s", err.Error())
        return
    }


	inWhiteList, err := s.ipCheck.InWhiteList(requestBody.WhiteListedCountries, requestBody.Ip)
	if err != nil {
		log.Error("error with checking in white list - %s", err.Error())
	}


	
	// set content type
	w.Header().Set("Content-Type", "application/json")
	// specify status code 
	w.WriteHeader(http.StatusOK)
	  
	// return a json letting know its working
	response := map[string]bool {"inList": inWhiteList}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		log.Error("error writing response - %s", err.Error())
	 	return
	}
	w.Write(jsonResponse)
}

