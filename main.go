package main

import (
	"fmt"

	"github.com/joaoino100/go-rest-api/database"
	"github.com/joaoino100/go-rest-api/models"
	"github.com/joaoino100/go-rest-api/routes"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "Nome 1", Historia: "Historia 1"},
		{Id: 2, Nome: "Nome 2", Historia: "Historia 2"},
	}

	database.Conectar()
	fmt.Println("Iniciando o servidor Rest")
	routes.HandleRequest()
}
