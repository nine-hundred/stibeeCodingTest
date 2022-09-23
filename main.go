package main

import (
	"stibee/controllers"
)

func main() {
	r := controllers.InitRouter()
	r.Run(":8090")
}
