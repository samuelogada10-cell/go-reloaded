package main

import "strings"

func fixA(s string) string {
	w := strings.Fields(s)

	for i := 0; i < len(w)-1; i++ {
		if w[i] == "a" || w[i] == "A" {
			next := strings.ToLower(w[i+1])

			if strings.ContainsAny(string(next[0]), "aeiou") {
				w[i] += "n"
				continue
			}

			if strings.HasPrefix(next, "honest") ||
				strings.HasPrefix(next, "hour") ||
				strings.HasPrefix(next, "heir") ||
				strings.HasPrefix(next, "honor") {
				w[i] += "n"
			}
		}
	}

	return strings.Join(w, " ")
}

func fixP(s string) string {
	for _, p := range []string{".", ",", "!", "?", ":", ";"} {
		s = strings.ReplaceAll(s, " "+p, p)
	}
	return s
}
