package main

import (
	mod "alchf/modules"
	"alchf/user"
	"flag"
)

func main() {

	// Warning
	user.PrintBanner()
	user.RandomMessage()

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

	case "word":
		mod.ExtractWords(file, url, cookie, header)

	default:
		println("Modules: recon, dis, word")
		println("Dis: extract all urls, api like and links from file")
		println("Word: Split all worlds in a file and output them.")
		println("Recon: TODO")
		println("usage: alchf -m MODULE -f FILE")
	}
}
