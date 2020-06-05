package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/common-nighthawk/go-figure"
	"github.com/fatih/color"
	"gopkg.in/yaml.v2"
)

var rawBinURL = "https://raw.githubusercontent.com/GTFOBins/GTFOBins.github.io/master/_gtfobins/%s.md"

func init() {
	flag.Usage = func() {
		h := []string{
			"Search gtfobin from terminal",
			"",
			"Options:",
			"  -b, --bin <binary>       Search Linux binaries on gtfobins",
			"",
		}

		fmt.Fprintf(os.Stderr, strings.Join(h, "\n"))
	}
}

// Function to get the gtfobins yaml file and parse it
// for proper displaying on the screen
func gtfobins(binary string) {
	config := make(map[interface{}]interface{})

	// Format the URL and send the get request.
	binaryURL := fmt.Sprintf(rawBinURL, binary)

	req, err := http.Get(binaryURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to create request: %s\n", err)
		return
	}

	defer req.Body.Close()

	// Just incase someone entered some random name
	if req.StatusCode == 404 {
		color.Red("[!] Binary not found on GTFObins")
		return
	}

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return
	}
	if err = yaml.Unmarshal(body, &config); err != nil {
		fmt.Println(err)
	}

	yellow := color.New(color.FgYellow)
	boldYellow := yellow.Add(color.Bold)
	green := color.New(color.FgGreen).SprintFunc()
	magenta := color.New(color.FgHiMagenta).SprintFunc()

	// This is a weird for loop to get out the required
	// values out of the map[interface{}]interface{}
	for _, key := range config {
		// Use switch case because some have direct strings
		// and some yamls have more information.
		switch key.(type) {
		case map[interface{}]interface{}:
			for k, v := range key.(map[interface{}]interface{}) {
				details := v.([]interface{})[0].(map[interface{}]interface{})

				// This is so that all the code section start from the same point.
				code := strings.ReplaceAll(fmt.Sprintf("%v", details["code"]), "\n", "\n\t")

				// Just formatting and printing.
				if details["description"] != nil {
					boldYellow.Println("\n# ", details["description"])
				}
				fmt.Printf("Code:\t%v \n", green(code))
				fmt.Printf("Type:\t%v\n", magenta(k))
				fmt.Println()
			}
		case string:
			boldYellow.Println("\n# ", key)

		}

	}
}

func main() {
	var bin string
	flag.StringVar(&bin, "bin", "", "")
	flag.StringVar(&bin, "b", "", "")

	flag.Parse()
	myFigure := figure.NewColorFigure("# gtfo", "big", "green", true)
	myFigure.Print()

	gtfobins(bin)
}
