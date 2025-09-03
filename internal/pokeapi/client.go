package pokeapi

import (
	"net/http"
	"time"

	"github.com/hectoribarra2024-eng/pokedex_go/internal/pokecache"
)

// Client -
type Client struct {
	httpClient 	http.Client 
	cache		pokecache.Cache
}

// NewClient -
func NewClient(timeout, cacheInterval time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		cache: pokecache.NewCache(cacheInterval),
	}
}
