package main

import (
	"babyblog/model"
	"babyblog/routes"
)

func main() {
	model.InitDb()
	routes.InitRouter()
}
