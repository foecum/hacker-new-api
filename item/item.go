package item

import (
	"fmt"
	"time"

	"github.com/foecum/hackernews/conn"
)

// Item fro Hackernews
type Item struct {
	Dead    bool          `json:"dead"`
	Deleted bool          `json:"deleted"`
	ID      int64         `json:"id,ommitempty"`
	Parent  int64         `json:"parent"`
	Score   int64         `json:"score"`
	Created time.Duration `json:"time"`
	Kids    []int64       `json:"kids"`
	Type    string        `json:"type"`
	Author  string        `json:"by"`
	URL     string        `json:"url"`
	Title   string        `json:"title"`
	Text    string        `json:"text"`
}

func getHackerNewItem(itemID int) (Item, error) {
	path := fmt.Sprintf("item/%d.json?print=pretty", itemID)
	c, err := conn.NewClient("")

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
