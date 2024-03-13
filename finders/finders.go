package finders

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Start() (data []string, err error) {
	var store []string
	fmt.Println("Func started")
	fmt.Println("The main purpose of this programme to find your file in OS")
	fmt.Println("Right now you will need to write file name that u want to find")

	reader := bufio.NewReader(os.Stdin)

	for i := 0; i < 2; i++ {
		if i == 0 {
			fmt.Print("Filename: ")
			input, err := reader.ReadString('\n')
			input = strings.TrimSpace(input)

			if err != nil {
				return nil, fmt.Errorf("You have typed incorrect file")
			}

			store = append(store, input)
		} else {
			fmt.Println("Now write the directory that your file locates")
			fmt.Print("DirName: ")
			input, err := reader.ReadString('\n')
			input = strings.TrimSpace(input)

			if err != nil {
				return nil, fmt.Errorf("You have typed incorrect directory")
			}

			store = append(store, input)
		}
	}

	return store, nil
}
