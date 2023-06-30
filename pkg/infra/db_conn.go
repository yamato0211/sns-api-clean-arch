package infra

import (
	"database/sql"
	"fmt"

	"clean-sns-api/pkg/config"
)

const driverName = "postgres"

type PostgresConnector struct {
	Conn *sql.DB
}

func NewPostgresConnector() *PostgresConnector {
	conf := config.LoadConfig()
	dsn := postgresConnInfo(*conf.PostgresInfo)
	conn, err := sql.Open(driverName, dsn)
	if err != nil {
		panic(err)
	}

	return &PostgresConnector{
		Conn: conn,
	}
}

func postgresConnInfo(postgresInfo config.PostgresInfo) string {
	dataSourceName := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		postgresInfo.PostgresHost,
		postgresInfo.PostgresPort,
		postgresInfo.PostgresUser,
		postgresInfo.PostgresPassword,
		postgresInfo.PostgresDB,
	)

	return dataSourceName
}
