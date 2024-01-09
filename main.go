package main

import (
	"connect-mysql/db"
	routers "connect-mysql/routes"
	"fmt"

	"github.com/gin-gonic/gin"
)

func init() {
	db.InitialMigration()
}

func main() {
	fmt.Println("Hello auth")
	r := gin.Default()
	routers.GetRoute(r)

	r.Run()
}
