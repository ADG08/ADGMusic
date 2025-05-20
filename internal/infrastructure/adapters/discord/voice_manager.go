package discord

import (
	"context"

	"github.com/bwmarrin/discordgo"
)

type VoiceManager struct {
	session *discordgo.Session
}

func NewVoiceManager(session *discordgo.Session) *VoiceManager {
	return &VoiceManager{session: session}
}

func (v *VoiceManager) Join(ctx context.Context, guildID, channelID string) error {
	_, err := v.session.ChannelVoiceJoin(guildID, channelID, false, false)
	return err
}

func (v *VoiceManager) Leave(ctx context.Context, guildID string) error {
	for _, voiceConnection := range v.session.VoiceConnections {
		if voiceConnection.GuildID == guildID {
			voiceConnection.Disconnect()
		}
	}
	return nil
}
