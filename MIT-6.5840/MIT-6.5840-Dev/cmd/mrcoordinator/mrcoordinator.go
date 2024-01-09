package main

//
// start the coordinator process, which is implemented
// in ../mr/coordinator.go
//
// go run mrcoordinator.go pg*.txt
//
// Please do not change this file.
//

import (
	"fmt"
	"os"
	"time"

	"6.5840/mr"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: mrcoordinator inputfiles...\n")
		os.Exit(1)
	}

	// cmd: go run mrcoordinator.go pg-*.txt
	// os.Args["mrcoordinator", "pg-*.txt"]
	m := mr.MakeCoordinator(os.Args[1:], 10)
	for m.Done() == true {
		time.Sleep(time.Second)
	}

	time.Sleep(time.Second)
}
