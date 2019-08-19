package status

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// SendStatus .
type SendStatus struct {
	GetSmsStatusResp `json:"getSmsStatusResp"`
}

// GetSmsStatusResp .
type GetSmsStatusResp struct {
	ID                 string `json:"id"`
	Received           string `json:"received"`
	Shortcode          int    `json:"shortcode"`
	MobileOperatorName string `json:"mobileOperatorName"`
	StatusCode         string `json:"statusCode"`
	StatusDescription  string `json:"statusDescription"`
	DetailCode         string `json:"detailCode"`
	DetailDescription  string `json:"detailDescription"`
}

// RegisterRouter .
func RegisterRouter(r *mux.Router) {
	r.HandleFunc("/services/get-sms-status/{id:[0-9]+}", statusHandler).Methods(http.MethodGet)
}

func statusHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	vars := mux.Vars(r)

	res := SendStatus{
		GetSmsStatusResp: GetSmsStatusResp{
			ID:                 vars["id"],
			Received:           "2014-08-23T02:01:23",
			Shortcode:          69788,
			MobileOperatorName: "claro",
			StatusCode:         "03",
			StatusDescription:  "Delivered",
			DetailCode:         "120",
			DetailDescription:  "Message received by mobile",
		},
	}

	json.NewEncoder(w).Encode(res)
}
