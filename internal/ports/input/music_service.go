package input

import "context"

type MusicService interface {
	PlayMusic(ctx context.Context, guildID, channelID, url string) error
	StopMusic(ctx context.Context, guildID string) error
	SkipMusic(ctx context.Context, guildID string) error
}
