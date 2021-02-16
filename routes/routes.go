package routes
import (
	"github.com/VarthanV/inventoryapi/controllers"
	"github.com/gin-gonic/gin"

)

func SetupRouter()* gin.Engine{
r:=gin.Default()
v1 := r.Group("/v1")
{
	v1.GET("/product",controllers.GetProducts)
	v1.GET("/product/:id",controllers.GetProductById)
	v1.POST("/product",controllers.AddProduct)
	v1.PUT("/product/:id",controllers.UpdateProduct)
	v1.DELETE("/product/:id",controllers.DeleteProduct)
}
return r;
}