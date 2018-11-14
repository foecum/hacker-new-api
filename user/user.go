package user

import (
	"fmt"
	"time"

	"github.com/foecum/hackernews/http"
)

// User ...
type User struct {
	ID        string        `json:"id,omitempty"`
	Delay     int64         `json:"delay"`
	Created   time.Duration `json:"created"`
	Karma     int64         `json:"karma"`
	About     string        `json:"about"`
	Submitted []int64       `json:"submited"`
}

func getHackerNewUser(username string) (User, error) {
	path := fmt.Sprintf("user/%s.json?print=pretty", username)
	c, err := client.NewHTTPClient()

	user := User{}

	if err != nil {
		return user, err
	}

	decoder, err := c.MakeRequest("GET", path, nil)

	if err != nil {
		return user, err
	}

	err = decoder.Decode(&user)

	if err != nil {
		return user, err
	}
	return user, nil
}
