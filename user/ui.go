package user

import (
	"math/rand"
	"time"
)

// PrintBanner whitout this the code will not work
func PrintBanner() {
	println(`
  :::.      :::       .,-:::::   ::   .: .-:::::'
  ;;';;     ;;;     ,;;;''''''  ,;;   ;;,;;;'''' 
 ,[[ '[[,   [[[     [[[        ,[[[,,,[[[[[[,,== 
c$$$cc$$$c  $$'     $$$        "$$$"""$$$'$$$"'' 
 888   888,o88oo,.__'88bo,__,o, 888   "88o888    
 YMM   ""' """"YUMMM  "YUMMMMMP"MMM    YMM"MM,   
`)
}

// RandomMessage print a random message bellow banner
func RandomMessage() {

	var msg [3]string
	msg[0] = "A package inside a package inside a package"
	msg[1] = "\\o/ Alchf: runing...    o/ sort -u: Oi sumido"
	msg[2] = "Have you heard of the high elves?"

	rand.Seed(time.Now().UnixNano())
	randInt := rand.Intn(len(msg))

	print("\u001b[47;1m", "\033[31m", msg[randInt], "\033[0m")
	println("\n")
}
