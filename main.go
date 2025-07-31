package main

import (
	"fmt"
	"log"

	"github.com/playwright-community/go-playwright"
)

func main() {
	pw, err := playwright.Run()
	if err != nil {
		log.Fatalf("could not start playwright: %v", err)
	}

	browser, err := pw.Chromium.Launch()
	if err != nil {
		log.Fatalf("could not launch browser: %v", err)
	}

	page, err := browser.NewPage()
	if err != nil {
		log.Fatalf("could not create page: %v", err)
	}

	if _, err = page.Goto("https://example.com"); err != nil {
		log.Fatalf("could not goto: %v", err)
	}

	title, err := page.Title()
	if err != nil {
		log.Fatalf("could not get title: %v", err)
	}
	fmt.Println(title)

	if err = browser.Close(); err != nil {
		log.Fatalf("could not close browser: %v", err)
	}

	if err = pw.Stop(); err != nil {
		log.Fatalf("could not stop Playwright: %v", err)
	}
}
