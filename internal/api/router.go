package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter(h Handlers) *mux.Router {
	r := mux.NewRouter()

	// Middlewares
	r.Use(CORSMiddleware)
	r.Use(LoggingMiddleware)

	// Routes
	r.HandleFunc("/records", h.GetRecords).Methods("GET")
	r.HandleFunc("/records/{id}", h.GetRecordByID).Methods("GET")
	r.HandleFunc("/records", h.CreateRecord).Methods("POST")

	r.HandleFunc("/stats", h.GetStats).Methods("GET")
	r.HandleFunc("/trends", h.GetTrends).Methods("GET")
	r.HandleFunc("/tags", h.GetTags).Methods("GET")

	r.HandleFunc("/analysis/summary", h.GetSummary).Methods("GET")
	r.HandleFunc("/analysis/patterns", h.GetPatterns).Methods("GET")
	r.HandleFunc("/analysis/monthly", h.GetMonthly).Methods("GET")

	// Health check
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	return r
}
