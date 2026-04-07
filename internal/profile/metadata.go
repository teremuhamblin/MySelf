package profile

import "time"

// Metadata contient des informations supplémentaires
// qui pourront être utilisées par l'analyse ou l'API.
type Metadata struct {
	WordCount int       `json:"word_count"`
	CharCount int       `json:"char_count"`
	DayOfWeek string    `json:"day_of_week"`
	Hour      int       `json:"hour"`
	CreatedAt time.Time `json:"created_at"`
}

// ComputeMetadata génère les métadonnées d'un record.
func ComputeMetadata(r Record) Metadata {
	return Metadata{
		WordCount: len(splitWords(r.Content)),
		CharCount: len(r.Content),
		DayOfWeek: r.CreatedAt.Weekday().String(),
		Hour:      r.CreatedAt.Hour(),
		CreatedAt: r.CreatedAt,
	}
}

func splitWords(s string) []string {
	var words []string
	current := ""

	for _, r := range s {
		if r == ' ' || r == '\n' || r == '\t' {
			if current != "" {
				words = append(words, current)
				current = ""
			}
		} else {
			current += string(r)
		}
	}

	if current != "" {
		words = append(words, current)
	}

	return words
}
