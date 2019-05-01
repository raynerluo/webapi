package routes

import (
	"webapi/controllers"

	"github.com/kataras/iris"
)

//注册路由
func RegisterRoutes(routes iris.Party) {
	routes.PartyFunc("/custom", controllers.Custom)
}
