package main

import "ordersystem/router"

func main() {
	router := router.SetupRouter()
	router.Run(":8000")
}
