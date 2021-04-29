package usecase



import (
	"github.com/labstack/echo"
	"web-shop/domain"
)

type categoryUseCase struct {
	CategoryRepository domain.CategoryRepository
}

func (c *categoryUseCase) GetByName(name string) (*domain.Category, error) {
	return c.CategoryRepository.GetByName(name)
}

func (c *categoryUseCase) Fetch(ctx echo.Context) ([]*domain.Category, error) {
	return c.CategoryRepository.Fetch()
}
func (c *categoryUseCase) GetByID(ctx echo.Context, id uint) (*domain.Category, error) {
	return c.CategoryRepository.GetByID(id)
}

func (c *categoryUseCase) Update(ctx echo.Context, pic *domain.Category) (*domain.Category, error) {
	return c.CategoryRepository.Update(pic)
}

func (c *categoryUseCase) Create(ctx echo.Context, pic *domain.Category) (*domain.Category, error) {
	return c.CategoryRepository.Create(pic)
}

func (c *categoryUseCase) Delete(ctx echo.Context, id uint) error {
	return c.CategoryRepository.Delete(id)
}

func NewCategoryUsecase(r domain.CategoryRepository) domain.CategoryUsecase {
	return &categoryUseCase{r}
}
