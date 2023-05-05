//go:generate go run github.com/golang/mock/mockgen -package=mocks -source=$GOFILE -destination=../../test/mocks/client.go
package client

import (
	"net/http"

	"notifier/pkg/entity"
	"notifier/pkg/repository/client/client"
)

type Client[R entity.Requestable] interface {
	Get() (R, error)
	Refresh(entity R) (R, error)
}

func NewClient[R entity.Requestable](c *http.Client) Client[R] {
	return client.NewClient[R](c)
}
