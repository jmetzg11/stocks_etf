package main

import (
	"fmt"
	"stocks_etf/backend/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	fmt.Println("server started")

	r := gin.Default()
	routes.SetupStaticRoutes(r)
	routes.SetupAPIRoutes(r)

	r.Run(":3000")
}
