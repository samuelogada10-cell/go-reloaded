package main

import (
	"strconv"
	"strings"
)

func conv(s string) string {
	s = strings.ReplaceAll(s, "(", " (")
	s = strings.ReplaceAll(s, ")", ") ")

	w := strings.Fields(s)
	var res []string

	for i := 0; i < len(w); i++ {
		v := w[i]

		if strings.HasPrefix(v, "(") {
			n := 1

			if strings.Contains(v, ",") && i+1 < len(w) {
				num := strings.TrimRight(w[i+1], ")")
				n, _ = strconv.Atoi(num)
				i++
			}

			for j := 1; j <= n && len(res)-j >= 0; j++ {
				idx := len(res) - j

				if strings.HasPrefix(v, "(up") {
					res[idx] = strings.ToUpper(res[idx])
				}
				if strings.HasPrefix(v, "(low") {
					res[idx] = strings.ToLower(res[idx])
				}
				if strings.HasPrefix(v, "(cap") {
					res[idx] = strings.Title(res[idx])
				}
				if v == "(hex)" {
					x, _ := strconv.ParseInt(res[idx], 16, 64)
					res[idx] = strconv.FormatInt(x, 10)
				}
				if v == "(bin)" {
					x, _ := strconv.ParseInt(res[idx], 2, 64)
					res[idx] = strconv.FormatInt(x, 10)
				}
			}
			continue
		}

		res = append(res, v)
	}

	return strings.Join(res, " ")
}
