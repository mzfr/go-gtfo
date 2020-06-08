package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/common-nighthawk/go-figure"
	"github.com/fatih/color"
	"gopkg.in/yaml.v2"
)

//TODO: Lot of code is repeating. Need to figure it out.
var rawBinURL = "https://raw.githubusercontent.com/GTFOBins/GTFOBins.github.io/master/_gtfobins/%s.md"
var rawExeURL = "https://raw.githubusercontent.com/LOLBAS-Project/LOLBAS-Project.github.io/master/_lolbas/%s.md"

func init() {
	flag.Usage = func() {
		h := []string{
			"Search gtfobin from terminal",
			"",
			"Options:",
			"  -b, --bin <binary>       Search Linux binaries on gtfobins",
			"  -e, --exe <EXE>          Search Linux binaries on gtfobins",
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

func lolbas(exe string) {
	config := make(map[interface{}]interface{})
	exeMap := make(map[string]string)

	doc, err := goquery.NewDocument("https://lolbas-project.github.io/")
	if err != nil {
		log.Fatal(err)
	}

	doc.Find(".bin-name").Each(func(index int, item *goquery.Selection) {
		href, _ := item.Attr("href")
		exeMap[item.Text()] = href[8 : len(href)-1]
	})

	// TODO: ignore case
	if val, ok := exeMap[exe]; ok {

		exeURL := fmt.Sprintf(rawExeURL, val)

		req, err := http.Get(exeURL)
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to create request: %s\n", err)
			return
		}

		defer req.Body.Close()

		// Just incase someone entered some random name
		if req.StatusCode == 404 {
			color.Red("[!] Exe not found on Lolbas")
			return
		}

		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			return
		}

		if err = yaml.Unmarshal(body, &config); err != nil {
			fmt.Println(err)
		}
		// fmt.Println(reflect.TypeOf(config["Commands"]))

		yellow := color.New(color.FgYellow)
		boldYellow := yellow.Add(color.Bold)
		green := color.New(color.FgGreen).SprintFunc()
		magenta := color.New(color.FgHiMagenta).SprintFunc()

		for _, key := range config["Commands"].([]interface{}) {
			details := key.(map[interface{}]interface{})
			boldYellow.Println("\n# ", details["Description"])
			fmt.Printf("CMD:\t\t%v \n", green(details["Command"]))
			fmt.Printf("Category:\t%v\n", magenta(details["Category"]))
			fmt.Printf("Privileges:\t%v\n", magenta(details["Privileges"]))
			fmt.Println()
		}
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
	myFigure := figure.NewColorFigure("# gtfo", "big", "green", true)
	myFigure.Print()

	if bin != "" {
		gtfobins(bin)
	} else if exe != "" {
		lolbas(exe)
	}
}
