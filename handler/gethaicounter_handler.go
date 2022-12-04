package handler

import (
	"net/http"
	"webserver/service"
)

func GetHaiCounterHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.URL.Path != "/haicounter/" {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"error": "Page Not Found"}`))

	} else if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(`{"error": "Method Not Allowed"}`))
	} else {
		hais, _ := service.GetHaisCounterService()
		w.WriteHeader(http.StatusOK)
		w.Write(hais)
	}

}
