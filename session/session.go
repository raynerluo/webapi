package session

import (
	"github.com/kataras/iris/sessions"
)

var (
	Sess *sessions.Sessions
)

//初始化session
func init() {
	Sess = sessions.New(sessions.Config{Cookie: "haha54xxx"})
}
