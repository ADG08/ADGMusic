package output

import (
	"context"

	"github.com/ADG08/ADGMusic/internal/domain/entities"
)

type ActiveChannelRepository interface {
	Add(ctx context.Context, channel *entities.ActiveChannel) error
	Remove(ctx context.Context, channel *entities.ActiveChannel) error
	GetAll(ctx context.Context) ([]*entities.ActiveChannel, error)
}
