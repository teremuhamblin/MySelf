package analysis

import (
	"myself/internal/profile"
)

type PatternsResult struct {
	MostUsedType string   `json:"most_used_type"`
	TopTags      []string `json:"top_tags"`
}

func DetectPatterns(records []profile.Record) PatternsResult {
	stats := ComputeStats(records)

	// Type dominant
	var dominantType string
	var maxCount int
	for t, c := range stats.ByType {
		if c > maxCount {
			maxCount = c
			dominantType = t
		}
	}

	// Top tags
	topTags := topN(stats.Tags, 3)

	return PatternsResult{
		MostUsedType: dominantType,
		TopTags:      topTags,
	}
}

func topN(m map[string]int, n int) []string {
	type kv struct {
		Key   string
		Value int
	}

	var arr []kv
	for k, v := range m {
		arr = append(arr, kv{k, v})
	}

	// tri simple
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[j].Value > arr[i].Value {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}

	var out []string
	for i := 0; i < len(arr) && i < n; i++ {
		out = append(out, arr[i].Key)
	}

	return out
}
