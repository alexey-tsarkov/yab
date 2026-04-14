package yab

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/alexey-tsarkov/yab/book"
)

type Client struct {
	httpClient *http.Client
	token      string
}

func NewClient(httpClient *http.Client, token string) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	return &Client{
		httpClient: httpClient,
		token:      token,
	}
}

func (c *Client) Book(ctx context.Context, uuid string) (*book.Book, error) {
	req := c.authenticate(book.NewBookRequest(ctx, uuid))
	var v struct {
		Book *book.Book `json:"book"`
	}
	err := c.perform(req, &v)

	return v.Book, err
}

func (c *Client) DownloadBook(ctx context.Context, uuid string, w io.Writer) error {
	req := c.authenticate(book.NewBookContentRequest(ctx, uuid))
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	io.Copy(io.Discard, resp.Body)

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("http status: %s", resp.Status)
	}

	_, err = io.Copy(w, resp.Body)
	return err
}

func (c *Client) authenticate(req *http.Request) *http.Request {
	req.Header.Add("Auth-Token", c.token)
	req.Header.Add("Authorization", fmt.Sprint("OAuth ", c.token))
	return req
}

func (c *Client) perform(req *http.Request, v any) error {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	defer io.Copy(io.Discard, resp.Body)

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("http status: %s", resp.Status)
	}

	return json.NewDecoder(resp.Body).Decode(&v)
}
