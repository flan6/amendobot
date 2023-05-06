package service

import (
	"context"
	"log"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/proto"

	"scraper/lib/linkedin"
)

type ScrapService interface {
	FindJobsPageURL(ctx context.Context, keyWords []string, location string) chan string
}

type scrapService struct {
	broswer *rod.Browser
	logger  *log.Logger
}

func NewScrapService(broswer *rod.Browser) ScrapService {
	logger := log.Default()
	logger.SetFlags(log.Llongfile | log.Ltime)

	return scrapService{broswer, logger}
}

func (s scrapService) FindJobsPageURL(ctx context.Context, keyWords []string, location string) chan string {
	result := make(chan string)

	go func(context.Context) {
		initialPosition, initialPageNumber := 1, 0
		searchUrl := linkedin.GetSearchURL(keyWords, location, initialPosition, initialPageNumber)

		page := s.broswer.MustPage(searchUrl)
		page.MustWaitLoad()

		seeAll := false
		for !seeAll {
			_, err := page.Eval("window.scrollTo(0, document.body.scrollHeight);")
			if err != nil {
				seeMoreBtn := page.MustElementX(linkedin.JobSeeMoreJobsButton)
				if seeMoreBtn != nil {
					err := page.Mouse.MoveTo(*seeMoreBtn.MustShape().OnePointInside()) // TODO : panics
					if err != nil {
						s.logger.Println(err)
						break
					}
					page.Mouse.MustClick(proto.InputMouseButtonLeft)
				}
			}
			page.MustWaitLoad()

			viewedAllElement := page.MustElementX(linkedin.JobViewedAll)
			if viewedAllElement != nil {
				seeAll = true
			}
		}

		jobList, err := page.ElementX(linkedin.JobListXPath)
		if err != nil {
			s.logger.Fatal(err)
		}

		jobsEntries, err := jobList.ElementsX(linkedin.JobEntriesXPath)
		if err != nil {
			s.logger.Fatal(err)
		}

		s.logger.Printf("found %d jobs", len(jobsEntries))
		for _, jobEntry := range jobsEntries {
			result <- *jobEntry.MustAttribute("href")
		}
	}(ctx)

	return result
}

func (s scrapService) CrawlJobPage(_ context.Context, pageURL chan string) error {
	return nil
}
