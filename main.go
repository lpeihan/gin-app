package main

import "gin-app/router"

func main() {
	r := router.Router()

	r.Run("0.0.0.0:7002")
}
