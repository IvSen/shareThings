package composites

import (
	"github.com/IvSen/shareThings/internal/jwt"
	"github.com/IvSen/shareThings/pkg/logging"
)

type JWTComposite struct {
	JWT jwt.Helper
}

func NewJWTComposite(cacheComposite CacheComposite, logger logging.Logger) (JWTComposite, error) {
	return JWTComposite{JWT: jwt.NewHelper(cacheComposite.Cache, logger)}, nil
}
