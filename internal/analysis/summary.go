package analysis

import (
	"fmt"
	"myself/internal/profile"
)

type SummaryResult struct {
	Text string `json:"text"`
}

func GenerateSummary(records []profile.Record) SummaryResult {
	if len(records) == 0 {
		return SummaryResult{Text: "Aucune donnée disponible."}
	}

	stats := ComputeStats(records)
	patterns := DetectPatterns(records)

	text := fmt.Sprintf(
		"Vous avez %d entrées. Le type dominant est '%s'. Les tags les plus fréquents sont : %v.",
		stats.Total,
		patterns.MostUsedType,
		patterns.TopTags,
	)

	return SummaryResult{Text: text}
}
