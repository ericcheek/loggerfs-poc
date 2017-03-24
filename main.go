package main


import (
	"fmt"
	"flag"
	"loggingfs"
	"os"
	)

func main() {
	flag.Parse()
	if flag.NArg() < 1 {
		fmt.Println("usage: loggerfs <config file>")
		os.Exit(2)
	}
	
	loggingfs.SetupAndMount(flag.Arg(0))
}