package controllers

type CartInput struct {
	Price     int `json:"price"`
	Total     int `json:"total"`
	ProductID int `json:"productId"`
}

// func GetAllCartByUser(c *gin.Context){

// }
