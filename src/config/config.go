package config

import (
	"database/sql/driver"
	"fmt"
	"log"
	"os"

	"github.com/uptrace/bun/driver/pgdriver"
)

type Conf struct {
	AppPort string
	DbArgs DbArgs
}

type DbArgs struct {
	Insecure bool
	Host string
	User string
	Password string
	Name string
	Port string
	MaxOpenConn int
	MinIdleConn int
}

var conf *Conf

func InitEnv() *Conf {
	conf = new(Conf)
	conf.DbArgs.Insecure = GetEnvBool("DB_INSECURE")
	conf.DbArgs.User = GetEnvOrPanic("DB_USER")
	conf.DbArgs.Password = GetEnvOrDefault("DB_PASSWORD", "123456789")
	conf.DbArgs.Name = GetEnvOrPanic("DB_NAME")
	conf.DbArgs.Host = GetEnvOrPanic("DB_HOST")
	conf.DbArgs.Port = GetEnvOrDefault("DB_PORT", "5432")

	return conf
}

func GetEnvBool(key string) bool {
	strval := os.Getenv(key)
	var b bool
	switch strval {
	case "TRUE", "true", "True", "T", "t":
		b = true
	case "FALSE", "false", "False", "F", "f":
		b = false
	default:
		log.Panicf("Invalid value for boolean environment variable %v=%v", key, strval)
	}

	return b
}

func GetEnvOrDefault(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func GetEnvOrPanic(key string) string {
	value, ok := os.LookupEnv(key)
	if !ok || len(value) == 0 {
		log.Panic("Environment variable not fount or is empty empty string: ", key)
	}
	return value
}

// GetDNS construye y devuelve una cadena de conexión (DNS) para una base de datos PostgreSQL
func (args DbArgs) GetDNS() string {
    // Extrae los valores de los campos de la estructura DbArgs
    host := args.Host
    user := args.User
    password := args.Password
    dbname := args.Name
    port := args.Port

    // Construye la cadena de conexión utilizando los valores extraídos
    dns := "host=" + host + " user=" + user + " password=" + password + " dbname=" + dbname + " port=" + port + " sslmode=disable"
    
    // Devuelve la cadena de conexión construida
    return dns
}

// GetConnector construye y devuelve un conector para la base de datos PostgreSQL
func (args DbArgs) GetConnector() driver.Connector {
    // Construye la dirección del servidor de la base de datos en el formato "host:port"
    addr := fmt.Sprintf("%s:%s", args.Host, args.Port)

    // Crea un nuevo conector de PostgreSQL utilizando los parámetros proporcionados
    pgconn := pgdriver.NewConnector(
        pgdriver.WithAddr(addr),          // Establece la dirección del servidor
        pgdriver.WithUser(args.User),     // Establece el nombre de usuario
        pgdriver.WithPassword(args.Password), // Establece la contraseña
        pgdriver.WithDatabase(args.Name), // Establece el nombre de la base de datos
        pgdriver.WithInsecure(args.Insecure), // Establece si la conexión es insegura (sin SSL)
    )

    // Devuelve el conector de PostgreSQL configurado
    return pgconn
}

