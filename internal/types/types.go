package types

import "time"

type URLInfo struct {
	URL           string
	RedirectCount int
	CreatedAt     time.Time
}

type ShortenResp struct {
	ShortenedURL string `json:"shortened_url"`
}
