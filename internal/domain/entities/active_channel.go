package entities

import "time"

type ActiveChannel struct {
	ID        int64
	GuildID   string
	ChannelID string
	CreatedAt time.Time
}

func NewActiveChannel(guildID, channelID string) *ActiveChannel {
	return &ActiveChannel{
		GuildID:   guildID,
		ChannelID: channelID,
		CreatedAt: time.Now(),
	}
}
