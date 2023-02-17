package composites

import (
	"github.com/IvSen/shareThings/internal/controller/http/v1/child"
	"github.com/IvSen/shareThings/internal/domain/child/dao"
	"github.com/IvSen/shareThings/internal/domain/child/service"
)

type ChildComposite struct {
	ChildStorage *dao.ChildDAO
	ChildService *service.ChildService
	ChildHandler *child.Handler
}

func NewChildComposite(pcc PgClientComposite, JWTHelper JWTComposite) (ChildComposite, error) {
	var ChildStorage = dao.NewGenderStorage(pcc.Db)
	var ChildService = service.NewChildService(ChildStorage)
	var ChildHandler = child.NewHandler(ChildService, JWTHelper.JWT)
	return ChildComposite{
		ChildStorage: ChildStorage,
		ChildService: ChildService,
		ChildHandler: ChildHandler,
	}, nil
}
