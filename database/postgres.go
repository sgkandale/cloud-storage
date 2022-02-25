package database

import (
	"context"
	"fmt"
	"log"

	"cloud_storage/config"

	"github.com/jackc/pgx/v4/pgxpool"
)

const FilesTable = "items"
const UsersTable = "users"

func NewPostgresConn() *pgxpool.Pool {

	log.Println("connecting Postgres")

	connStr := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s",
		config.Config.Postgres.User, config.Config.Postgres.Password,
		config.Config.Postgres.Host, config.Config.Postgres.Port,
		config.Config.Postgres.DBName,
	)

	conn, err := pgxpool.Connect(
		context.TODO(),
		connStr,
	)
	if err != nil {
		log.Fatal("error opening Postgres : ", err)
	}

	pingErr := conn.Ping(
		context.TODO(),
	)
	if pingErr != nil {
		log.Fatal("error pinging Postgres : ", pingErr)
	}

	log.Println("Postgres Connected")

	return conn
}

var DB = NewPostgresConn()
