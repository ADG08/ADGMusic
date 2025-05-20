package output

import (
	"context"

	"github.com/ADG08/ADGMusic/internal/domain/entities"
)

type SoundRepository interface {
	Save(ctx context.Context, sound *entities.Sound) error
	GetAll(ctx context.Context) ([]*entities.Sound, error)
	FindByID(ctx context.Context, id int64) (*entities.Sound, error)
	Delete(ctx context.Context, id int64) error
}

type SoundPlayer interface {
	Play(ctx context.Context, sound *entities.Sound) error
}
