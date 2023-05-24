package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
	"time"
)

func timer(c *int, d int) {
	time.Sleep(time.Duration(d) * time.Second)
	fmt.Println("----")
	fmt.Println("Time's up!")
	fmt.Printf("You answered %v questions correctly.\n", *c)
	os.Exit(0)
}

func main() {
	var problems *string = flag.String(
		"csv",
		"problems.csv",
		"Provide a csv file with two columns: 'question, answer'",
	)
	var duration *int = flag.Int("duration", 30, "Provide the time duration for the quiz.")
	flag.Parse()

	file, err := os.Open(*problems)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var qa *csv.Reader = csv.NewReader(file)
	var r []string

	fmt.Println("----")
	var count int = 0
	go timer(&count, *duration)
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
			count++
		} else {
			fmt.Println("Incorrect.")
		}

		fmt.Println("----")
	}
}
