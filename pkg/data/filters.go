package data

import (
	"myself/internal/profile"
	"strings"
	"time"
)

func FilterByType(records []profile.Record, t string) []profile.Record {
	if t == "" {
		return records
	}

	var out []profile.Record
	for _, r := range records {
		if r.Type == t {
			out = append(out, r)
		}
	}
	return out
}

func FilterByTag(records []profile.Record, tag string) []profile.Record {
	if tag == "" {
		return records
	}

	var out []profile.Record
	for _, r := range records {
		for _, t := range r.Tags {
			if t == tag {
				out = append(out, r)
				break
			}
		}
	}
	return out
}

func FilterByDateRange(records []profile.Record, start, end time.Time) []profile.Record {
	var out []profile.Record
	for _, r := range records {
		if r.CreatedAt.After(start) && r.CreatedAt.Before(end) {
			out = append(out, r)
		}
	}
	return out
}

func FilterByText(records []profile.Record, text string) []profile.Record {
	if text == "" {
		return records
	}

	var out []profile.Record
	for _, r := range records {
		if strings.Contains(strings.ToLower(r.Content), strings.ToLower(text)) {
			out = append(out, r)
		}
	}
	return out
}
