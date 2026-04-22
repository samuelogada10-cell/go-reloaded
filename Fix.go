package main

import "strings"

func fixA(s string) string {
	w := strings.Fields(s)
	for i := 0; i < len(w)-1; i++ {
		if (w[i] == "a" || w[i] == "A") && strings.ContainsAny(string(w[i+1][0]), "aeiouAEIOU") {
			w[i] += "n"
		}
	}
	return strings.Join(w, " ")
}

func fixP(s string) string {
	for _, p := range []string{".", ",", "!", "?", ":", ";"} {
		s = strings.ReplaceAll(s, " "+p, p)
	}
	return strings.ReplaceAll(s, " ' ", " '")
}
