package controllers

import (
	"webapi/session"

	"github.com/kataras/iris"
)

func Custom(custom iris.Party) {

	custom.Post("/login", func(ctx iris.Context) {
		sess := session.Sess.Start(ctx)
		result := iris.Map{"status": false}
		username := ctx.PostValue("username")
		password := ctx.PostValue("password")
		if username != "" && password != "" {
			sess.Set("username", username)
			sess.Set("password", password)
			result["status"] = true
		}
		ctx.JSON(result)
	})
}
