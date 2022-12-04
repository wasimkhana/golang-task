package main

import (
	"fmt"
	"io"
	"net/http"
	"webserver/handler"
)

func main() {

	http.HandleFunc("/", MainHandler)
	http.HandleFunc("/hai/", handler.GetHaiHandler)
	http.HandleFunc("/haicounter/", handler.GetHaiCounterHandler)
	fmt.Println("starting web server...accessible at 127.0.0.1:3333")
	http.ListenAndServe(":3333", nil) //you can reach the server at 127.0.0.1:3333
}

func MainHandler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path == "/" && r.Method == "GET" {
		fmt.Println("request received ")
		message := fmt.Sprintf("Hello World!")
		io.WriteString(w, message)
	} else if r.URL.Path != "/" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"error": "Page Not Found"}`))
		//http.Error(w, "404 Not Found", http.StatusNotFound)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(`{"error": "Method Not Allowed"}`))
	}
}
