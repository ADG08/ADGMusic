package domain

import "time"

// Music représente une entité de musique dans le domaine
type Music struct {
	ID        int64
	GuildID   string
	ChannelID string
	URL       string
	Title     string
	AddedBy   string
	CreatedAt time.Time
}

// NewMusic crée une nouvelle instance de Music
func NewMusic(guildID, channelID, url, title, addedBy string) *Music {
	return &Music{
		GuildID:   guildID,
		ChannelID: channelID,
		URL:       url,
		Title:     title,
		AddedBy:   addedBy,
		CreatedAt: time.Now(),
	}
}
