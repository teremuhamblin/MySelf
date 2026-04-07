package main

import (
	"flag"
	"strings"
)

type AddFlags struct {
	Type    string
	Tags    []string
	Content string
}

func ParseAddFlags() AddFlags {
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)

	recordType := addCmd.String("type", "note", "Record type (emotion, habit, event, note)")
	tags := addCmd.String("tags", "", "Comma-separated list of tags")
	content := addCmd.String("content", "", "Record content")

	_ = addCmd.Parse(os.Args[2:])

	var tagList []string
	if *tags != "" {
		tagList = strings.Split(*tags, ",")
	}

	return AddFlags{
		Type:    *recordType,
		Tags:    tagList,
		Content: *content,
	}
}
