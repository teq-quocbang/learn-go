package auth

import "gorm.io/gorm"

type AuthRepo struct {
	client *gorm.DB
}

func NewAuthRepo(c *gorm.DB) *AuthRepo {
	return &AuthRepo{
		client: c,
	}
}

func (ar *AuthRepo) Login() error {
	return nil
}
