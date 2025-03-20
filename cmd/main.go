package main

import (
	"fmt"
	"os"

	"github.com/sdslabs/Zeus/pkg/initialize"
)

/* "runc" is the low-level runtime for Zeus.
It is responsible for setting up the container environment and running the container. */

func main() {
	//version checking
	fmt.Println("Zeus runz v0.1")
	// path to the root filesystem
	newrootPath := "/tmp/newroot"
	//checks if the number of arguments is correct else, gives a usage message
	if len(os.Args) == 2 {
		newrootPath = os.Args[1] // path to the root filesystem
		memoryLimit := "512M"     // max memory limit
		maxPid := "10000"         //  max processes

		// The begin function has been separated from main to make multi-container execution easier in the future
		initialize.Begin(newrootPath, memoryLimit, maxPid)
	} else if len(os.Args) == 4 {
		newrootPath = os.Args[1]
		// memory limit for the container in MBs
		memoryLimit := os.Args[2]
		// maximum number of processes that can be created in the container
		maxPid := os.Args[3]

		initialize.Begin(newrootPath, memoryLimit, maxPid)
	} else {
		fmt.Println("Usage: zeus <newrootPath> <memoryLimit in xM format> <maxPid>")
	}
}
