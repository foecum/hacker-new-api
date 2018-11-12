package user

import "time"

// User ...
type User struct {
	ID        int64     `json:"id,omitempty"`
	Delay     int64     `json:"delay"`
	Created   time.Time `json:"created"`
	Karma     int64     `json:"karma"`
	About     string    `json:"about"`
	Submitted []int64   `json:"submited"`
}
