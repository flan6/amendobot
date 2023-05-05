package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"notifier/pkg/config"
	"notifier/pkg/entity"
)

type Client[R entity.Requestable] struct {
	httpClient *http.Client
}

func NewClient[R entity.Requestable](c *http.Client) Client[R] {
	return Client[R]{c}
}

func (c Client[R]) Get() (R, error) {
	var entity R

	req, err := http.NewRequest(
		http.MethodGet,
		fmt.Sprint(
			config.WarApiURL(),
			entity.ApiEndpoint(),
		),
		nil,
	)
	if err != nil {
		return entity, err
	}

	etag := entity.Etag()
	if etag != "" {
		req.Header.Add("If-None-Match", etag)
	}

	response, err := c.httpClient.Do(req)
	if err != nil {
		return entity, err
	}
	defer response.Body.Close()

	if response.StatusCode == http.StatusNotModified {
		return entity, nil
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return entity, err
	}

	err = json.Unmarshal(body, &entity)
	if err != nil {
		return entity, err
	}

	entity.GenId()
	entity.SetEtag(response.Header.Get("etag"))
	return entity, nil
}

func (c Client[R]) Refresh(entity R) (R, error) {
	req, err := http.NewRequest(
		http.MethodGet,
		fmt.Sprint(
			config.WarApiURL(),
			entity.ApiEndpoint(),
		),
		nil,
	)
	if err != nil {
		return entity, err
	}

	if etag := entity.Etag(); etag != "" {
		req.Header.Set("If-None-Match", fmt.Sprintf(`"%s"`, etag))
	}

	response, err := c.httpClient.Do(req)
	if err != nil {
		return entity, err
	}
	defer response.Body.Close()

	if response.StatusCode == http.StatusNotModified {
		return entity, nil
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return entity, err
	}

	err = json.Unmarshal(body, &entity)
	if err != nil {
		return entity, err
	}

	entity.GenId()
	entity.SetEtag(response.Header.Get("etag"))
	return entity, nil
}
