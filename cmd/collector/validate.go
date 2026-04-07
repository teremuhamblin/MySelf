package main

import "errors"

func ValidateRecord(r Record) error {
	if r.Type == "" {
		return errors.New("record type cannot be empty")
	}

	if r.Content == "" {
		return errors.New("record content cannot be empty")
	}

	return nil
}
