package datastore

import (
	"gorm.io/gorm"
	"web-shop/domain"
)

type roleRepository struct {
	Conn *gorm.DB
}

func NewRoleRepository(Conn *gorm.DB) domain.RoleRepository {
	return &roleRepository{Conn: Conn}
}


func (r *roleRepository) Fetch() ([]*domain.Role, error) {
	var (
		roles []*domain.Role
		err error
	)

	err = r.Conn.Order("id desc").Find(&roles).Error
	return roles, err
}

func (r *roleRepository) GetByID(id uint) (*domain.Role, error) {
	role := &domain.Role{Model: gorm.Model{ID: id}}
	err := r.Conn.First(role).Error
	return role, err
}

func (r *roleRepository) Update(o *domain.Role) (*domain.Role, error) {
	err := r.Conn.Save(o).Error
	return o, err
}

func (r roleRepository) Create(o *domain.Role) (*domain.Role, error) {
	err := r.Conn.Create(o).Error
	return  o, err
}

func (r *roleRepository) Delete(id uint) error {
	role := &domain.Role{Model: gorm.Model{ID: id}}
	err := r.Conn.Delete(role).Error
	return err
}
