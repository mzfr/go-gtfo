package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/fatih/color"
	"gopkg.in/yaml.v2"
)

var rawBinURL = "https://raw.githubusercontent.com/GTFOBins/GTFOBins.github.io/master/_gtfobins/%s.md"

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

// Function to get the gtfobins yaml file and parse it
// for proper displaying on the screen
func gtfobins(binary string) {
	config := make(map[interface{}]interface{})

	// Format the URL and send the get request.
	binary_url := fmt.Sprintf(RAW_BIN_URL, binary)
	req, err := http.Get(binary_url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to create request: %s\n", err)
		return
	}

	defer req.Body.Close()

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return
	}

	// unmarshal for yaml
	if err = yaml.Unmarshal(body, &config); err != nil {
		fmt.Println("error: %v", err)
	}

	red := color.New(color.FgRed)
	boldRed := red.Add(color.Bold)
	yellow := color.New(color.FgYellow)
	boldYellow := yellow.Add(color.Bold)
	green := color.New(color.FgGreen).SprintFunc()
	// This is a weird for loop to get out the required
	// values out of the map[interface{}]interface{}
	for _, key := range config {
		for k, v := range key.(map[interface{}]interface{}) {
			details := v.([]interface{})[0].(map[interface{}]interface{})

			if details["description"] != nil {
				boldYellow.Println("# ", details["description"])
			}
			code := strings.ReplaceAll(fmt.Sprintf("%v", details["code"]), "\n", "\n\t")
			fmt.Printf("\n\nCode: \t %v \n", green(code))
			boldRed.Println("Type:\t", k)
			fmt.Println("\n")
		}
	}
}

//Main function
func main() {
	//define variables to hold flag value
	var bin string
	flag.StringVar(&bin, "bin", "", "")
	flag.StringVar(&bin, "b", "", "")

	var exe string
	flag.StringVar(&exe, "exe", "", "")
	flag.StringVar(&exe, "e", "", "")

	flag.Parse()
	// TODO: https://github.com/theckman/yacspin
	if bin != "" {
		gtfobins(bin)
	} else if exe != "" {
		fmt.Println("Windows sucks")
	} else {
		fmt.Println("No option selected")
		os.Exit(2)
	}
}
