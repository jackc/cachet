package cachet_test

import (
	"testing"

	"github.com/jackc/cachet"
	"github.com/stretchr/testify/require"
)

func TestCachetCacheLoadedValueIsCached(t *testing.T) {
	callCount := 0
	cache := cachet.Cache[int]{
		Load: func() (int, error) {
			callCount++
			return 42, nil
		},
	}

	require.Equal(t, 42, cache.MustGet())
	require.Equal(t, 42, cache.MustGet())
	require.Equal(t, 1, callCount)
}

func TestCachetCacheUsesIsStaleToControlReload(t *testing.T) {
	value := 0
	isStale := false
	cache := cachet.Cache[int]{
		Load: func() (int, error) {
			value++
			return value, nil
		},
		IsStale: func() (bool, error) {
			return isStale, nil
		},
	}

	require.Equal(t, 1, cache.MustGet())
	require.Equal(t, 1, cache.MustGet())
	isStale = true
	require.Equal(t, 2, cache.MustGet())
	require.Equal(t, 3, cache.MustGet())
	isStale = false
	require.Equal(t, 3, cache.MustGet())
}
