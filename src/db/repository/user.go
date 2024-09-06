package repository

import (
	"context"
	"development-environment-api-go-manager/src/models"

	"github.com/uptrace/bun"
)

type user struct {
	con *bun.DB
}

func NewUserRepository(con *bun.DB) *user {
	return &user{con: con,}
}

func (userRepository *user) ReadUser(id int64) (*models.UserResponse, error) {
	model := new(models.UserResponse)
	
	data := userRepository.con.NewSelect().
	Model(model).
	Column("u.*"). // Usar el alias configurado en la estructura
	Where("u.id = ?", id)
	err := data.Scan(context.Background())

	if err != nil {
		return nil, err
	}
	
	return model, nil
}