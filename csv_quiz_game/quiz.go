package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func main() {
	problems := os.Args[1]
	file, err := os.Open(problems)
	if err != nil {
		fmt.Println(err)
		return
	}

	qa := csv.NewReader(file)

	var r []string
	fmt.Println("----")
	for err != io.EOF {
		r, err = qa.Read()
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("What is %v?\n", r[0])
		var n string
		fmt.Scanf("%s", &n)
		if n == r[1] {
			fmt.Println("Correct!")
		} else {
			fmt.Println("Incorrect.")
		}
		fmt.Println("----")
	}
}
