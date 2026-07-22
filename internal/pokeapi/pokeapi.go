package pokeapi

import "time"

const (
	baseURL                  = "https://pokeapi.co/api/v2"
	DefaultClientTimeout     = (5 * time.Second)
	DefaultCacheReapInterval = (5 * time.Minute)
)
