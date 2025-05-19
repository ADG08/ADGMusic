package discord

import (
	"context"
	"strings"

	"github.com/ADG08/ADGMusic/internal/ports/output"
	"github.com/bwmarrin/discordgo"
)

type CommandHandler struct {
	voiceManager output.VoiceManager
}

func NewCommandHandler(voiceManager output.VoiceManager) *CommandHandler {
	return &CommandHandler{voiceManager: voiceManager}
}

func (h *CommandHandler) HandleCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}
	if !strings.HasPrefix(m.Content, "!") {
		return
	}
	args := strings.Fields(m.Content[1:])
	if len(args) == 0 {
		return
	}
	switch strings.ToLower(args[0]) {
	case "join":
		guildID := m.GuildID
		userID := m.Author.ID
		var channelID string
		guild, err := s.State.Guild(guildID)
		if err == nil {
			for _, vs := range guild.VoiceStates {
				if vs.UserID == userID {
					channelID = vs.ChannelID
					break
				}
			}
		}
		if channelID == "" {
			s.ChannelMessageSend(m.ChannelID, "Tu dois être dans un salon vocal !")
			return
		}
		err = h.voiceManager.Join(context.Background(), guildID, channelID)
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, "Erreur lors de la connexion vocale.")
		} else {
			s.ChannelMessageSend(m.ChannelID, "J'ai rejoint le vocal !")
		}

	case "leave":
		guildID := m.GuildID
		err := h.voiceManager.Leave(context.Background(), guildID)
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, "Erreur lors de la déconnexion vocale.")
		} else {
			s.ChannelMessageSend(m.ChannelID, "J'ai quitté le vocal !")
		}
	}
}
