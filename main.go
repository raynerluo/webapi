package main

import (
	_ "webapi/config"
	_ "webapi/database"
	"webapi/routes"
	_ "webapi/session"

	"github.com/kataras/iris"
)

func main() {
	app := iris.Default()
	app.PartyFunc("/webapi", routes.RegisterRoutes)
	app.Run(iris.Addr(":8080"))
}
