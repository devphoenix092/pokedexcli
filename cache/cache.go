package cache

import (
	"pokedexcli/cache/pokecache"
)

func Add(key string, val []byte) {
	// Check cache
	pokecache.ReapLoop()

	// Add new pokecache
	pokecache.NewCache(key, val)
}

func Get(key string) (val []byte, isExists bool) {
	// Check cache
	pokecache.ReapLoop()

	// Find the cache using key
	for _, item := range pokecache.PokeCache {
		if key == item.Key {
			return item.Val, true
		}
	}
	return nil, false
}
