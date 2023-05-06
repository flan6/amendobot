package linkedin

import (
	"fmt"
	"strings"
)

const (
	baseLinkedinSearchURL = "https://www.linkedin.com/jobs/search?" +
		"keywords=%s" +
		"&location=%s" +
		"&geoId=106057199" +
		"&trk=public_jobs_jobs-search-bar_search-submit" +
		"&position=%d" +
		"&pageNum=%d`"
)

const (
	JobListXPath         = `//ul[@class="jobs-search__results-list"]`
	JobEntriesXPath      = `li/div/a`
	JobViewedAll         = `//p[contains(text(), "You've viewed all jobs for this search")]`
	JobSeeMoreJobsButton = `//button[contains(text(), "See more jobs")]`
)

func GetSearchURL(keyWords []string, location string, position, pageNumber int) string {
	paramKeywords := strings.Join(keyWords, "%2B")

	return fmt.Sprintf(
		baseLinkedinSearchURL,
		paramKeywords,
		location,
		position,
		pageNumber,
	)
}
