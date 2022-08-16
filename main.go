package main

import (
	"ordersystem/dao"
	"ordersystem/router"
)

func main() {
	connectSqlx()
	// dao.QueryRow()
	defer dao.CloseSqlxDatabases()

	router := router.SetupRouter()
	// Run("里面不指定端口号默认为8080")
	router.Run(":8000")
}

func connectSqlx() {
	dao.ConnectSqlxDatabases()
}
