package main

import (
	"github.com/claudiootte/restapiexample/controller"
	"github.com/claudiootte/restapiexample/model"
)

func main() {
	model.InitialMigration()
	controller.InitializeRouter()
}
