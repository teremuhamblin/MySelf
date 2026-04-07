package main

import (
	"fmt"
	"os"
)

func PrintAvailableCommands() {
	fmt.Println("\nAvailable commands:")
	fmt.Println("  add     Add a new record")
	fmt.Println("  list    List all records")
	fmt.Println("  view    View a specific record by ID")
}

func RunAddCommand() {
	flags := ParseAddFlags()

	content := flags.Content
	if content == "" {
		content = ReadContentFromInput()
	}

	record := Record{
		ID:        GenerateID(),
		Type:      flags.Type,
		Tags:      flags.Tags,
		Content:   content,
		CreatedAt: Now(),
	}

	if err := ValidateRecord(record); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	if err := SaveRecord(record); err != nil {
		fmt.Println("Error saving record:", err)
		os.Exit(1)
	}

	fmt.Println("Record saved with ID:", record.ID)
}

func RunListCommand() {
	records, err := LoadAllRecords()
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	for _, r := range records {
		fmt.Printf("[%s] %s (%s)\n", r.ID, r.Type, r.CreatedAt)
	}
}

func RunViewCommand() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: myself view <id>")
		os.Exit(1)
	}

	id := os.Args[2]

	record, err := LoadRecordByID(id)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Printf("ID: %s\nType: %s\nTags: %v\nDate: %s\n\n%s\n",
		record.ID, record.Type, record.Tags, record.CreatedAt, record.Content)
}
