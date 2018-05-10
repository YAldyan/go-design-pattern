package main

import (
	"flag"
	"fmt"
	"log"
)

func main() {
	var output = flag.String("output", "console", "The output to use between 'console' and 'image' file")

	flag.Parse()

	var activeStrategy PrintStrategy
	switch *output {
	case "console":
		activeStrategy = &TextSquare{}
	case "image":
		activeStrategy = &ImageSquare{"/tmp/image.jpg"}
	default:
		activeStrategy = &TextSquare{}
	}

	err := activeStrategy.Print()
	if err != nil {
		log.Fatal(err)
	}
}
