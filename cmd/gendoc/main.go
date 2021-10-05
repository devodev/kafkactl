package main

import (
	"fmt"
	"os"

	"github.com/devodev/kafkactl/internal/kafkactl"
)

func checkError(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func main() {
	if len(os.Args) != 2 {
		checkError(fmt.Errorf("invalid number of args passed, dirpath not found"))
	}
	dirpath := os.Args[1]
	checkError(kafkactl.GenDoc(dirpath))
}
