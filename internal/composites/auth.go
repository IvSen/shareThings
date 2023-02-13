package composites

import (
	"github.com/IvSen/shareThings/internal/controller/http/v1/auth"
	"github.com/IvSen/shareThings/internal/domain/user/dao"
	"github.com/IvSen/shareThings/internal/domain/user/service"
)

type AuthComposite struct {
	UserStorage *dao.UserDAO
	UserService *service.UserService
	AuthHandler *auth.Handler
}

func NewAuthComposite(pcc PgClientComposite, JWTHelper JWTComposite) (AuthComposite, error) {
	var userStorage = dao.NewUserStorage(pcc.Db)
	var userService = service.NewUserService(userStorage)
	var authHandler = auth.NewHandler(userService, JWTHelper.JWT)

	return AuthComposite{
		UserStorage: userStorage,
		UserService: userService,
		AuthHandler: authHandler,
	}, nil
}
