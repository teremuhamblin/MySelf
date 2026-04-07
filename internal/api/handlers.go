package api

import (
	"encoding/json"
	"myself/internal/analysis"
	"myself/internal/profile"
	"myself/pkg/data"
	"net/http"

	"github.com/gorilla/mux"
)

type Handlers struct {
	Store data.Store
}

func writeJSON(w http.ResponseWriter, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(v)
}

// ----------------------
// RECORDS
// ----------------------

func (h Handlers) GetRecords(w http.ResponseWriter, r *http.Request) {
	records, err := h.Store.LoadAllRecords()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	writeJSON(w, records)
}

func (h Handlers) GetRecordByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	record, err := h.Store.LoadRecordByID(id)
	if err != nil {
		http.Error(w, "Record not found", 404)
		return
	}

	writeJSON(w, record)
}

func (h Handlers) CreateRecord(w http.ResponseWriter, r *http.Request) {
	var input profile.Record
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid JSON", 400)
		return
	}

	if err := input.Validate(); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	if err := h.Store.SaveRecord(input); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	writeJSON(w, input)
}

// ----------------------
// STATS / TAGS / TRENDS
// ----------------------

func (h Handlers) GetStats(w http.ResponseWriter, r *http.Request) {
	records, _ := h.Store.LoadAllRecords()
	stats := analysis.ComputeStats(records)
	writeJSON(w, stats)
}

func (h Handlers) GetTags(w http.ResponseWriter, r *http.Request) {
	records, _ := h.Store.LoadAllRecords()
	stats := analysis.ComputeStats(records)
	writeJSON(w, stats.Tags)
}

func (h Handlers) GetTrends(w http.ResponseWriter, r *http.Request) {
	records, _ := h.Store.LoadAllRecords()
	stats := analysis.ComputeStats(records)
	writeJSON(w, stats.PerMonth)
}

// ----------------------
// ANALYSIS
// ----------------------

func (h Handlers) GetSummary(w http.ResponseWriter, r *http.Request) {
	records, _ := h.Store.LoadAllRecords()
	summary := analysis.GenerateSummary(records)
	writeJSON(w, summary)
}

func (h Handlers) GetPatterns(w http.ResponseWriter, r *http.Request) {
	records, _ := h.Store.LoadAllRecords()
	patterns := analysis.DetectPatterns(records)
	writeJSON(w, patterns)
}

func (h Handlers) GetMonthly(w http.ResponseWriter, r *http.Request) {
	records, _ := h.Store.LoadAllRecords()
	a := analysis.NewAnalyzer(records)
	writeJSON(w, a.Monthly())
}
