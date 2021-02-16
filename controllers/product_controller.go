package controllers

import (
	"fmt"
	"net/http"

	"github.com/VarthanV/inventoryapi/models"
	"github.com/gin-gonic/gin"
)


func GetProducts(c*gin.Context){
	var product []models.Product;
	fmt.Println(product)
	err := models.GetAllProducts(&product)
	if err != nil{
		c.JSON(http.StatusBadRequest,gin.H{"err":err.Error()})
	}else{
		c.JSON(http.StatusOK,product)
	}
}

func AddProduct(c*gin.Context){
	var product models.Product
	if err := c.ShouldBindJSON(&product) ; err !=nil{
		 c.JSON(http.StatusBadRequest,gin.H{"err":err.Error()})
		 return
	}
	err := models.AddProduct(&product)
	if err != nil{
		c.JSON(http.StatusInternalServerError,gin.H{"err":err.Error()})
	}else{
		c.JSON(http.StatusCreated,product)
	}

}

func GetProductById(c*gin.Context){
	var product models.Product
	id := c.Param("id")
	err := models.GetProductById(&product,id)
	if err != nil{
		c.JSON(http.StatusInternalServerError ,gin.H{"err":err.Error()})
	}else{
		c.JSON(http.StatusOK  ,product)
	}
}

func UpdateProduct(c*gin.Context){
	var product models.Product
	id := c.Param("id")
	c.BindJSON(&product)
	err := models.UpdateProduct(&product,id)
	if err != nil{
		c.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
	}else{
		c.JSON(http.StatusAccepted,product)
	}
}

func DeleteProduct(c*gin.Context){
	var product models.Product
	id := c.Param("id")
	err := models.DeleteProduct(&product,id)
	if err != nil{
	   c.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
	}else{
		c.JSON(http.StatusContinue,gin.H{"message":"Succes"})
	}
}