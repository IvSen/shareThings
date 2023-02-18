package category

import (
	"net/http"

	"github.com/IvSen/shareThings/internal/domain/category/service"
	"github.com/IvSen/shareThings/internal/jwt"
	"github.com/IvSen/shareThings/pkg/apperror"
	"github.com/IvSen/shareThings/pkg/response"
	"github.com/julienschmidt/httprouter"
)

type Handler struct {
	CategoryService *service.CategoryService
	JWTHelper       jwt.Helper
}

func NewHandler(
	categoryService *service.CategoryService,
	JWTHelper jwt.Helper,
) *Handler {
	return &Handler{
		CategoryService: categoryService,
		JWTHelper:       JWTHelper,
	}
}

func (h *Handler) Register(router *httprouter.Router) {
	router.HandlerFunc(http.MethodGet, "/category/", jwt.Middleware(apperror.Middleware(h.getAll)))
	router.HandlerFunc(http.MethodGet, "/category/:id/", jwt.Middleware(apperror.Middleware(h.getById)))
}

func (h *Handler) getAll(w http.ResponseWriter, r *http.Request) error {
	allGender, err := h.CategoryService.All(r.Context())
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
	gender, err := h.CategoryService.One(r.Context(), id)
	if err != nil {
		return err
	}

	response.JSON(w, http.StatusOK, gender)
	return nil
}
