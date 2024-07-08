package main

import (
	"strings"
)

func split_white_spaces(s string) []string {
	open_quote := false
	split := []string{}

	if len(s) == 0 {
		return split
	}

	for s[0] == ' ' || s[0] == '\n' || s[0] == '\t' || s[0] == '\r' || s[0] == '\v' || s[0] == '\f' {
		s = s[1:]
	}

	for i := 0; i < len(s); i++ {
		if s[i] == '"' && (i == 0 || s[i-1] != '\\') {
			open_quote = !open_quote
		}
		if (s[i] == ' ' || s[i] == '\n' || s[i] == '\t' || s[i] == '\r' || s[i] == '\v' || s[i] == '\f') && !open_quote {
			for j := i; j < len(s); j++ {
				if s[j] != ' ' && s[j] != '\n' && s[j] != '\t' && s[j] != '\r' && s[j] != '\v' && s[j] != '\f' {
					split = append(split, s[:i])
					s = s[j:]
					i = -1
					break
				}
			}
		}
		if i == len(s)-1 {
			split = append(split, s)
		}
	}

	split_len := len(split)
	split_str_len := len(split[split_len-1])
	char := split[split_len-1][split_str_len-1]
	for char == ' ' || char == '\n' || char == '\t' || char == '\r' || char == '\v' || char == '\f' {
		split[split_len-1] = split[split_len-1][:split_str_len-1]
		split_len = len(split)
		split_str_len = len(split[split_len-1])
		char = split[split_len-1][split_str_len-1]
	}

	return split
}

func split_new_lines(s string) []string {
	return strings.Split(strings.ReplaceAll(s, "\r", ""), "\n")
}

func parse_file(s string) [][]string {
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
