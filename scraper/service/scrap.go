package service

import (
	"context"
	"log"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/proto"

	"scraper/lib/linkedin"
)

type ScrapService interface {
	FindJobsPageURL(ctx context.Context, keyWords []string, location string) []string
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

func (s scrapService) FindJobsPageURL(ctx context.Context, keyWords []string, location string) []string {
	searchUrl := linkedin.GetSearchURL(keyWords, location, 1, 0)

	page := s.broswer.MustPage(searchUrl)
	page.MustWaitLoad()

	for {
		cctx, cancel := context.WithCancel(ctx)
		defer cancel()

		go func(cctx context.Context) {
			for {
				page.Eval("window.scrollTo(0, document.body.scrollHeight);")
			}
		}(cctx)

		go func(cctx context.Context) {
			for {
				seeMoreBtn := page.MustElementX(linkedin.JobSeeMoreJobsButton)
				if seeMoreBtn != nil {
					seeMoreBtn.MustWaitVisible()

					err := page.Mouse.MoveTo(*seeMoreBtn.MustShape().OnePointInside())
					if err != nil {
						s.logger.Println(err)
						return
					}

					page.Mouse.MustClick(proto.InputMouseButtonLeft)
				}
			}
		}(cctx)

		viewedAllElement := page.MustElementX(linkedin.JobViewedAll)
		if viewedAllElement != nil {
			viewedAllElement.MustWaitVisible()

			cancel()
			break
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
	result := make([]string, 0, len(jobsEntries))
	for _, jobEntry := range jobsEntries {
		result = append(result, *jobEntry.MustAttribute("href"))
	}

	return result
}

func (s scrapService) CrawlJobPage(_ context.Context, pageURL chan string) error {
	return nil
}
