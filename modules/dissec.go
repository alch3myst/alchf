package modules

import (
	"alchf/util"
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
)

// Dissect module
func Dissect(file string, target string, cookie string, header string) {

	if target != "" {

		// Load url page content
		body := util.RequestURL(target, cookie, header)

		// Extract api like strings from page
		apiLike(string(body))

		// Extract links from page
		httpLink(string(body))
	}

	if file != "" {

		// Read file data
		file, err := ioutil.ReadFile(file)
		if err != nil {
			log.Fatal(err)
		}

		// Extract api like strings from file
		apiLike(string(file))

		// Extract links from file
		httpLink(string(file))
	}

	return
}

func apiLike(data string) {
	// Reg api like strings
	var apiLikeRe = regexp.MustCompile(`(?m)/[a-zA-Z0-9?\.=\-&\/_\+]*`)

	// Print match
	for _, apiFound := range apiLikeRe.FindAllString(string(data), -1) {
		fmt.Println("[Api Like]", apiFound)
	}
}

func httpLink(data string) {
	// Reg http links
	var httpRe = regexp.MustCompile(`(?m)(http|https)://[a-zA-Z0-9?\.=\-&\/_\+]*`)

	// Print match
	for _, httpFound := range httpRe.FindAllString(string(data), -1) {
		fmt.Println("[Link]", httpFound)
	}
}
