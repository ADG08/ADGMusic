package services

import (
	"context"

	"github.com/ADG08/ADGMusic/internal/domain/entities"
	"github.com/ADG08/ADGMusic/internal/domain/errors"
	"github.com/ADG08/ADGMusic/internal/ports/input"
	"github.com/ADG08/ADGMusic/internal/ports/output"
)

type SoundService struct {
	repository output.SoundRepository
	player     output.SoundPlayer
}

func NewSoundService(repository output.SoundRepository, player output.SoundPlayer) input.SoundUseCase {
	return &SoundService{
		repository: repository,
		player:     player,
	}
}

func (s *SoundService) SaveSound(ctx context.Context, name, url string) error {
	if name == "" || url == "" {
		return errors.ErrInvalidInput
	}

	return s.repository.Save(ctx, &entities.Sound{Name: name, URL: url})
}

func (s *SoundService) GetAllSounds(ctx context.Context) ([]*entities.Sound, error) {
	return s.repository.GetAll(ctx)
}

func (s *SoundService) FindSoundByID(ctx context.Context, id int64) (*entities.Sound, error) {
	return s.repository.FindByID(ctx, id)
}

func (s *SoundService) DeleteSound(ctx context.Context, id int64) error {
	return s.repository.Delete(ctx, id)
}

func (s *SoundService) PlaySound(ctx context.Context, id int64) error {
	sound, err := s.FindSoundByID(ctx, id)
	if err != nil {
		return err
	}
	return s.player.Play(ctx, sound)
}
