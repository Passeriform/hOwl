package main

import (
	"flag"
	"os"
)

func main() {
	flag.Parse()
	if flag.NArg() == 0 {
		flag.Usage()
		os.Exit(1)
	}

	location := flag.Arg(0)

}
