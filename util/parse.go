package util

import (
	"strings"
)

// ParseCookie array of array
func ParseCookie(cookie string) [][2]string {

	// Cookies array
	var cookies [][2]string

	splitSemiCol := strings.Split(cookie, ";")

	for i := range splitSemiCol {

		splitEqual := strings.Split(splitSemiCol[i], "=")

		if len(splitEqual) > 1 {

			data := [2]string{splitEqual[0], splitEqual[1]}

			cookies = append(cookies, data)
		}
	}

	// return array of cookies
	return cookies
}

// ParseHeader array of array
func ParseHeader(header string) [][2]string {

	// headers array
	var headers [][2]string

	splitSemiCol := strings.Split(header, ";")

	for i := range splitSemiCol {

		splitEqual := strings.Split(splitSemiCol[i], ":")

		if len(splitEqual) > 1 {

			data := [2]string{splitEqual[0], splitEqual[1]}

			headers = append(headers, data)
		}
	}

	// return array of headers
	return headers
}
