package main

import (
	"fmt"
)

var calls = map[string]func([]stack) []stack{
	"print": func(s []stack) (res []stack) {
		msg := ""
		for i := range s {
			msg = fmt.Sprint(msg, s[i].data)
		}
		fmt.Print(msg)

		return res
	},
}
