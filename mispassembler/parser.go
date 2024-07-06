package main

import (
	"strings"
)

func split_white_spaces(s string) []string {
	return strings.Fields(s)
}

func split_new_lines(s string) []string {
	return strings.Split(strings.ReplaceAll(s, "\r", ""), "\n")
}

func split_file(s string) [][]string {
	final := [][]string{}

	lines := split_new_lines(s)
	for _, line := range lines {
		split := split_white_spaces(line)
		if len(split) == 0 {
			continue
		}
		final = append(final, split)
	}

	return final
}
