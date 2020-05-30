package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/fatih/color"
	"gopkg.in/yaml.v2"
)

var rawBinURL = "https://raw.githubusercontent.com/GTFOBins/GTFOBins.github.io/master/_gtfobins/%s.md"
var rawExeURL = "https://raw.githubusercontent.com/LOLBAS-Project/LOLBAS-Project.github.io/master/_lolbas/%s.md"

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
		for k, v := range key.(map[interface{}]interface{}) {
			details := v.([]interface{})[0].(map[interface{}]interface{})

			// This is so that all the code section start from the same point.
			code := strings.ReplaceAll(fmt.Sprintf("%v", details["code"]), "\n", "\n\t")

			// Just formatting and printing.
			if details["description"] != nil {
				boldYellow.Println("# ", details["description"])
			}
			fmt.Printf("\nCode:\t%v \n", green(code))
			fmt.Printf("Type:\t%v\n", magenta(k))
			fmt.Println()
		}
	}
}

func lolbas(exe string) {
	mapExe := make(map[string]string)

	doc, err := goquery.NewDocument("https://lolbas-project.github.io/")
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to create request: %s\n", err)
		return
	}
	doc.Find(".bin-name").Each(func(index int, item *goquery.Selection) {
		href, _ := item.Attr("href")
		mapExe[item.Text()] = strings.TrimSuffix(href[8:], "/")
	})

	exeURL := fmt.Sprintf(rawExeURL, mapExe[exe])
	fmt.Println(exeURL)
	fmt.Println(mapExe[exe])
	req, err := http.Get(exeURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to create request: %s\n", err)
		return
	}
	defer req.Body.Close()

	// Just incase someone entered some random name
	if req.StatusCode == 404 {
		color.Red("[!] Exe not found on lolbas")
		return
	}

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return
	}
	config := make(map[interface{}]interface{})
	if err = yaml.Unmarshal(body, &config); err != nil {
		fmt.Println(err)
	}
	fmt.Println(config)
}

func main() {
	var bin string
	flag.StringVar(&bin, "bin", "", "")
	flag.StringVar(&bin, "b", "", "")

	var exe string
	flag.StringVar(&exe, "exe", "", "")
	flag.StringVar(&exe, "e", "", "")

	flag.Parse()

	if bin != "" {
		gtfobins(bin)
	} else if exe != "" {
		// TODO: Implement support for lolbas
		lolbas(exe)
	} else {
		fmt.Println("No option selected")
		os.Exit(2)
	}
}
