package user

import (
	"github.com/IvSen/shareThings/internal/domain/user/service"
	"github.com/IvSen/shareThings/internal/jwt"
	"github.com/julienschmidt/httprouter"
)

type Handler struct {
	UserService *service.UserService
	JWTHelper   jwt.Helper
}

func NewHandler(
	userService *service.UserService,
	JWTHelper jwt.Helper,
) *Handler {
	return &Handler{
		UserService: userService,
		JWTHelper:   JWTHelper,
	}
}

func (h *Handler) Register(router *httprouter.Router) {

}
