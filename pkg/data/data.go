package data

import (
	"encoding/json"
	"fmt"
	"myself/internal/profile"
	"os"
	"path/filepath"
	"time"
)

type Store struct {
	BasePath string
}

func NewStore(basePath string) Store {
	return Store{BasePath: basePath}
}

// SaveRecord enregistre un record dans records/YYYY/MM/ID.json
func (s Store) SaveRecord(r profile.Record) error {
	year := r.CreatedAt.Format("2006")
	month := r.CreatedAt.Format("01")

	dir := filepath.Join(s.BasePath, year, month)
	os.MkdirAll(dir, 0755)

	path := filepath.Join(dir, r.ID+".json")

	data, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0644)
}

// LoadAllRecords charge tous les records du disque
func (s Store) LoadAllRecords() ([]profile.Record, error) {
	var records []profile.Record

	err := filepath.Walk(s.BasePath, func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) == ".json" && filepath.Base(path) != "index.json" {
			data, err := os.ReadFile(path)
			if err != nil {
				return err
			}

			var r profile.Record
			if err := json.Unmarshal(data, &r); err != nil {
				return err
			}

			records = append(records, r)
		}
		return nil
	})

	return records, err
}

// LoadRecordByID recherche un record par ID
func (s Store) LoadRecordByID(id string) (profile.Record, error) {
	records, err := s.LoadAllRecords()
	if err != nil {
		return profile.Record{}, err
	}

	for _, r := range records {
		if r.ID == id {
			return r, nil
		}
	}

	return profile.Record{}, fmt.Errorf("record not found: %s", id)
}

// UpdateRecord met à jour un record existant
func (s Store) UpdateRecord(updated profile.Record) error {
	updated.UpdatedAt = time.Now()
	return s.SaveRecord(updated)
}
