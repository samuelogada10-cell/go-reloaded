package main

import "os"

func main() {
	a := os.Args
	if len(a) < 3 {
		return
	}

	d, _ := os.ReadFile(a[1])
	s := fixP(fixA(conv(string(d))))
	os.WriteFile(a[2], []byte(s), 0644)
}
