package probe

import (
	"testing"
	"time"
)

func TestDefaultCacheConfig(t *testing.T) {
	cfg := DefaultCacheConfig()
	if cfg == nil {
		t.Fatal("expected non-nil config")
	}
	if cfg.Enabled {
		t.Error("expected caching to be disabled by default")
	}
	if cfg.TTL != 5*time.Second {
		t.Errorf("expected default TTL 5s, got %v", cfg.TTL)
	}
}

func TestCacheConfig_Validate_Nil(t *testing.T) {
	var cfg *CacheConfig
	if err := cfg.Validate(); err == nil {
		t.Error("expected error for nil config")
	}
}

func TestCacheConfig_Validate_Disabled(t *testing.T) {
	cfg := &CacheConfig{Enabled: false, TTL: 0}
	if err := cfg.Validate(); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestCacheConfig_Validate_EnabledPositiveTTL(t *testing.T) {
	cfg := &CacheConfig{Enabled: true, TTL: 10 * time.Second}
	if err := cfg.Validate(); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestCacheConfig_Validate_EnabledZeroTTL(t *testing.T) {
	cfg := &CacheConfig{Enabled: true, TTL: 0}
	if err := cfg.Validate(); err == nil {
		t.Error("expected error for zero TTL with caching enabled")
	}
}

func TestResultCache_GetMiss(t *testing.T) {
	rc := NewResultCache(time.Second)
	_, ok := rc.Get("missing")
	if ok {
		t.Error("expected cache miss")
	}
}

func TestResultCache_SetAndGet(t *testing.T) {
	rc := NewResultCache(time.Second)
	r := &Result{Status: StatusServing}
	rc.Set("svc", r)
	got, ok := rc.Get("svc")
	if !ok {
		t.Fatal("expected cache hit")
	}
	if got.Status != StatusServing {
		t.Errorf("expected SERVING, got %v", got.Status)
	}
}

func TestResultCache_Expiry(t *testing.T) {
	rc := NewResultCache(10 * time.Millisecond)
	rc.Set("svc", &Result{Status: StatusServing})
	time.Sleep(20 * time.Millisecond)
	_, ok := rc.Get("svc")
	if ok {
		t.Error("expected cache miss after TTL expiry")
	}
}

func TestResultCache_Invalidate(t *testing.T) {
	rc := NewResultCache(time.Second)
	rc.Set("svc", &Result{Status: StatusServing})
	rc.Invalidate("svc")
	_, ok := rc.Get("svc")
	if ok {
		t.Error("expected cache miss after invalidation")
	}
}

func TestResultCache_ConcurrentAccess(t *testing.T) {
	rc := NewResultCache(time.Second)
	done := make(chan struct{})
	for i := 0; i < 10; i++ {
		go func() {
			rc.Set("key", &Result{Status: StatusServing})
			rc.Get("key")
			done <- struct{}{}
		}()
	}
	for i := 0; i < 10; i++ {
		<-done
	}
}
