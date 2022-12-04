package handler

import (
	"fmt"
	"net/http"
	"webserver/service"
)

type Hais struct {
	ID   int    `json:"h_id,omitempty"`
	City string `json:"city,omitempty"`
	Hai  string `json:"hai,omitempty"`
}

func GetHaiHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if len(r.URL.Query()) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf(`{"error":"missing params"}`)))
		return
	}

	if r.URL.Path != "/hai/" {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"error": "Page Not Found"}`))
		return

	} else if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(`{"error": "Method Not Allowed"}`))
	} else {
		hais, err := service.GetHaisData(r.URL.Query())
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(fmt.Sprintf(`{"error":"Details not Found"}`)))
			return
		}

		result := []byte(fmt.Sprintf(`{"result":%s}`, hais))
		w.WriteHeader(http.StatusOK)
		w.Write(result)
	}

}
