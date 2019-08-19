package mutiple

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jeffersonsc/gozenviamock/pkg/single"
)

// SMSRequest .
type SMSRequest struct {
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

// SMSResponse .
type SMSResponse struct {
	SendSmsMultiResponse `json:"sendSmsMultiResponse"`
}

// SendSmsMultiResponse .
type SendSmsMultiResponse struct {
	SendSmsResponseList []single.SendSmsResponse `json:"sendSmsResponseList"`
}

// RegisterRouter .
func RegisterRouter(r *mux.Router) {
	r.HandleFunc("/services/send-sms-multiple", mutipleHandler).Methods(http.MethodPost)
}

func mutipleHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	msg := SMSRequest{}
	if err := json.NewDecoder(r.Body).Decode(&msg); err != nil {
		log.Println("[singleHandler] ERROR Failed decode single message ", err.Error())
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(map[string]string{"error": "failed decode json"})
		return
	}

	msgr := SMSResponse{}
	for range msg.SendSmsMultiRequest.SendSmsRequestList {
		msgr.SendSmsResponseList = append(msgr.SendSmsResponseList, single.SendSmsResponse{
			StatusCode:        "00",
			StatusDescription: "OK",
			DetailCode:        "000",
			DetailDescription: "Message sent",
		})
	}

	w.WriteHeader(200)
	json.NewEncoder(w).Encode(msgr)
}
