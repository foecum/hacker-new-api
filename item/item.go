package item

import (
	"fmt"
	"time"

	"github.com/foecum/hackernews/http"
)

// Item fro Hackernews
type Item struct {
	ID      int64         `json:"id,ommitempty"`
	Deleted bool          `json:"deleted"`
	Type    string        `json:"type"`
	Author  string        `json:"by"`
	Created time.Duration `json:"time"`
	Dead    bool          `json:"dead"`
	Parent  int64         `json:"parent"`
	Kids    []int64       `json:"kids"`
	URL     string        `json:"url"`
	Score   int64         `json:"score"`
	Title   string        `json:"title"`
	Text    string        `json:"text"`
}

func getHackerNewItem(itemID int) (Item, error) {
	path := fmt.Sprintf("item/%d.json?print=pretty", itemID)
	c, err := client.NewHTTPClient()

	item := Item{}

	if err != nil {
		return item, err
	}

	decoder, err := c.MakeRequest("GET", path, nil)

	if err != nil {
		return item, err
	}

	err = decoder.Decode(&item)

	if err != nil {
		return item, err
	}
	return item, nil
}
