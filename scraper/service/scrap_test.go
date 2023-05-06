package service_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"

	"scraper/service"
)

func TestFindJobsAndSave(t *testing.T) {
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

	launcher.Open(browser.ServeMonitor(""))

	scrapService := service.NewScrapService(browser)

	result := scrapService.FindJobsPageURL(
		context.Background(),
		[]string{"backend", "developer"},
		"brazil",
	)

	for link := range result {
		fmt.Println(link)
	}
}
