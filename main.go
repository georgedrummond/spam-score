package main

import (
	"encoding/json"
	"github.com/georgedrummond/spamscore/message"
	"github.com/gorilla/mux"
	"net/http"
	"html/template"
)

func HomepageHandler(res http.ResponseWriter, req *http.Request) {
	message := message.New()
	t, _ := template.ParseFiles("views/index.html")
	
	t.Execute(res, message)
}

func MessageNewHandler(res http.ResponseWriter, req *http.Request) {
	message := message.New()
	json, _ := json.Marshal(message)

	res.Header().Set("Content-Type", "application/json")
	res.Write(json)
}

func MessageViewHandler(res http.ResponseWriter, req *http.Request) {
	message := message.New()
	json, _ := json.Marshal(message)

	res.Header().Set("Content-Type", "application/json")
	res.Write(json)
}

func MessageStatsHandler(res http.ResponseWriter, req *http.Request) {
	stats := message.Stats()
	json, _ := json.Marshal(stats)

	res.Header().Set("Content-Type", "application/json")
	res.Write(json)
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", HomepageHandler).Methods("GET")
	r.HandleFunc("/api/v1/messages/new", MessageNewHandler).Methods("GET")
	r.HandleFunc("/api/v1/messages/stats", MessageStatsHandler).Methods("GET")
	r.HandleFunc("/api/v1/messages/{uuid}", MessageViewHandler).Methods("GET")

	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}
