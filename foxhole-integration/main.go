package main

import (
	"context"
	"net/http"
	"time"

	tcache "github.com/jellydator/ttlcache/v3"

	"notifier/pkg/repository"
	"notifier/pkg/service"
)

func main() {
	ctx := context.Background()
	c := tcache.New(
		tcache.WithTTL[string, any](time.Hour),
	)
	go c.Start()

	httpClient := &http.Client{
		Timeout: time.Hour,
	}

	serv := initService(
		c,
		httpClient,
	)

	serv.PrintRepositoryMethods(ctx)
}

func initService(
	ttlCache *tcache.Cache[string, any],
	httpClient *http.Client,
) service.Service {
	all := repository.GetAll(ttlCache, httpClient)

	return service.NewService(
		all.WarRepository,
		all.MapsRepository,
		all.WarReportRepository,
	)
}
