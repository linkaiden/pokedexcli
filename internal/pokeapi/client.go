package pokeapi

import (
	"net/http"
	"time"

	"github.com/linkaiden/pokedexcli/internal/pokecache"
)

type Client struct {
	httpClient http.Client
	cache      *pokecache.Cache
}

func NewClient(timeout, reapInterval time.Duration) Client {
	return Client{
		cache: pokecache.NewCache(reapInterval),
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}
