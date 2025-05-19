package input

import (
	"context"

	"github.com/ADG08/ADGMusic/internal/domain/entities"
)

type SoundUseCase interface {
	SaveSound(ctx context.Context, name, url string) error
	GetAllSounds(ctx context.Context) ([]*entities.Sound, error)
	FindSoundByID(ctx context.Context, id int64) (*entities.Sound, error)
	DeleteSound(ctx context.Context, id int64) error
	PlaySound(ctx context.Context, id int64) error
}
