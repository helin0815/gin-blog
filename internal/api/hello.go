package api

import "github.com/gin-gonic/gin"

func (as *Apis) HelloWorld(c *gin.Context) {
	c.JSON(200, "hello world")
}
