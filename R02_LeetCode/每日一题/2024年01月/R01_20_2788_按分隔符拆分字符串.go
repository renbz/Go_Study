package main

import "strings"

func splitWordsBySeparator(words []string, separator byte) (ans []string) {

	for _, word := range words {
		split := strings.Split(word, string(separator))
		for _, s := range split {
			if len(s) > 0 {
				ans = append(ans, s)
			}
		}
	}
	return ans
}
