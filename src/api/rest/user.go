package rest

import (
	"development-environment-api-go-manager/src/db/repository"
	"development-environment-api-go-manager/src/models"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserRepository interface {
	ReadUser(id int64) (*models.UserResponse, error)
	DeleteUser(id int64) error
    CreateUser(user models.UserCreateSchema) (*models.UserResponse, error)
    UpdateUser(id int64, user models.UserUpdateSchema) (*models.UserResponse, error)
    PartialUpdateUser(id int64, user models.UserUpdateSchema) (*models.UserResponse, error)
}

type UserServiceApi struct {
	repository UserRepository
}

func NewUserHandlers(repository UserRepository) *UserServiceApi {
	return &UserServiceApi{repository: repository}
}

// @Tags users
// @Summary Get user by ID
// @Description Get user by ID
// @Tags users
// @Accept  json
// @Produce  json
// @Param   id     path    string     true        "User ID"
// @Success 200 {object} models.UserResponse
// @Failure 400 {object} models.NotFountResponse
// @Router /users/{id} [get]
func (usa *UserServiceApi) ReadUser(c *gin.Context) {
    // Obtener el ID del usuario desde los parámetros de la URL
    userID := c.Param("id")

	 // Convertir userID de string a int64
	 id, err := strconv.ParseInt(userID, 10, 64)
	 if err != nil {
        log.Println("Error while parsing id:", err)
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
        return
    }

	model, err := usa.repository.ReadUser(id)
	if err != nil {
		log.Println("Error while reading id:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}



    // Devolver el usuario en la respuesta
    c.JSON(http.StatusOK, model)
}

// @Tags users
// @Summary Delete user by ID
// @Description Delete user by ID
// @Tags users
// @Accept  json
// @Produce  json
// @Param   id     path    string     true        "User ID"
// @Success 200 {object} models.UserResponse
// @Failure 400 {object} models.NotFountResponse
// @Failure 404 {object} models.NotFountResponse
// @Router /users/{id} [delete]
func (usa *UserServiceApi) DeleteUser(c *gin.Context) {
	
    // Obtener el ID del usuario desde los parámetros de la URL
    userID := c.Param("id")
	

    // Convertir userID de string a int64
    id, err := strconv.ParseInt(userID, 10, 64)
    if err != nil {
        log.Println("Error while parsing id:", err)
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
        return
    }

    // Llamar al método DeleteUser del repositorio
    err = usa.repository.DeleteUser(id)
    if err != nil {
        if err == repository.ErrUserNotFound {
            log.Println("User not found:", err)
            c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
            return
        }
        log.Println("Error while deleting user:", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
        return
    }

    // Devolver un mensaje de éxito en la respuesta
    c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}


// @Tags users
// @Summary Create a new user
// @Description Create a new user
// @Tags users
// @Accept  json
// @Produce  json
// @Param   user  body    models.UserCreateSchema  true  "User Data"
// @Success 201 {object} models.UserResponse
// @Failure 400 {object} models.NotFountResponse
// @Failure 500 {object} models.NotFountResponse
// @Router /users [post]
func (usa *UserServiceApi) CreateUser(c *gin.Context) {
    var newUser models.UserCreateSchema

    // Vincular el JSON recibido al modelo UserCreateSchema
    if err := c.ShouldBindJSON(&newUser); err != nil {
        log.Println("Error while binding JSON:", err)
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
        return
    }

    // Llamar al método CreateUser del repositorio
    createdUser, err := usa.repository.CreateUser(newUser)
    if err != nil {
        log.Println("Error while creating user:", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
        return
    }

    // Devolver el usuario creado en la respuesta
    c.JSON(http.StatusCreated, createdUser)
}

// @Tags users
// @Summary Update an existing user
// @Description Update an existing user
// @Tags users
// @Accept  json
// @Produce  json
// @Param   id    path    int64                   true  "User ID"
// @Param   user  body    models.UserUpdateSchema true  "User Data"
// @Success 200 {object} models.UserResponse
// @Failure 400 {object} models.NotFountResponse
// @Failure 404 {object} models.NotFountResponse
// @Failure 500 {object} models.NotFountResponse
// @Router /users/{id} [put]
func (usa *UserServiceApi) UpdateUser(c *gin.Context) {
    var updateUser models.UserUpdateSchema

    // Obtener el ID del usuario desde los parámetros de la URL
    id := c.Param("id")
    userID, err := strconv.ParseInt(id, 10, 64)
    if err != nil {
        log.Println("Invalid user ID:", err)
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
        return
    }

    // Vincular el JSON recibido al modelo UserUpdateSchema
    if err := c.ShouldBindJSON(&updateUser); err != nil {
        log.Println("Error while binding JSON:", err)
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
        return
    }

    // Llamar al método UpdateUser del repositorio
    updatedUser, err := usa.repository.UpdateUser(userID, updateUser)
    if err != nil {
        if errors.Is(err, repository.ErrUserNotFound) {
            c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        } else {
            log.Println("Error while updating user:", err)
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
        }
        return
    }

    // Devolver el usuario actualizado en la respuesta
    c.JSON(http.StatusOK, updatedUser)
}

// @Tags users
// @Summary Partially update an existing user
// @Description Partially update an existing user
// @Tags users
// @Accept  json
// @Produce  json
// @Param   id    path    int64                   true  "User ID"
// @Param   user  body    models.UserUpdateSchema true  "User Data"
// @Success 200 {object} models.UserResponse
// @Failure 400 {object} models.NotFountResponse
// @Failure 404 {object} models.NotFountResponse
// @Failure 500 {object} models.NotFountResponse
// @Router /users/{id} [patch]
func (usa *UserServiceApi) PartialUpdateUser(c *gin.Context) {
    var updateUser models.UserUpdateSchema

    // Obtener el ID del usuario desde los parámetros de la URL
    id := c.Param("id")
    userID, err := strconv.ParseInt(id, 10, 64)
    if err != nil {
        log.Println("Invalid user ID:", err)
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
        return
    }

    // Vincular el JSON recibido al modelo UserUpdateSchema
    if err := c.ShouldBindJSON(&updateUser); err != nil {
        log.Println("Error while binding JSON:", err)
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
        return
    }

    // Llamar al método PartialUpdateUser del repositorio
    updatedUser, err := usa.repository.PartialUpdateUser(userID, updateUser)
    if err != nil {
        if errors.Is(err, repository.ErrUserNotFound) {
            c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        } else {
            log.Println("Error while updating user:", err)
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
        }
        return
    }

    // Devolver el usuario actualizado en la respuesta
    c.JSON(http.StatusOK, updatedUser)
}
