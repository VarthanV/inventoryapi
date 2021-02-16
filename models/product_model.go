package models
import(
	"github.com/VarthanV/inventoryapi/config"
	_ "github.com/go-sql-driver/mysql"
)
type Product struct{
	Id    	uint    `json:"id" gorm:"primaryKey"` 
	Name  	string  `json:"name"`
	Price 	float64 `json:"price"`
	Stock 	int64   `json:"stock"` 
	InStock bool	`json:"in_stock"`
}

func (t*Product) TableName() string{
	return "product"
}

func GetAllProducts(product *[]Product)(err error){
	if err = config.DB.Find(product).Error; err != nil {
		return err
	   }
	   return nil
}

func GetProductById(product *Product,id string) (err error){
	if err = config.DB.Where("id= ?",id).First(product).Error; err !=nil{
		return err
	}
	return nil
}

func AddProduct(product *Product) (err error){
	if err := config.DB.Create(product).Error ; err !=nil{
		return err
	}
	return nil
}

func UpdateProduct(product *Product,id string) (err error){
	if err := config.DB.Update(product).Error  ; err !=nil{
		return err
	}
	return nil
}

func DeleteProduct(product *Product ,id string)(err error){
	if err := config.DB.Where("id= ?",id).Delete(product).Error; err !=nil{
		return err
	}
	return nil
}