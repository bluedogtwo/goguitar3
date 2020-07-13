package main

import (
	"fmt"
	"time"

	"github.com/go-rod/rod"
)

const (
	row  string = ".flex.items-center.bb.pb2.b--light-gray"
	row2 string = ".flex.items-center.bb.pv2.b--light-gray"

	band string = ".link.outline-0.source-sans-b.black.w-40.pr3.h2.flex.items-center"
	song string = ".black.db.w-40.pr2.source-sans.ttc.h2.flex.items-center"
)

func main() {
	browser := rod.New().Connect()

	defer browser.Close()

	page := browser.Timeout(time.Minute).Page("https://www.fender.com/play/guitar/artists/songs")

	page.Window(0, 0, 1200, 600)

	fmt.Println("done1")

	vv := page.Element(row).Text() // doesnt work if removed?
	fmt.Printf("%s \n", vv)        //

	for _, e := range page.Elements(row + "," + row2) {
		bandPrint := e.Element(band).Text()

		for _, es := range e.Elements(song) {
			fmt.Printf("%s - %s\n", bandPrint, es.Text())

		}
	}
}
