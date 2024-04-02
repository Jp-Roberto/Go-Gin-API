package main

import (
	"github.com/Jp-Roberto/api-go-gin/database"
	"github.com/Jp-Roberto/api-go-gin/routes"
)

func main() {
	database.ConectaComBancoDeDados()

	routes.HandleRquest()
}
