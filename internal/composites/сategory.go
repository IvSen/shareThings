package composites

import (
	"github.com/IvSen/shareThings/internal/controller/http/v1/category"
	"github.com/IvSen/shareThings/internal/domain/category/dao"
	"github.com/IvSen/shareThings/internal/domain/category/service"
)

type CategoryComposite struct {
	CategoryStorage *dao.CategoryDAO
	CategoryService *service.CategoryService
	CategoryHandler *category.Handler
}

func NewCategoryComposite(pcc PgClientComposite, JWTHelper JWTComposite) (CategoryComposite, error) {
	var categoryStorage = dao.NewCategoryStorage(pcc.Db)
	var categoryService = service.NewCategoryService(categoryStorage)
	var categoryHandler = category.NewHandler(categoryService, JWTHelper.JWT)
	return CategoryComposite{
		CategoryStorage: categoryStorage,
		CategoryService: categoryService,
		CategoryHandler: categoryHandler,
	}, nil
}
