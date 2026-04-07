package profile

import (
	"encoding/json"
)

// ToJSON convertit un record en JSON.
func (r Record) ToJSON() ([]byte, error) {
	return json.MarshalIndent(r, "", "  ")
}

// FromJSON convertit du JSON en record.
func FromJSON(data []byte) (Record, error) {
	var r Record
	err := json.Unmarshal(data, &r)
	return r, err
}

// ToMap convertit un record en map[string]interface{}.
// Utile pour l'API REST ou les templates.
func (r Record) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"id":         r.ID,
		"type":       r.Type,
		"tags":       r.Tags,
		"content":    r.Content,
		"created_at": r.CreatedAt,
		"updated_at": r.UpdatedAt,
		"version":    r.Version,
	}
}
