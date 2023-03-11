[![Go Reference](https://pkg.go.dev/badge/github.com/jackc/cachet.svg)](https://pkg.go.dev/github.com/jackc/cachet)
![Build Status](https://github.com/jackc/cachet/actions/workflows/ci.yml/badge.svg)

# CacheT

CacheT is a tiny, generic cache. It was initially designed to assist with live reloading HTML templates in development
while caching them in production.

## Example Usage

```go
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

cache.MustGet() // => 1
cache.MustGet() // => 1
isStale = true
cache.MustGet() // => 2
cache.MustGet() // => 3
isStale = false
cache.MustGet() // => 3
```
