package main

import (
	"os"

	"github.com/devodev/kafkactl/internal/kafkactl"
)

func main() {
	if err := kafkactl.Execute(); err != nil {
		os.Exit(1)
	}
}
