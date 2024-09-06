package rest

import (
	"development-environment-api-go-manager/src/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserRepository interface {
	ReadUser(id int64) (*models.UserResponse, error)
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
// @Success 200 {object} models.User
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



