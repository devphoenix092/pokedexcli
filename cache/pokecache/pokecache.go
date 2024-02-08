package pokecache

import "time"

type CacheEntry struct {
	Key       string
	CreatedAt time.Time
	Val       []byte
}

const CACHE_DURATION = 5 // 5s

var PokeCache []CacheEntry

func ReapLoop() {
	// Remove items that expired
	currentTime := time.Now()
	for i := 0; i < len(PokeCache); i++ {
		diffTime := currentTime.Sub(PokeCache[i].CreatedAt)
		if diffTime.Seconds() > CACHE_DURATION {
			PokeCache = append(PokeCache[:i], PokeCache[i+1:]...)
			i--
		}
	}
}

func NewCache(key string, val []byte) {
	newCache := CacheEntry{
		Key:       key,
		Val:       val,
		CreatedAt: time.Now(),
	}

	// Remove item that has same key
	for index, item := range PokeCache {
		if item.Key == key {
			PokeCache = append(PokeCache[:index], PokeCache[index+1:]...)
		}
	}

	// Add new item to cache
	PokeCache = append(PokeCache, newCache)
}
