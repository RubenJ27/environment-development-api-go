package repository

import (
	"context"
	"development-environment-api-go-manager/src/models"
	"errors"

	"github.com/uptrace/bun"
)

type user struct {
	con *bun.DB
}
var ErrUserNotFound = errors.New("user not found")

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


func (userRepository *user) DeleteUser(id int64) error {
    // Crear un modelo de usuario para la eliminación
    model := &models.UserResponse{ID: id}

    // Ejecutar la eliminación
    result, err := userRepository.con.NewDelete().
        Model(model).
        Where("id = ?", id).
        Exec(context.Background())

    if err != nil {
        return err
    }
	// Verificar si alguna fila fue afectada
    rowsAffected, err := result.RowsAffected()
    if err != nil {
        return err
    }
    if rowsAffected == 0 {
        return ErrUserNotFound
    }

    return nil
}