// Package cachet is a tiny, generic cache.
package cachet

import (
	"fmt"
	"sync"
)

// Cache is a generic cache for a single value. It must not be mutated after first use. All methods are concurrency
// safe.
type Cache[T any] struct {
	// Load loads the value to be cached. It is required.
	Load func() (T, error)

	// IsStale is called before retrieving the value from the cache. If it returns true, Load will be called.
	//
	// IsStale is optional. If it is not provided the cached value will never be automatically reloaded.
	IsStale func() (bool, error)

	mutex  sync.Mutex
	value  T
	loaded bool
}

// Get gets the value from the cache. It will reload the value if it is stale.
func (c *Cache[T]) Get() (T, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	var zero T

	shouldLoad, err := c.shouldLoad()
	if err != nil {
		return zero, fmt.Errorf("cachet.Cache IsStale: %w", err)
	}

	if shouldLoad {
		value, err := c.Load()
		if err != nil {
			return zero, fmt.Errorf("cachet.Cache Load: %w", err)
		}
		c.value = value
		c.loaded = true
	}

	return c.value, nil
}

// MustGet gets a value from the cache. It will reload the value if it is stale. It panics if an error occurs.
func (c *Cache[T]) MustGet() T {
	tmpl, err := c.Get()
	if err != nil {
		panic(err)
	}
	return tmpl
}

func (c *Cache[T]) shouldLoad() (bool, error) {
	if !c.loaded {
		return true, nil
	}

	if c.IsStale == nil {
		return false, nil
	}

	return c.IsStale()
}
