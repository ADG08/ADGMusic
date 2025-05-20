package services

import (
	"context"
	"math/rand"
	"time"

	"github.com/ADG08/ADGMusic/internal/domain/entities"
	"github.com/ADG08/ADGMusic/internal/ports/output"
)

type ActiveChannelService struct {
	repo output.ActiveChannelRepository
}

func (s *ActiveChannelService) UserJoined(ctx context.Context, guildID, channelID string) error {
	return s.repo.Add(ctx, &entities.ActiveChannel{
		GuildID:   guildID,
		ChannelID: channelID,
		CreatedAt: time.Now(),
	})
}

func (s *ActiveChannelService) UserLeft(ctx context.Context, guildID, channelID string) error {
	return s.repo.Remove(ctx, &entities.ActiveChannel{
		GuildID:   guildID,
		ChannelID: channelID,
	})
}

func (s *ActiveChannelService) GetRandomActiveChannel(ctx context.Context) (*entities.ActiveChannel, error) {
	channels, err := s.repo.GetAll(ctx)
	if err != nil || len(channels) == 0 {
		return nil, err
	}

	return channels[rand.Intn(len(channels))], nil
}
