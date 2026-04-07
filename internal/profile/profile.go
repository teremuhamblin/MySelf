package profile

import (
	"errors"
	"time"
)

// Record est la structure centrale de MySelf.
// Elle est utilisée par :
// - le collecteur (A/B/C)
// - l'analyse (1.2.0)
// - l'API REST (1.3.0)
// - le dashboard web (1.3.0)
type Record struct {
	ID        string    `json:"id"`
	Type      string    `json:"type"`
	Tags      []string  `json:"tags"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Version   string    `json:"version"` // version du format (utile pour 1.4.0)
}

// NewRecord crée un record propre et complet.
func NewRecord(id, rtype, content string, tags []string) Record {
	now := time.Now()

	return Record{
		ID:        id,
		Type:      rtype,
		Tags:      tags,
		Content:   content,
		CreatedAt: now,
		UpdatedAt: now,
		Version:   "1.0",
	}
}

// Validate vérifie que le record est cohérent.
func (r Record) Validate() error {
	if r.ID == "" {
		return errors.New("record ID cannot be empty")
	}
	if r.Type == "" {
		return errors.New("record type cannot be empty")
	}
	if r.Content == "" {
		return errors.New("record content cannot be empty")
	}
	return nil
}

// UpdateContent modifie le contenu et met à jour la date.
func (r *Record) UpdateContent(newContent string) {
	r.Content = newContent
	r.UpdatedAt = time.Now()
}

// UpdateTags remplace les tags et met à jour la date.
func (r *Record) UpdateTags(tags []string) {
	r.Tags = tags
	r.UpdatedAt = time.Now()
}

// UpdateType change le type du record.
func (r *Record) UpdateType(t string) {
	r.Type = t
	r.UpdatedAt = time.Now()
}
