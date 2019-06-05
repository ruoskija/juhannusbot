package jbot

import "strings"

func stringHasAnyPrefix(s string, prefixes []string) bool {
	for _, prefix := range prefixes {
		if strings.HasPrefix(strings.ToLower(s), prefix) {
			return true
		}
	}
	return false
}
