package main

import (
	"context"
	"fmt"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"

	"scraper/service"
)

func main() {
	l := launcher.New().
		Headless(false).
		Devtools(true)
	defer l.Cleanup()

	url := l.MustLaunch()

	browser := rod.New().
		ControlURL(url).
		Trace(true).
		SlowMotion(2 * time.Second).
		MustConnect()
	defer browser.MustClose()

	scrapService := service.NewScrapService(browser)

	result := scrapService.FindJobsPageURL(
		context.Background(),
		[]string{"backend", "developer", "go"},
		"brazil",
	)

	for _, link := range result {
		fmt.Println(link)
	}

	fmt.Println("done")

	l.Kill()
}
