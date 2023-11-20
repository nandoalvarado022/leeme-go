package storage

import (
	"database/sql"
	"fmt"
	"log"
)

func NewMysqlClient(DSN string) (client *sql.DB, disconnectFunc func(), err error) {
	client, err = sql.Open("mysql", DSN)
	if err != nil {
		log.Println("error connecting mysql client", "error", err.Error())
		return nil, nil, fmt.Errorf("error connecting mysql client: %w", err)
	}

	disconnectFunc = func() {
		err := client.Close()
		if err != nil {
			log.Println("error disconnecting mongo client", "error", err.Error())
		}
		log.Println("mongo client has been successfully disconnected")
	}

	// defer client.Close()

	err = client.Ping() // Checamos si es exitosa la conexión. Esto devuelve un puntero a un tipo de dato error o en su defecto nil.
	if err != nil {
		panic(err.Error())
	}

	return client, disconnectFunc, nil
}
