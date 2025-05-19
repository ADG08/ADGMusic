package output

import (
	"context"
	"time"
)

// MusicQueue représente une entrée dans la file d'attente de musique
type MusicQueue struct {
	ID        int64
	GuildID   string
	ChannelID string
	URL       string
	Title     string
	AddedBy   string
	CreatedAt time.Time
}

// MusicRepository définit les opérations de persistance pour la musique
type MusicRepository interface {
	// AddToQueue ajoute une musique à la file d'attente
	AddToQueue(ctx context.Context, queue *MusicQueue) error

	// GetQueue récupère la file d'attente pour un serveur
	GetQueue(ctx context.Context, guildID string) ([]*MusicQueue, error)

	// RemoveFromQueue retire une musique de la file d'attente
	RemoveFromQueue(ctx context.Context, id int64) error

	// ClearQueue vide la file d'attente d'un serveur
	ClearQueue(ctx context.Context, guildID string) error
}
