package models

import "time"

type Result struct {
	ID        int64     `json:"id"`
	Query     string    `json:"query"`
	Source    string    `json:"source"`
	Summary   string    `json:"summary"`
	URL       string    `json:"url"`
	CreatedAt time.Time `json:"created_at"`
}
