package main

import "ordersystem/router"

func main() {
	router := router.SetupRouter()
	// Run("里面不指定端口号默认为8080")
	// router.Run(":8000")
	router.Run(":8000")
}
