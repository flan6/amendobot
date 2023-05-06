package repository

import (
	"context"
	"encoding/csv"
	"os"
	"path"
	"runtime"

	"scraper/entity"
)

type JobPostingRepository interface {
	Save(_ context.Context, posting entity.JobPosting) error
}

func NewJobPostingRepository() JobPostingRepository {
	_, currentFile, _, ok := runtime.Caller(0)
	if !ok {
		panic("could not find current file directory in job_posting.go")
	}

	return jobPostingRepository{path.Dir(currentFile)}
}

const filePath = "/storage/job_posting.csv"

type jobPostingRepository struct {
	curentPath string
}

func (j jobPostingRepository) Save(_ context.Context, posting entity.JobPosting) error {
	csvFile, err := os.OpenFile(j.curentPath+filePath, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}
	defer csvFile.Close()

	writer := csv.NewWriter(csvFile)
	writer.Comma = ';'
	defer writer.Flush()

	return writer.Write(posting.ToCSV())
}

func (j jobPostingRepository) ReadAll(_ context.Context) ([]entity.JobPosting, error) {
	csvFile, err := os.Open(j.curentPath + filePath)
	if err != nil {
		return nil, err
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)

	data, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	content := make([]entity.JobPosting, 0, len(data))
	for _, row := range data[1:] {
		posting := entity.JobPosting{
			PostURL:     row[0],
			Company:     row[1],
			ContentText: row[2],
			ApplyURL:    row[4],
		}

		if row[3] == "1" {
			posting.EasyApply = true
		}

		content = append(content, posting)
	}

	return content, nil
}
