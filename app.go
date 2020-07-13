package main

import (
	"fmt"
	"time"

	"github.com/go-rod/rod"
)

func main() {
	// Launch a new browser with default options, and connect to it.
	browser := rod.New().Connect()

	// Even you forget to close, rod will close it after main process ends.
	defer browser.Close()

	// Timeout will be passed to all chained function calls.
	// The code will panic out if any chained call is used after the timeout.
	page := browser.Timeout(time.Minute).Page("https://www.fender.com/play/guitar/artists/songs")

	// Resize the window make sure window size is always consistent.
	page.Window(0, 0, 1200, 600)

	fmt.Println("done1")

	// Wait until css selector get the element then get the text content of it.
	// band
	// text := page.Element(".link.outline-0.source-sans-b.black.w-40.pr3.h2.flex.items-center").Text()
	// song
	// text := page.Element(".black.db.w-40.pr2.source-sans.ttc.h2.flex.items-center").Text()
	// runtom
	text := page.Element(".flex.items-center.bb.pb2.b--light-gray").HTML()

	fmt.Println(text)

	// Output: Git is the most widely used version control system.
}
