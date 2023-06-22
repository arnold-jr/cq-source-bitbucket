package lib

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"
)

type Client struct {
	workspace string
	username  string
	password  string
}

func New() *Client {
	return &Client{
		workspace: os.Getenv("BITBUCKET_WORKSPACE"),
		username:  os.Getenv("BITBUCKET_USERNAME"),
		password:  os.Getenv("BITBUCKET_PASSWORD"),
	}
}

func (c *Client) request(url string, pageLength, curPage int) (*http.Response, error) {
	client := http.Client{Timeout: 5 * time.Second}

	req, err := http.NewRequest(http.MethodGet, url, http.NoBody)
	if err != nil {
		return nil, fmt.Errorf("Error creating users requests url: %s", err.Error())
	}

	q := req.URL.Query()
	q.Set("pagelen", strconv.Itoa(pageLength))
	q.Set("page", strconv.Itoa(curPage))
	req.URL.RawQuery = q.Encode()

	req.SetBasicAuth(c.username, c.password)

	return client.Do(req)
}
