package service

import (
	"context"
	"fmt"

	"notifier/pkg/repository"
)

type Service interface {
	PrintRepositoryMethods(ctx context.Context)
}

type service struct {
	warRepository       repository.WarRepository
	mapsRepository      repository.MapsRepository
	warReportRepository repository.WarReportRepository
}

func NewService(
	warRepository repository.WarRepository,
	mapRepository repository.MapsRepository,
	warReportRepository repository.WarReportRepository,
) Service {
	return service{
		warRepository,
		mapRepository,
		warReportRepository,
	}
}

func (s service) PrintRepositoryMethods(ctx context.Context) {
	war, err := s.warRepository.GetLatestWarState(ctx)
	if err != nil {
		fmt.Printf("could not get latest war, error: %s", err.Error())
	}

	fmt.Println(war)

	maps, err := s.mapsRepository.GetMaps(ctx)
	if err != nil {
		fmt.Printf("get maps error %s", err.Error())
	}

	for _, m := range *maps {
		report, err := s.warReportRepository.GetLatestReport(ctx, m)
		if err != nil {
			fmt.Printf("got error in report from map: %s error: %s", m, err)
			return
		}

		fmt.Println(report)
	}
}
