package analysis

import (
	"myself/internal/profile"
	"time"
)

type Analyzer struct {
	Records []profile.Record
}

func NewAnalyzer(records []profile.Record) Analyzer {
	return Analyzer{Records: records}
}

func (a Analyzer) Stats() StatsResult {
	return ComputeStats(a.Records)
}

func (a Analyzer) Patterns() PatternsResult {
	return DetectPatterns(a.Records)
}

func (a Analyzer) Summary() SummaryResult {
	return GenerateSummary(a.Records)
}

func (a Analyzer) Monthly() map[string]StatsResult {
	result := make(map[string]StatsResult)

	for _, r := range a.Records {
		key := r.CreatedAt.Format("2006-01")
		result[key] = ComputeStats(filterByMonth(a.Records, r.CreatedAt))
	}

	return result
}

func filterByMonth(records []profile.Record, t time.Time) []profile.Record {
	var out []profile.Record
	for _, r := range records {
		if r.CreatedAt.Year() == t.Year() && r.CreatedAt.Month() == t.Month() {
			out = append(out, r)
		}
	}
	return out
}
