package main

import (
	"fmt"
)

var calls = map[string]func([]stack) []stack{
	"print": func(s []stack) (res []stack) {
		msg := ""
		for _, v := range s {
			msg = fmt.Sprint(msg, v.data)
		}
		fmt.Print(msg)

		return res
	},
}
