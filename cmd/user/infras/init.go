package infras

import "github.com/kalandramo/appdemo/cmd/user/infras/mysql"

// Init init dal
func Init() {
	mysql.Init() // mysql init
}
