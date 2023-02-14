package gender

import (
	"net/http"
	"strconv"

	"github.com/IvSen/shareThings/pkg/response"

	"github.com/IvSen/shareThings/pkg/apperror"

	"github.com/IvSen/shareThings/internal/domain/gender/service"
	"github.com/IvSen/shareThings/internal/jwt"
	"github.com/julienschmidt/httprouter"
)

type Handler struct {
	GenderService *service.GenderService
	JWTHelper     jwt.Helper
}

func NewHandler(
	genderService *service.GenderService,
	JWTHelper jwt.Helper,
) *Handler {
	return &Handler{
		GenderService: genderService,
		JWTHelper:     JWTHelper,
	}
}

func (h *Handler) Register(router *httprouter.Router) {
	router.HandlerFunc(http.MethodGet, "/gender/", jwt.Middleware(apperror.Middleware(h.getAll)))
	router.HandlerFunc(http.MethodGet, "/gender/:id/", jwt.Middleware(apperror.Middleware(h.getById)))
}

func (h *Handler) getAll(w http.ResponseWriter, r *http.Request) error {
	allGender, err := h.GenderService.All(r.Context())
	if err != nil {
		return err
	}

	response.JSON(w, http.StatusOK, allGender)
	return nil
}

func (h *Handler) getById(w http.ResponseWriter, r *http.Request) error {
	//logger := logging.GetLogger()
	params := httprouter.ParamsFromContext(r.Context())
	var id = params.ByName("id")
	var userId, _ = strconv.ParseUint(id, 10, 0)

	gender, err := h.GenderService.One(r.Context(), uint(userId))
	if err != nil {
		return err
	}

	response.JSON(w, http.StatusOK, gender)
	return nil
}
