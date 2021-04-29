package usecase

import (
	"github.com/labstack/echo"
	"web-shop/domain"
)

type roleUsecase struct {
	RoleRepository domain.RoleRepository
}

func NewRoleUsecase(r domain.RoleRepository) domain.RoleUsecase {
	return &roleUsecase{RoleRepository: r}
}

func (r *roleUsecase) Fetch(ctx echo.Context) ([]*domain.Role, error) {
	return r.RoleRepository.Fetch()
}

func (r *roleUsecase) GetByID(ctx echo.Context, id uint) (*domain.Role, error) {
	return r.RoleRepository.GetByID(id)
}

func (r *roleUsecase) Update(ctx echo.Context, o *domain.Role) (*domain.Role, error) {
	return r.RoleRepository.Update(o)
}

func (r *roleUsecase) Create(ctx echo.Context, o *domain.Role) (*domain.Role, error) {
	return r.RoleRepository.Create(o)
}

func (r *roleUsecase) Delete(ctx echo.Context, id uint) error {
	return r.RoleRepository.Delete(id)
}