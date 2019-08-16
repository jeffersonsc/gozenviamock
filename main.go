package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// SingleSMSRequest .
type SingleSMSRequest struct {
	SendSmsRequest struct {
		From           string `json:"from"`
		To             string `json:"to"`
		Schedule       string `json:"schedule"`
		Msg            string `json:"msg"`
		CallbackOption string `json:"callbackOption"`
		ID             string `json:"id"`
		AggregateID    int    `json:"aggregateId"`
		FlashSms       bool   `json:"flashSms"`
	} `json:"sendSmsRequest"`
}

// SingleSMSResponse .
type SingleSMSResponse struct {
	SendSmsResponse `json:"sendSmsResponse"`
}

// SendSmsResponse .
type SendSmsResponse struct {
	StatusCode        string `json:"statusCode"`
	StatusDescription string `json:"statusDescription"`
	DetailCode        string `json:"detailCode"`
	DetailDescription string `json:"detailDescription"`
}

// MutipleSMSRequest .
type MutipleSMSRequest struct {
	SendSmsMultiRequest struct {
		AggregateID        int `json:"aggregateId"`
		SendSmsRequestList []struct {
			Msg            string `json:"msg"`
			Schedule       string `json:"schedule"`
			From           string `json:"from"`
			To             string `json:"to"`
			CallbackOption string `json:"callbackOption"`
			ID             string `json:"id"`
		} `json:"sendSmsRequestList"`
	} `json:"sendSmsMultiRequest"`
}

// MutipleSMSResponse .
type MutipleSMSResponse struct {
	SendSmsMultiResponse `json:"sendSmsMultiResponse"`
}

// SendSmsMultiResponse .
type SendSmsMultiResponse struct {
	SendSmsResponseList []SendSmsResponse `json:"sendSmsResponseList"`
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("IT'S WORK!"))
	}).Methods(http.MethodGet)
	r.HandleFunc("/services/send-sms", singleHandler).Methods(http.MethodPost)
	r.HandleFunc("/services/send-sms-multiple", mutipleHandler).Methods(http.MethodPost)

	log.Println("[main] Server started")
	log.Fatal(http.ListenAndServe(":3000", r))
}

func singleHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	msg := SingleSMSRequest{}
	if err := json.NewDecoder(r.Body).Decode(&msg); err != nil {
		log.Println("[singleHandler] ERROR Failed decode single message ", err.Error())
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(map[string]string{"error": "failed decode json"})
		return
	}

	msgr := SingleSMSResponse{
		SendSmsResponse: SendSmsResponse{
			StatusCode:        "00",
			StatusDescription: "OK",
			DetailCode:        "000",
			DetailDescription: "Message sent",
		},
	}
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(msgr)
}

func mutipleHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	msg := MutipleSMSRequest{}
	if err := json.NewDecoder(r.Body).Decode(&msg); err != nil {
		log.Println("[singleHandler] ERROR Failed decode single message ", err.Error())
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(map[string]string{"error": "failed decode json"})
		return
	}

	msgr := MutipleSMSResponse{}
	for range msg.SendSmsMultiRequest.SendSmsRequestList {
		msgr.SendSmsResponseList = append(msgr.SendSmsResponseList, SendSmsResponse{
			StatusCode:        "00",
			StatusDescription: "OK",
			DetailCode:        "000",
			DetailDescription: "Message sent",
		})
	}

	<-time.After(time.Second * 3)

	w.WriteHeader(200)
	json.NewEncoder(w).Encode(msgr)
}
