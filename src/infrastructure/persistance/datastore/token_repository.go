package datastore

import (
	"context"
	"gorm.io/gorm"
	"web-shop/domain"
)

type TokenRepository interface {
	SaveToken(ctx context.Context, token string, userId string) error
	FetchToken(ctx context.Context, userId string) (string, error)
}

type tokenRepository struct {
	Conn *gorm.DB
}

func (t tokenRepository) SaveToken(ctx context.Context, token string, userId string) error {
	adminToken := domain.AdminToken{Token: token, UserId: userId}
	var tokenToDelete *domain.AdminToken
	err := t.Conn.Where("user_id = ?", userId).Take(&tokenToDelete).Error
	if err == nil {
		t.Conn.Delete(&tokenToDelete)
	}

	return t.Conn.Save(&adminToken).Error
}

func (t tokenRepository) FetchToken(ctx context.Context, userId string) (string, error) {
	var token *domain.AdminToken
	err := t.Conn.Where("user_id = ?", userId).Take(&token).Error
	if err != nil {
		return "", err
	}
	return token.Token, nil
}

func NewTokenRepository(conn *gorm.DB) TokenRepository {
	return &tokenRepository{Conn: conn}
}