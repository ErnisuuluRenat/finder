package main

import (
	"finder/finders"
	"finder/logic"
	"fmt"
	"log"
)

func main() {
	fmt.Println("Hello boss")
	data, err := finders.Start()

	if err != nil {
		log.Fatal(err)
	}

	fileFound, err := logic.FindFile(data...)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Yeah, file was found at: ", fileFound)
}
