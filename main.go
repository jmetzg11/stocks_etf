package main

import (
	"fmt"
	"stocks_etf/backend/routes"

	"stocks_etf/backend/database"

	"github.com/gin-gonic/gin"
)

func main() {

	fmt.Println("server started")
	database.Connect()

	r := gin.Default()
	routes.SetupStaticRoutes(r)
	routes.SetupAPIRoutes(r)

	r.Run(":3000")
}
