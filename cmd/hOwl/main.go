package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	flag.Parse()
	if flag.NArg() == 0 {
		flag.Usage()
		os.Exit(1)
	}

	location := flag.Arg(0)
	fmt.Sprintf("%s", location)
}
