package auth

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/IvSen/shareThings/internal/controller/http/v1/dto"

	"github.com/IvSen/shareThings/internal/domain/user/service"
	"github.com/IvSen/shareThings/internal/jwt"
	"github.com/IvSen/shareThings/pkg/apperror"
	"github.com/julienschmidt/httprouter"
)

const (
	authURL   = "/api/auth"
	signupURL = "/api/signup"
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
	router.HandlerFunc(http.MethodPost, authURL, apperror.Middleware(h.Auth))
	router.HandlerFunc(http.MethodPut, authURL, apperror.Middleware(h.Auth))
	router.HandlerFunc(http.MethodPost, signupURL, apperror.Middleware(h.Signup))
}

func (h *Handler) Signup(w http.ResponseWriter, r *http.Request) error {
	defer r.Body.Close()
	var userDto dto.CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&userDto); err != nil {
		return errors.New("failed to decode data")
	}

	u, err := h.UserService.Create(r.Context(), &userDto)
	if err != nil {
		return err
	}
	token, err := h.JWTHelper.GenerateAccessToken(jwt.User{
		UUID:     u.Id,
		Email:    u.Email,
		Password: h.UserService.GeneratePasswordHash(u.Password),
	})
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(token)

	return nil
}

func (h *Handler) Auth(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "application/json")

	var token []byte
	var err error
	switch r.Method {
	case http.MethodPost:
		defer r.Body.Close()
		var dto dto.SignInUserDTO
		if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
			return errors.New("failed to decode data")
		}
		u, err := h.UserService.GetByEmailAndPassword(r.Context(), dto.Email, dto.Password)
		if err != nil {
			return err
		}
		token, err = h.JWTHelper.GenerateAccessToken(jwt.User{
			UUID:     u.Id,
			Email:    u.Email,
			Password: h.UserService.GeneratePasswordHash(u.Password),
		})
		if err != nil {
			return err
		}
	case http.MethodPut:
		defer r.Body.Close()
		var rt jwt.RT
		if err := json.NewDecoder(r.Body).Decode(&rt); err != nil {
			return errors.New("failed to decode data")
		}
		token, err = h.JWTHelper.UpdateRefreshToken(rt)
		if err != nil {
			return err
		}
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(token)

	return err
}
