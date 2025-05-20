package output

import (
	"context"
)

type VoiceManager interface {
	Join(ctx context.Context, guildID, channelID string) error
	Leave(ctx context.Context, guildID string) error
}
