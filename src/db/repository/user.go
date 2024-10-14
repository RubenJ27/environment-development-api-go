package repository

import (
	"context"
	"database/sql"
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

func (userRepository *user) CreateUser(newUser models.UserCreateSchema) (*models.UserResponse, error) {
    // Crear una instancia de models.UserEntity para la inserción en la base de datos
    user := &models.UserEntity{
        Name:     newUser.Name,
        Lastname: newUser.Lastname,
        Age:      newUser.Age,
        Email:    newUser.Email,
    }

    // Insertar el nuevo usuario en la base de datos
    _, err := userRepository.con.NewInsert().Model(user).Exec(context.Background())
    if err != nil {
        return nil, err
    }

    // Crear una instancia de models.UserResponse para la respuesta
    userResponse := &models.UserResponse{
        ID:       user.ID,
        Name:     user.Name,
        Lastname: user.Lastname,
        Age:      user.Age,
        Email:    user.Email,
    }

    return userResponse, nil
}

func (userRepository *user) UpdateUser(id int64, updateUser models.UserUpdateSchema) (*models.UserResponse, error) {
    // Verificar si el usuario existe
    user := &models.UserEntity{}
    err := userRepository.con.NewSelect().Model(user).Where("id = ?", id).Scan(context.Background())
    if err != nil {
        if err == sql.ErrNoRows {
            return nil, ErrUserNotFound
        }
        return nil, err
    }

    // Ejecutar la actualización
    _, err = userRepository.con.NewUpdate().
        Model(user).
        Set("name = COALESCE(?, name)", updateUser.Name).
        Set("lastname = COALESCE(?, lastname)", updateUser.Lastname).
        Set("age = COALESCE(?, age)", updateUser.Age).
        Set("email = COALESCE(?, email)", updateUser.Email).
        Where("id = ?", id).
        Exec(context.Background())
    if err != nil {
        return nil, err
    }

    // Crear una instancia de models.UserResponse para la respuesta
    userResponse := &models.UserResponse{
        ID:       user.ID,
        Name:     *updateUser.Name,
        Lastname: *updateUser.Lastname,
        Age:      *updateUser.Age,
        Email:    *updateUser.Email,
    }

    return userResponse, nil
}

func (userRepository *user) PartialUpdateUser(id int64, updateUser models.UserUpdateSchema) (*models.UserResponse, error) {
    // Verificar si el usuario existe
    user := &models.UserEntity{}
    err := userRepository.con.NewSelect().Model(user).Where("id = ?", id).Scan(context.Background())
    if err != nil {
        if err == sql.ErrNoRows {
            return nil, ErrUserNotFound
        }
        return nil, err
    }

    // Ejecutar la actualización parcial
    _, err = userRepository.con.NewUpdate().
        Model(user).
        Set("name = COALESCE(?, name)", updateUser.Name).
        Set("lastname = COALESCE(?, lastname)", updateUser.Lastname).
        Set("age = COALESCE(?, age)", updateUser.Age).
        Set("email = COALESCE(?, email)", updateUser.Email).
        Where("id = ?", id).
        Exec(context.Background())
    if err != nil {
        return nil, err
    }

    // Crear una instancia de models.UserResponse para la respuesta
    userResponse := &models.UserResponse{
        ID:       user.ID,
        Name:     user.Name,
        Lastname: user.Lastname,
        Age:      user.Age,
        Email:    user.Email,
    }

    return userResponse, nil
}