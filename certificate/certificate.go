package certificate

import "time"

type Certificate struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Authority   string    `json:"authority"`
	Created     time.Time `json:"created_timestamp"`
}
