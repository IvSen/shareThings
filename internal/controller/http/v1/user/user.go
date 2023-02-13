package user

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/IvSen/shareThings/pkg/logging"

	"github.com/IvSen/shareThings/pkg/response"

	"github.com/IvSen/shareThings/internal/domain/user/model"

	"github.com/IvSen/shareThings/internal/domain/user/service"
	"github.com/IvSen/shareThings/internal/jwt"
	"github.com/IvSen/shareThings/pkg/apperror"
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
	router.HandlerFunc(http.MethodPut, "/user/", jwt.Middleware(apperror.Middleware(h.updateUser)))
}

func (h *Handler) updateUser(w http.ResponseWriter, r *http.Request) error {
	logger := logging.GetLogger()
	//params := httprouter.ParamsFromContext(r.Context())
	//logger.WithFields(map[string]interface{}{
	//	"Method": r.Method,
	//	"Path":   r.URL.Path,
	//	"Params": params,
	//})
	var user model.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		logger.Error(r.Context(), err)
		// TODO: log
		//controller.NewErrorResponse(gctx, http.StatusBadRequest, "invalid input body")
		//return errors.New("invalid input body")
		return errors.New(http.StatusText(http.StatusBadRequest))
	}
	user.Id = r.Context().Value("user_id").(string)

	modelUser, err := h.UserService.Update(r.Context(), &user)
	if err != nil {
		logger.Error(r.Context(), err)
		return err
	}

	response.JSON(w, http.StatusOK, modelUser)
	return nil
}
