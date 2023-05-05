package repository

import (
	"context"

	"notifier/pkg/entity"
	"notifier/pkg/repository/cache"
	"notifier/pkg/repository/client"
	"notifier/pkg/repository/storage"
)

type WarReportRepository interface {
	GetLatestReport(ctx context.Context, m entity.Map) (*entity.WarReport, error)
}

func newWarReportRepository(
	client client.Client[*entity.WarReport],
	cache cache.Cache[*entity.WarReport],
	storage storage.Storage[*entity.WarReport],
) WarReportRepository {
	return warReportRepository{
		client,
		cache,
		storage,
	}
}

type warReportRepository struct {
	client  client.Client[*entity.WarReport]
	cache   cache.Cache[*entity.WarReport]
	storage storage.Storage[*entity.WarReport]
}

func (w warReportRepository) GetLatestReport(ctx context.Context, m entity.Map) (*entity.WarReport, error) {
	report := &entity.WarReport{
		Map: m,
	}

	report, err := w.client.Refresh(report)
	if err != nil {
		return nil, err
	}

	storedReport, err := w.storage.Save(ctx, report)
	if err != nil {
		return nil, err
	}

	return storedReport, nil
}
