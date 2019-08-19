package cancel

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// SmsCacelResp .
type SmsCacelResp struct {
	SmsResp `json:"cancelSmsResp"`
}

// SmsResp .
type SmsResp struct {
	StatusCode        string `json:"statusCode"`
	StatusDescription string `json:"statusDescription"`
	DetailCode        string `json:"detailCode"`
	DetailDescription string `json:"detailDescription"`
}

// RegisterRouter .
func RegisterRouter(r *mux.Router) {
	r.HandleFunc("/services/cancel-sms/{id:[0-9]+}", cancelHandler).Methods(http.MethodPost)
}

func cancelHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	res := SmsCacelResp{
		SmsResp: SmsResp{
			StatusCode:        "09",
			StatusDescription: "Blocked",
			DetailCode:        "002",
			DetailDescription: "Message successfully canceled",
		},
	}
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(res)
}
