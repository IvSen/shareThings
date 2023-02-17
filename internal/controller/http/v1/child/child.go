package child

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/IvSen/shareThings/internal/controller/http/v1/dto"

	"github.com/IvSen/shareThings/pkg/response"

	"github.com/IvSen/shareThings/internal/domain/child/service"
	"github.com/IvSen/shareThings/internal/jwt"
	"github.com/IvSen/shareThings/pkg/apperror"
	"github.com/julienschmidt/httprouter"
)

type Handler struct {
	ChildService *service.ChildService
	JWTHelper    jwt.Helper
}

func NewHandler(
	ChildService *service.ChildService,
	JWTHelper jwt.Helper,
) *Handler {
	return &Handler{
		ChildService: ChildService,
		JWTHelper:    JWTHelper,
	}
}

func (h *Handler) Register(router *httprouter.Router) {
	router.HandlerFunc(http.MethodGet, "/child/", jwt.Middleware(apperror.Middleware(h.getAll)))
	router.HandlerFunc(http.MethodPost, "/child/", jwt.Middleware(apperror.Middleware(h.create)))
	router.HandlerFunc(http.MethodPut, "/child/:id/", jwt.Middleware(apperror.Middleware(h.update)))
	router.HandlerFunc(http.MethodGet, "/child/:id/", jwt.Middleware(apperror.Middleware(h.getById)))
	router.HandlerFunc(http.MethodDelete, "/child/:id/", jwt.Middleware(apperror.Middleware(h.delete)))
}

func (h *Handler) update(w http.ResponseWriter, r *http.Request) error {
	//logger := logging.GetLogger()
	params := httprouter.ParamsFromContext(r.Context())

	var dtoData dto.CreateUpdateChildRequest

	if err := json.NewDecoder(r.Body).Decode(&dtoData); err != nil {
		//logger.Error(r.Context(), err)
		// TODO: log
		//controller.NewErrorResponse(gctx, http.StatusBadRequest, "invalid input body")
		//return errors.New("invalid input body")
		return errors.New(http.StatusText(http.StatusBadRequest))
	}
	dtoData.Id = params.ByName("id")
	modelUser, err := h.ChildService.Update(r.Context(), &dtoData)
	if err != nil {
		//logger.Error(r.Context(), err)
		return err
	}

	response.JSON(w, http.StatusOK, modelUser)
	return nil
}

func (h *Handler) create(w http.ResponseWriter, r *http.Request) error {
	defer r.Body.Close()
	var dtoData dto.CreateUpdateChildRequest
	if err := json.NewDecoder(r.Body).Decode(&dtoData); err != nil {
		return errors.New("failed to decode data")
	}

	e, err := h.ChildService.Create(r.Context(), &dtoData)
	if err != nil {
		return err
	}

	response.JSON(w, http.StatusOK, e)

	return nil
}

func (h *Handler) getById(w http.ResponseWriter, r *http.Request) error {
	params := httprouter.ParamsFromContext(r.Context())
	var id = params.ByName("id")
	gender, err := h.ChildService.One(r.Context(), id)
	if err != nil {
		return err
	}

	response.JSON(w, http.StatusOK, gender)
	return nil
}

func (h *Handler) delete(w http.ResponseWriter, r *http.Request) error {
	params := httprouter.ParamsFromContext(r.Context())
	var id = params.ByName("id")
	err := h.ChildService.Delete(r.Context(), id)
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusOK)

	return nil
}

func (h *Handler) getAll(w http.ResponseWriter, r *http.Request) error {
	all, err := h.ChildService.All(r.Context())
	if err != nil {
		return err
	}

	response.JSON(w, http.StatusOK, all)
	return nil
}
