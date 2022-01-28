package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func TestId(c *gin.Context) {
	request := c.GetUint("currentUser")
	fmt.Println(request)
}
