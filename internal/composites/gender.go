package composites

import (
	"github.com/IvSen/shareThings/internal/controller/http/v1/gender"
	"github.com/IvSen/shareThings/internal/domain/gender/dao"
	"github.com/IvSen/shareThings/internal/domain/gender/service"
)

type GenderComposite struct {
	GenderStorage *dao.GenderDAO
	GenderService *service.GenderService
	GenderHandler *gender.Handler
}

func NewGenderComposite(pcc PgClientComposite, JWTHelper JWTComposite) (GenderComposite, error) {
	var genderStorage = dao.NewGenderStorage(pcc.Db)
	var genderService = service.NewGenderService(genderStorage)
	var genderHandler = gender.NewHandler(genderService, JWTHelper.JWT)
	return GenderComposite{
		GenderStorage: genderStorage,
		GenderService: genderService,
		GenderHandler: genderHandler,
	}, nil
}
