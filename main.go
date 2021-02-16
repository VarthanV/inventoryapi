package main

import (
	"github.com/VarthanV/inventoryapi/config"
	"github.com/VarthanV/inventoryapi/models"
	"github.com/VarthanV/inventoryapi/routes"
	"github.com/jinzhu/gorm"
)
var err error
func main(){
	config.DB ,err =gorm.Open("mysql",config.DbURL(config.BuildDBConfig()))
	if err != nil{
		panic(err)
	}
	defer config.DB.Close()
	config.DB.AutoMigrate(&models.Product{})
	r := routes.SetupRouter()
	r.Run()
}