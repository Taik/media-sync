package main

import (
	"flag"
	"fmt"
)

func main() {
	hostFlag := flag.String("host", "", "Host to sync to")
	flag.Parse()

	if *hostFlag == "" {
		fmt.Println("Please provide the host to sync to")
		return
	}
}
