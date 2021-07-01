package Controllers

import (
	"Cart/Models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddToCart(c *gin.Context) {
	var cart Models.Cart
	c.BindJSON(&cart)
	fmt.Println(cart)
	err := Models.AddToCart(&cart)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, cart)
	}
}

func GetCartDetailsById(c *gin.Context) {
	id := c.Params.ByName("custid")
	var cart []Models.Cart
	err := Models.GetCartDetailsById(&cart, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, cart)
	}
}

func UpdateItemsInCart(c *gin.Context) {
	var userCart []Models.Cart
	c.BindJSON(&userCart)
	id := c.Params.ByName("custid")
	//id := userCart.CustomerId
	err := Models.GetCartDetailsById(&userCart, string(id))
	if err != nil {
		c.JSON(http.StatusNotFound, userCart)
	}
	c.BindJSON(&userCart)
	var User1 Models.Cart
	id2 := c.Params.ByName("cid")
	id1 := c.Params.ByName("prodid")
	err = Models.UpdateItemsInCart(&User1, id2, id1)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, User1)
	}
}
