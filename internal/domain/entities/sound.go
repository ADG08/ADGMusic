package entities

import "time"

type Sound struct {
	ID        int64
	Name      string
	URL       string
	CreatedAt time.Time
}

func NewSound(name, url string) *Sound {
	return &Sound{
		Name:      name,
		URL:       url,
		CreatedAt: time.Now(),
	}
}
