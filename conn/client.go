package conn

import (
	"time"

	"github.com/foecum/going/client"
)

// NewClient creates a new client
func NewClient(endpoint string, timeout int) (client.Requester, error) {
	cfg := client.Config{
		Endpoint: endpoint,
		Timeout:  time.Duration(timeout) * time.Second,
	}
	c, err := client.NewHTTPClient(cfg)

	if err != nil {
		return client.Requester{}, err
	}

	return c, nil
}
