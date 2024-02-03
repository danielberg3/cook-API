package main

import (
	"Cook_API/APP/API/config"
	"Cook_API/APP/API/endpoints/router"
	"Cook_API/infra/postgres"
	"fmt"
)

func main() {
	SetUpPostgres()
	startServer(config.ServerHost, config.ServerPort)
}

func SetUpPostgres() {
	err := postgres.SetupCredentials(
		config.PostgresUser,
		config.PostgresPassword,
		config.PostgresDBName,
		config.ServerHost,
		config.PostgresPort,
	)
	if err != nil {
		panic(err)
	}
}

func startServer(host string, port int) {
	server := router.Start()
	server.Debug = true
	adress := fmt.Sprintf("%s:%d", host, port)
	server.Logger.Info(server.Start(adress))
}
