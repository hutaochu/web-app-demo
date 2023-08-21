package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"k8s.io/klog/v2"
)

type HelloRequest struct {
	Name string `form:"name" binding:"required"`
}

// Hello godoc
// @Summary      Get hello info
// @Description  get hello info by name
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        name   query   string  true  "username"
// @Success      200  {object}  string
// @Router       /user/hello [get]
func Hello(c *gin.Context) {
	var req HelloRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, "request parameters missing")
		return
	}
	if req.Name == "test" {
		r1 := []int{1, 2, 3, 0}
		for _, v := range r1 {
			klog.Info(10 / v)
		}
	}
	c.JSON(http.StatusOK, "hello")
}
