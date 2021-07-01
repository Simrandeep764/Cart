package Routes

import (
	"Cart/Controllers"

	"github.com/gin-gonic/gin"
)

//SetupRouter ... Configure routes
func SetupRouter() {

	r2 := gin.Default()
	grp3 := r2.Group("v1/")
	{
		grp3.GET("cart/:custid", Controllers.GetCartDetailsById)
		grp3.POST("cart", Controllers.AddToCart)
		grp3.PUT("cart/:cid/:prodid", Controllers.UpdateItemsInCart)
	}

	r2.Run(":7002")

}
