package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

// User define la estructura del usuario para la inserción en la base de datos
type UserEntity struct {
    bun.BaseModel `bun:"table:public.users,alias:u"`
    ID            uuid.UUID  `bun:"id,pk,autoincrement"`
    Name          string `bun:"name"`
    Lastname      string `bun:"lastname"`
    Age           int    `bun:"age"`
    Email         string `bun:"email,unique"`
}

// UserSchema define la estructura completa del usuario
type UserResponse struct {
    bun.BaseModel `bun:"table:public.users,alias:u" swaggerignore:"true"`

    ID       uuid.UUID   `json:"id,omitempty" bun:"id" description:"ID of the user"`
    Name     string  `json:"name" bun:"name" description:"Name of the user"`
    Lastname string  `json:"lastname" bun:"lastname" description:"Lastname of the user"`
    Age      int     `json:"age" bun:"age" description:"Age of the user"`
    Email    string  `json:"email" bun:"email" description:"Email of the user"`
}

// UserCreateSchema define la estructura para la creación de usuarios
type UserCreateSchema struct {
    Name     string `json:"name" description:"Name of the user"`
    Lastname string `json:"lastname" description:"Lastname of the user"`
    Age      int    `json:"age" description:"Age of the user"`
    Email    string `json:"email" description:"Email of the user"`
}

// UserUpdateSchema define la estructura para la actualización de usuarios
type UserUpdateSchema struct {
    Name     *string `json:"name,omitempty" description:"Name of the user"`
    Lastname *string `json:"lastname,omitempty" description:"Lastname of the user"`
    Age      *int    `json:"age,omitempty" description:"Age of the user"`
    Email    *string `json:"email,omitempty" description:"Email of the user"`
}

type NotFountResponse struct {
    Msg string `json:"msg"`
}