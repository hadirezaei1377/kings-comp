package repository

import (
	"kings-comp/internal/entity"

	"github.com/redis/rueidis"
)

var _ LobbyRepository = &LobbyRedisRepository{}

type LobbyRedisRepository struct {
	*RedisCommonBehaviour[entity.Lobby]
}

func NewLobbyRedisRepository(client rueidis.Client) *LobbyRedisRepository {
	return &LobbyRedisRepository{
		NewRedisCommonBehaviour[entity.Lobby](client),
	}
}
