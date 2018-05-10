package main

import (
	"fmt"
	"strings"
)

func main() {
	toUpperSync("Hello Callbacks!", func(v string) {
		fmt.Printf("Callback: %s\n", v)
	})
}

func toUpperSync(word string, f func(string)) {

	/*
		1. String Word dijadikan uppercase.
		2. Hasil Uppercase tersebut dijadikan
		   parameter yang dikirimkan ke fungsi
		   yang ada di main program untuk di-
		   cetak ke layar.
	*/

	/*
		Variabel f dibawah menandakan bahwa
		itu adalah func sebagaimana paremeter
		pada fungsi toUpperSync, dan fungsi
		string berikutnya adalah passing param
		untuk fungsi anonymous yang ada di main
		program
	*/
	f(strings.ToUpper(word))
}
