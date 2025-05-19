package discord

import (
	"fmt"
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
)

type CommandHandler struct {
	session *discordgo.Session
	// Map pour stocker les connexions vocales actives
	voiceConnections map[string]*discordgo.VoiceConnection
}

func NewCommandHandler(session *discordgo.Session) *CommandHandler {
	return &CommandHandler{
		session:          session,
		voiceConnections: make(map[string]*discordgo.VoiceConnection),
	}
}

func (h *CommandHandler) RegisterCommands() {
	h.session.AddHandler(h.handleMessageCreate)
	h.session.AddHandler(h.handleVoiceStateUpdate)
}

func (h *CommandHandler) handleVoiceStateUpdate(s *discordgo.Session, vs *discordgo.VoiceStateUpdate) {
	// Si l'utilisateur a quitté un canal vocal
	if vs.BeforeUpdate != nil && vs.BeforeUpdate.ChannelID != "" && vs.ChannelID == "" {
		guildID := vs.GuildID
		channelID := vs.BeforeUpdate.ChannelID

		// Vérifier si le bot est dans ce canal
		if vc, ok := h.voiceConnections[guildID]; ok && vc.ChannelID == channelID {
			// Obtenir tous les utilisateurs dans le canal
			channel, err := s.Channel(channelID)
			if err != nil {
				log.Printf("Erreur lors de la récupération du canal: %v", err)
				return
			}

			userCount := 0
			guild, err := s.Guild(guildID)
			if err != nil {
				log.Printf("Erreur lors de la récupération du serveur: %v", err)
				return
			}

			for _, voiceState := range guild.VoiceStates {
				if voiceState.ChannelID == channelID && voiceState.UserID != s.State.User.ID {
					userCount++
				}
			}

			if userCount == 0 {
				log.Printf("Déconnexion du canal vocal %s car vide", channel.Name)
				vc.Disconnect()
				delete(h.voiceConnections, guildID)
			}
		}
	}
}

func (h *CommandHandler) handleMessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	log.Printf("Message reçu: %s", m.Content)

	if m.Author.ID == s.State.User.ID {
		return
	}

	if !strings.HasPrefix(m.Content, "!") {
		return
	}

	log.Printf("Commande détectée: %s", m.Content)

	args := strings.Fields(m.Content)
	command := strings.ToLower(args[0][1:])

	switch command {
	case "play":
		if len(args) < 2 {
			s.ChannelMessageSend(m.ChannelID, "Usage: !play <url>")
			return
		}
		url := args[1]
		h.handlePlayCommand(s, m, url)
	case "stop":
		h.handleStopCommand(s, m)
	case "skip":
		h.handleSkipCommand(s, m)
	default:
		s.ChannelMessageSend(m.ChannelID, "Commande inconnue. Commandes disponibles: !play, !stop, !skip, !test")
	}
}

func (h *CommandHandler) handlePlayCommand(s *discordgo.Session, m *discordgo.MessageCreate, url string) {
	// Vérifier si l'utilisateur est dans un canal vocal
	voiceState, err := s.State.VoiceState(m.GuildID, m.Author.ID)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "Tu dois être dans un canal vocal pour utiliser cette commande.")
		return
	}

	// Rejoindre le canal vocal
	vc, err := s.ChannelVoiceJoin(m.GuildID, voiceState.ChannelID, false, true)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "Impossible de rejoindre le canal vocal.")
		return
	}

	// Stocker la connexion vocale
	h.voiceConnections[m.GuildID] = vc

	// TODO: Implémenter la logique de lecture de musique
	s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Tentative de lecture de: %s", url))
}

func (h *CommandHandler) handleStopCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSend(m.ChannelID, "Arrêt de la lecture...")
}

func (h *CommandHandler) handleSkipCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSend(m.ChannelID, "Passage à la musique suivante...")
}
