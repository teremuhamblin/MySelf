package data

import "myself/internal/profile"

type Aggregation struct {
	CountByType map[string]int `json:"count_by_type"`
	CountByTag  map[string]int `json:"count_by_tag"`
}

func Aggregate(records []profile.Record) Aggregation {
	agg := Aggregation{
		CountByType: make(map[string]int),
		CountByTag:  make(map[string]int),
	}

	for _, r := range records {
		agg.CountByType[r.Type]++

		for _, tag := range r.Tags {
			agg.CountByTag[tag]++
		}
	}

	return agg
}
