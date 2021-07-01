package Models

import (
	"Cart/Config"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func AddToCart(cart *Cart) (err error) {
	if err = Config.DB.Create(cart).Error; err != nil {
		return err
	}
	return nil
}

func GetCartDetailsById(cart *[]Cart, id string) (err error) {
	if err = Config.DB.Where("customer_id = ?", id).Find(cart).Error; err != nil {
		return err
	}
	return nil
}

func UpdateItemsInCart(cart *Cart, id string, id1 string) (err error) {
	fmt.Println(cart)
	Config.DB.Save(cart)
	return nil
}
