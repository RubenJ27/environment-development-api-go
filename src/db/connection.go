package db

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"errors"
	"time"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
)

type Args interface {
	GetDNS() string
	GetConnector() driver.Connector
}

// GetBunConnection establece una conexión con la base de datos utilizando la configuración proporcionada.
// Devuelve una instancia de bun.DB y un error en caso de que ocurra algún problema.
func GetBunConnection(dbconf Args) (*bun.DB, error) {
    // Abre una conexión a la base de datos utilizando el conector proporcionado en dbconf.
    sqldb := sql.OpenDB(dbconf.GetConnector())
    
    // Crea una nueva instancia de bun.DB utilizando la conexión SQL y el dialecto de PostgreSQL.
    dbBun := bun.NewDB(sqldb, pgdialect.New())

    // Crea un contexto con un tiempo de espera de 5 segundos para la operación de ping.
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    // Asegura que el contexto se cancele después de que la función termine.
    defer cancel()

    // Verifica la conexión a la base de datos utilizando el contexto con tiempo de espera.
    err := dbBun.PingContext(ctx)
    if err != nil {
        // Si ocurre un error al conectar, devuelve nil y un error con un mensaje descriptivo.
        return nil, errors.New("cannot connect to database: " + err.Error())
    }

    // Si la conexión es exitosa, devuelve la instancia de bun.DB y nil como error.
    return dbBun, nil
}
