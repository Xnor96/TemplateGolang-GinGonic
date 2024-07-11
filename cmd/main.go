package main

import (
	"fmt"
	"templategolang-gingonic/internal/config"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {

	//* OBTENGO LAS LAS VARIABLES DE ENTORNO
	cnf := config.NewConfig()
	fmt.Println(cnf)

	// Inicializar el router de Gin
	routes := gin.Default()

	routes.GET("/", nil)

}
