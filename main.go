package main

import (
	"github.com/Emon46/vector-config-server/api"
	"log"
)

func main() {
	config, err := api.LoadConfig("/")
	if err != nil {
		log.Fatal(err)
	}

	server, err := api.NewServer(config)
	if err != nil {
		log.Fatal("cannot create server: ", err)
	}
	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}

}
