package main

import (
	"os"
	"fmt"
	"flag"
	"strings"
)

func init() {
	flag.Usage = func() {
		h := []string{
			"Search gtfobin and lolbas files from terminal",
			"",
			"Options:",
			"  -b, --bin <binary>       Search Linux binaries on gtfobins",
			"  -e, --exe <EXE>       	Search Windows EXE on lolbas",
			"",
		}

		fmt.Fprintf(os.Stderr, strings.Join(h, "\n"))
	}
}

func main() {
	var bin string
	flag.StringVar(&bin, "bin", "", "")
	flag.StringVar(&bin, "b", "", "")

	var exe string
	flag.StringVar(&exe, "exe", "", "")
	flag.StringVar(&exe, "e", "", "")

	flag.Parse()

	if bin != ""  {
		fmt.Println("Binaries")
		gtfobins(bin)
	} else if exe != "" {
		fmt.Println("Windows sucks")
	} else {
		fmt.Println("No option selected")
		os.Exit(2)
	}
}
