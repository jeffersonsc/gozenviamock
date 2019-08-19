package single

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// SMSRequest .
type SMSRequest struct {
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

// SMSResponse .
type SMSResponse struct {
	SendSmsResponse `json:"sendSmsResponse"`
}

// SendSmsResponse .
type SendSmsResponse struct {
	StatusCode        string `json:"statusCode"`
	StatusDescription string `json:"statusDescription"`
	DetailCode        string `json:"detailCode"`
	DetailDescription string `json:"detailDescription"`
}

// RegisterRouter .
func RegisterRouter(r *mux.Router) {
	r.HandleFunc("/services/send-sms", singleHandler).Methods(http.MethodPost)
}

func singleHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	msg := SMSRequest{}
	if err := json.NewDecoder(r.Body).Decode(&msg); err != nil {
		log.Println("[singleHandler] ERROR Failed decode single message ", err.Error())
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(map[string]string{"error": "failed decode json"})
		return
	}

	msgr := SMSResponse{
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
