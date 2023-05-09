package infras

import "github.com/kalandramo/mocheng/cmd/user/infras/mysql"

// Init init dal
func Init() {
	mysql.Init() // mysql init
}
