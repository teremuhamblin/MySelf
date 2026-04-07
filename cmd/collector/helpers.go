package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

type Record struct {
	ID        string    `json:"id"`
	Type      string    `json:"type"`
	Tags      []string  `json:"tags"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func GenerateID() string {
	return fmt.Sprintf("%d", time.Now().UnixNano())
}

func Now() time.Time {
	return time.Now()
}

func SaveRecord(r Record) error {
	year := r.CreatedAt.Format("2006")
	month := r.CreatedAt.Format("01")

	dir := filepath.Join("records", year, month)
	os.MkdirAll(dir, 0755)

	path := filepath.Join(dir, r.ID+".json")

	data, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0644)
}

func LoadAllRecords() ([]Record, error) {
	var records []Record

	err := filepath.Walk("records", func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) == ".json" {
			data, err := os.ReadFile(path)
			if err != nil {
				return err
			}

			var r Record
			if err := json.Unmarshal(data, &r); err != nil {
				return err
			}

			records = append(records, r)
		}
		return nil
	})

	return records, err
}

func LoadRecordByID(id string) (Record, error) {
	records, err := LoadAllRecords()
	if err != nil {
		return Record{}, err
	}

	for _, r := range records {
		if r.ID == id {
			return r, nil
		}
	}

	return Record{}, fmt.Errorf("record not found: %s", id)
}
