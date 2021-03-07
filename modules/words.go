package modules

import (
	"alchf/util"
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strings"
)

// ExtractWords extract all words from a page
func ExtractWords(file string, target string, cookie string, header string) {

	if target != "" {
		pageData := util.RequestURL(target, cookie, header)

		var stringData = string(pageData)

		// Remove all non alpha numeric from string
		var ripNon = regexp.MustCompile(`(?m)[^a-zA-Z0-9]+`)

		// Clean string
		var ripNonAlpha = ripNon.ReplaceAllString(stringData, " ")

		// split all words
		var re = regexp.MustCompile(`(?m).+?[\W]`)

		// Want some tee?
		for _, word := range re.FindAllString(ripNonAlpha, -1) {
			// Print word by word
			fmt.Println(strings.Replace(word, " ", "", -1))
		}
	}

	if file != "" {

		// Read file data
		file, err := ioutil.ReadFile(file)
		if err != nil {
			log.Fatal(err)
		}

		// File to string
		var stringData = string(file)

		// Remove all non alpha numeric from string
		var ripNon = regexp.MustCompile(`(?m)[^a-zA-Z0-9]+`)

		// Clean string
		var ripNonAlpha = ripNon.ReplaceAllString(stringData, " ")

		// split all words
		var re = regexp.MustCompile(`(?m).+?[\s]`)

		// More tee?
		for _, word := range re.FindAllString(ripNonAlpha, -1) {
			// Print words
			fmt.Println(strings.Replace(word, " ", "", -1))
		}
	}
}
