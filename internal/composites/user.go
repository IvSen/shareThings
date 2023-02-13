package composites

import (
	"github.com/IvSen/shareThings/internal/controller/http/v1/user"
	"github.com/IvSen/shareThings/internal/domain/user/dao"
	"github.com/IvSen/shareThings/internal/domain/user/service"
)

type UserComposite struct {
	UserStorage *dao.UserDAO
	UserService *service.UserService
	UserHandler *user.Handler
}

func NewUserComposite(pcc PgClientComposite, JWTHelper JWTComposite) (UserComposite, error) {
	var userStorage = dao.NewUserStorage(pcc.Db)
	var userService = service.NewUserService(userStorage)
	var userHandler = user.NewHandler(userService, JWTHelper.JWT)
	return UserComposite{
		UserStorage: userStorage,
		UserService: userService,
		UserHandler: userHandler,
	}, nil
}
