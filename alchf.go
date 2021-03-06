package main

import (
	mod "alchf/modules"
	"alchf/user"
	"flag"
)

func main() {

	// Warning
	user.PrintBanner()

	// function flag
	var module string
	flag.StringVar(&module, "m", "", "Select module")

	// url flag
	var url string
	flag.StringVar(&url, "u", "", "Target url")

	// file flag
	var file string
	flag.StringVar(&file, "f", "", "Target file")

	// Cookies
	var cookie string
	flag.StringVar(&cookie, "C", "", "Request Cookies (Optional)")

	// Header
	var header string
	flag.StringVar(&header, "H", "", "Request Header (Optional)")

	// Parse flag values
	flag.Parse()

	switch module {
	case "recon":
		mod.Recon()
		break

	case "dis":
		mod.Dissect(file, url, cookie, header)
		break

	default:
		println("Chose a module")
	}

}
