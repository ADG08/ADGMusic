package application

import (
	"context"
	"fmt"

	"github.com/ADG08/ADGMusic/internal/domain"
	"github.com/ADG08/ADGMusic/internal/ports/output"
)

type MusicService struct {
	repository output.MusicRepository
}

func NewMusicService(repository output.MusicRepository) *MusicService {
	return &MusicService{
		repository: repository,
	}
}

func (s *MusicService) PlayMusic(ctx context.Context, guildID, channelID, url, userID string) error {
	music := domain.NewMusic(guildID, channelID, url, "Titre à extraire", userID)

	queueEntry := &output.MusicQueue{
		GuildID:   music.GuildID,
		ChannelID: music.ChannelID,
		URL:       music.URL,
		Title:     music.Title,
		AddedBy:   music.AddedBy,
	}

	if err := s.repository.AddToQueue(ctx, queueEntry); err != nil {
		return fmt.Errorf("erreur lors de l'ajout à la file d'attente: %w", err)
	}

	return nil
}

func (s *MusicService) StopMusic(ctx context.Context, guildID string) error {
	return nil
}

func (s *MusicService) SkipMusic(ctx context.Context, guildID string) error {
	return nil
}

func (s *MusicService) GetQueue(ctx context.Context, guildID string) ([]*output.MusicQueue, error) {
	return s.repository.GetQueue(ctx, guildID)
}
