package audio

import (
	"context"
	"fmt"
	"net/http"
)

const Origin = "https://api.bookmate.yandex.net"

func NewAudioBookRequest(ctx context.Context, uuid string) *http.Request {
	url := fmt.Sprint(Origin, "/api/v5/audiobooks/", uuid)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		panic(err.Error())
	}
	return req
}
