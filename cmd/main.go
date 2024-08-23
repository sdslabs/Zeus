package main

import (
	"os"

	"github.com/sdslabs/Zeus/pkg/initialize"
)

func main() {
	newrootPath := os.Args[1]
	initialize.Begin(newrootPath)
}
