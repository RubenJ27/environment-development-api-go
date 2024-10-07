package main

import (
	"log"

	"development-environment-api-go-manager/src/api/rest"
	"development-environment-api-go-manager/src/api/rest/routes"
	"development-environment-api-go-manager/src/config"
	"development-environment-api-go-manager/src/db"
	"development-environment-api-go-manager/src/db/repository"
	_ "development-environment-api-go-manager/src/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
    // Imprime un mensaje en la consola indicando que el servidor está en ejecución
    log.Println("Server Go is running...")

    // Inicializa la configuración del entorno desde variables de entorno o archivos de configuración
    conf := config.InitEnv()

    // Establece una conexión con la base de datos utilizando Bun ORM
    connBun, err := db.GetBunConnection(conf.DbArgs)
    if err != nil {
        // Si hay un error al conectar con la base de datos, detiene la ejecución del programa
        panic(err)
    }

    // Crea una nueva instancia del repositorio de usuarios utilizando la conexión a la base de datos
    userRepo := repository.NewUserRepository(connBun)
    // Crea una instancia de UserServiceApi
    userServiceApi := rest.NewUserHandlers(userRepo)

    // Configura el servidor HTTP con las rutas necesarias y el repositorio de usuarios
    r := routes.GetServer(userRepo)

    // Configura la ruta para la documentación de Swagger
    r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
    // Configura la ruta para eliminar un usuario
    r.DELETE("/users/:id", userServiceApi.DeleteUser)

    // Inicia el servidor HTTP en el puerto 8080
    r.Run()
}