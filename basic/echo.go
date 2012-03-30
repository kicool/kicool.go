package main

import (
	"os"
	"flag" // command line option parser
)

var omitNewLine = flag.Bool("n", false, "don't print final new line")

const (
	Space = " "
	Newline = "\n"
)

func main() {
	flag.Parse() // Scans the arg list and setups flags
	var s string = ""
	for i :=0; i < flag.NArg(); i++ {
		if i > 0 {
			s += Space
		}
		s += flag.Arg(i)
	}

	if !*omitNewLine {
		s += Newline
	}

	os.Stdout.WriteString(s)
}
