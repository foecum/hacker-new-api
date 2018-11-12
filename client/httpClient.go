package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"time"
)

const hackerRankAPIURL = "https://hacker-news.firebaseio.com/v0/"

// Requester ...
type Requester struct {
	c http.Client
}

var timeout = time.Duration(2 * time.Second)

func dialTimeout(network, addr string) (net.Conn, error) {
	return net.DialTimeout(network, addr, timeout)
}

// NewHTTPClient creates a new http client
func NewHTTPClient() Requester {
	t := http.Transport{
		Dial: dialTimeout,
	}

	c := http.Client{
		Transport: &t,
	}

	return Requester{c: c}
}

// MakeRequest ...
func (r Requester) MakeRequest(method, path string, body io.Reader) (*json.Decoder, error) {
	url := fmt.Sprintf("%s%s", hackerRankAPIURL, path)

	req, err := http.NewRequest(method, url, body)

	if err != nil {
		return nil, err
	}

	resp, err := r.c.Do(req)
	if err != nil {
		return nil, err
	}
	return json.NewDecoder(resp.Body), nil
}