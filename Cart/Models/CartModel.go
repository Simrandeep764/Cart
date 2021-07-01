package Models

type Cart struct {
	CartId     uint `gorm:"primary_key;auto_increment:false" json:"cid"`
	ProductId  uint `gorm:"primary_key;auto_increment:false" json:"prodid"`
	CustomerId uint `json:"custid"`
	Qty        int  `json:"prodqty" gorm:"type:int"`
}

func (c *Cart) TableName() string {
	return "cart"
}
