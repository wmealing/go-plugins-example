package main

import (
	"fmt"
	"log"
	"os"

	"github.com/sourcegraph/go-diff/diff"
)

type greeting string

func (g greeting) Greet(d diff.FileDiff) {
	fmt.Println("Hello Universe")
	printed, err := diff.PrintFileDiff(&d)

	if err != nil {
		fmt.Println("wat")

	}
	if _, err := os.Stdout.Write(printed); err != nil {
		log.Fatal(err)
	}

}

// exported
var Greeter greeting
