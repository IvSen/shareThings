package composites

import (
	"github.com/IvSen/shareThings/pkg/cache"
	"github.com/IvSen/shareThings/pkg/cache/freecache"
)

type CacheComposite struct {
	Cache cache.Repository
}

func NewCacheComposite(cacheSize int) (CacheComposite, error) {
	return CacheComposite{Cache: freecache.NewCacheRepo(cacheSize)}, nil
}
