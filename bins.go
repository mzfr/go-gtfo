package bins

import (
	"fmt"
)

var URL = "https://gtfobins.github.io/"
var RAW_URL = "https://raw.githubusercontent.com/GTFOBins/GTFOBins.github.io/master/_gtfobins/%s.md"

func gtfobins(binary) {
	val, err := url.ParseRequestURI(URL)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(val)
	

}
