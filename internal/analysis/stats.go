package analysis

import (
	"myself/internal/profile"
)

type StatsResult struct {
	Total       int               `json:"total"`
	ByType      map[string]int    `json:"by_type"`
	Tags        map[string]int    `json:"tags"`
	PerMonth    map[string]int    `json:"per_month"`
}

func ComputeStats(records []profile.Record) StatsResult {
	stats := StatsResult{
		Total:    len(records),
		ByType:   make(map[string]int),
		Tags:     make(map[string]int),
		PerMonth: make(map[string]int),
	}

	for _, r := range records {
		stats.ByType[r.Type]++
		stats.PerMonth[r.CreatedAt.Format("2006-01")]++

		for _, tag := range r.Tags {
			stats.Tags[tag]++
		}
	}

	return stats
}
