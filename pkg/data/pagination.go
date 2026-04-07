package data

import "myself/internal/profile"

func Paginate(records []profile.Record, limit, offset int) []profile.Record {
	if limit <= 0 {
		return records
	}

	if offset >= len(records) {
		return []profile.Record{}
	}

	end := offset + limit
	if end > len(records) {
		end = len(records)
	}

	return records[offset:end]
}
