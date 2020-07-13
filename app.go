package main

import (
	"fmt"
	"time"

	"github.com/go-rod/rod"
)

const (
	row  string = ".flex.items-center.bb.pb2.b--light-gray"
	band string = ".link.outline-0.source-sans-b.black.w-40.pr3.h2.flex.items-center"

	// salary float64 = 50000.0
)

func main() {
	browser := rod.New().Connect()

	defer browser.Close()

	page := browser.Timeout(time.Minute).Page("https://www.fender.com/play/guitar/artists/songs")

	page.Window(0, 0, 1200, 600)

	fmt.Println("done1")

	// Wait until css selector get the element then get the text content of it.
	// band
	// text := page.Element(".link.outline-0.source-sans-b.black.w-40.pr3.h2.flex.items-center").Text()
	// song

	text := page.Element(row).Text()
	fmt.Printf("Hello %s\n", text)

	for i, e := range page.Elements(row) {
		fmt.Printf("%d \n", i)
		temp := e.Element(band).Text()
		fmt.Printf("%d %s\n", i, temp)
	}
}
