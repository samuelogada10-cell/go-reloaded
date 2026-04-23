package main

import "strings"

func fixA(s string) string {
	w := strings.Fields(s)

	for i := 0; i < len(w)-1; i++ {
		word := w[i]
		next := strings.ToLower(w[i+1])

		useAn := false

		// vowel words
		if strings.ContainsAny(string(next[0]), "aeiou") {
			useAn = true
		}

		// words that sound like "you"
		if strings.HasPrefix(next, "uni") ||
			strings.HasPrefix(next, "use") ||
			strings.HasPrefix(next, "euro") {
			useAn = false
		}

		// silent h
		if next == "hour" ||
			next == "honest" ||
			next == "honor" ||
			next == "honour" ||
			next == "heir" {
			useAn = true
		}

		if word == "a" || word == "an" {
			if useAn {
				w[i] = "an"
			} else {
				w[i] = "a"
			}
		}

		if word == "A" || word == "An" {
			if useAn {
				w[i] = "An"
			} else {
				w[i] = "A"
			}
		}
	}

	return strings.Join(w, " ")
}
