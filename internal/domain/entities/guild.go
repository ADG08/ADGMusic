package entities

import "time"

type Guild struct {
	ID        int64
	Name      string
	CreatedAt time.Time
}

func NewGuild(name string) *Guild {
	return &Guild{
		Name:      name,
		CreatedAt: time.Now(),
	}
}
