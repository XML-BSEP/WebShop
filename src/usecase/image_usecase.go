package usecase

import (
	"github.com/labstack/echo"
	"web-shop/domain"
)

type imageUseCase struct {
	ImageRepository domain.ImageRepository
}

func (i *imageUseCase) GetyByPath(path string) ([]*domain.Image, error) {
	return i.ImageRepository.GetyByPath(path)
}

func NewImageUseCase(i domain.ImageRepository) domain.ImageUseCase {
	return &imageUseCase{ImageRepository : i}
}

func (i *imageUseCase) Fetch(ctx echo.Context) ([]*domain.Image, error) {
	return i.ImageRepository.Fetch()
}

func (i *imageUseCase) GetByID(ctx echo.Context, id uint) (*domain.Image, error) {
	return i.ImageRepository.GetByID(id)
}

func (i *imageUseCase) Update(ctx echo.Context, o *domain.Image) (*domain.Image, error) {
	return i.ImageRepository.Update(o)
}

func (i *imageUseCase) Create(ctx echo.Context, o *domain.Image) (*domain.Image, error) {
	return i.ImageRepository.Create(o)
}

func (i *imageUseCase) Delete(ctx echo.Context, id uint) error {
	return i.ImageRepository.Delete(id)
}
