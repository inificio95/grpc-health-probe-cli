package probe

import (
	"errors"
	"sync"
	"time"
)

// CacheConfig controls result caching behaviour between probe attempts.
type CacheConfig struct {
	Enabled bool
	TTL     time.Duration
}

// DefaultCacheConfig returns a CacheConfig with caching disabled.
func DefaultCacheConfig() *CacheConfig {
	return &CacheConfig{
		Enabled: false,
		TTL:     5 * time.Second,
	}
}

// Validate checks that the CacheConfig is well-formed.
func (c *CacheConfig) Validate() error {
	if c == nil {
		return errors.New("cache config must not be nil")
	}
	if c.Enabled && c.TTL <= 0 {
		return errors.New("cache TTL must be positive when caching is enabled")
	}
	return nil
}

// cachedEntry holds a single cached Result and its expiry time.
type cachedEntry struct {
	result    *Result
	expiresAt time.Time
}

// ResultCache is a thread-safe, TTL-based cache for probe Results.
type ResultCache struct {
	mu      sync.Mutex
	entries map[string]*cachedEntry
	ttl     time.Duration
}

// NewResultCache creates a ResultCache with the given TTL.
func NewResultCache(ttl time.Duration) *ResultCache {
	return &ResultCache{
		entries: make(map[string]*cachedEntry),
		ttl:     ttl,
	}
}

// Get returns the cached Result for key and true if it exists and has not expired.
func (rc *ResultCache) Get(key string) (*Result, bool) {
	rc.mu.Lock()
	defer rc.mu.Unlock()
	e, ok := rc.entries[key]
	if !ok || time.Now().After(e.expiresAt) {
		delete(rc.entries, key)
		return nil, false
	}
	return e.result, true
}

// Set stores result under key with the configured TTL.
func (rc *ResultCache) Set(key string, result *Result) {
	rc.mu.Lock()
	defer rc.mu.Unlock()
	rc.entries[key] = &cachedEntry{
		result:    result,
		expiresAt: time.Now().Add(rc.ttl),
	}
}

// Invalidate removes the entry for key from the cache.
func (rc *ResultCache) Invalidate(key string) {
	rc.mu.Lock()
	defer rc.mu.Unlock()
	delete(rc.entries, key)
}
