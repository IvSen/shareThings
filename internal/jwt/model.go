package jwt

import (
	"github.com/IvSen/shareThings/pkg/cache"
	"github.com/IvSen/shareThings/pkg/logging"
	"github.com/cristalhq/jwt/v3"
)

type User struct {
	UUID     string
	Email    string
	Password string
}

type RT struct {
	RefreshToken string `json:"refresh_token"`
}

type UserClaims struct {
	jwt.RegisteredClaims
	Email string `json:"email"`
}

type helper struct {
	Logger  logging.Logger
	RTCache cache.Repository
}
