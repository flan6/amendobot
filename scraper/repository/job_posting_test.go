package repository_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"scraper/entity"
	"scraper/repository"
)

func TestSave(t *testing.T) {
	repo := repository.NewJobPostingRepository()

	err := repo.Save(context.Background(), entity.JobPosting{
		PostURL: "someUrl",
	})
	require.NoError(t, err)
}
