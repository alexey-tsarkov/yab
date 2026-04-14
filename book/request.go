package book

import (
	"context"
	"fmt"
	"net/http"
)

const Origin = "https://api.bookmate.yandex.net"

func NewBookRequest(ctx context.Context, uuid UUID) *http.Request {
	url := fmt.Sprint(Origin, "/api/v5/books/", uuid)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		panic(err.Error())
	}
	return req
}

func NewBookContentRequest(ctx context.Context, uuid UUID) *http.Request {
	url := fmt.Sprint(Origin, "/api/v5/books/", uuid, "/content/v4")
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		panic(err.Error())
	}
	return req
}
