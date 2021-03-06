package util

import (
	"io/ioutil"
	"log"
	"net/http"
)

// RequestURL - Download page form given url
func RequestURL(url string, cookie string, header string) string {

	// Request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	// If has cookies to parse
	if cookie != "" {

		// Parse cookies from flag
		cookies := ParseCookie(cookie)

		// add all cookies to request
		for c := range cookies {

			// Cookie name, Cookie value
			req.AddCookie(&http.Cookie{Name: cookies[c][0], Value: cookies[c][1]})
		}
	}

	// if has headers to parse
	if header != "" {

		// parse headers from flag
		headers := ParseHeader(header)

		// Add headers to request
		for c := range headers {

			// Header Name, Header value
			req.Header.Set(headers[c][0], headers[c][1])
		}
	}

	// make client
	client := &http.Client{}

	// Execute request
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	// Close request body
	defer res.Body.Close()

	// Parse request body
	body, _ := ioutil.ReadAll(res.Body)

	// return page as string
	return string(body)
}
