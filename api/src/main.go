package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"BF/src/database"
	"BF/src/report"
)

func main() {
	database.Connect()

	r := mux.NewRouter()

	r.HandleFunc("/report", report.ReportHandler).Methods("POST")
	r.HandleFunc("/reports", report.GetReportsHandler).Methods("GET", "OPTIONS")

	log.Println("server started on :8080")
	if err := http.ListenAndServe(":8080", withCORS(r)); err != nil {
		log.Println("failed to start server")
	}
}

func withCORS(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		h.ServeHTTP(w, r)
	})
}
