package actions

import (
	"instasafe/models"
	"instasafe/mp"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Delete(c *gin.Context) {
	//mp.M = make(map[string]models.CorrReq)
	for k := range mp.M {
		delete(mp.M, k)
	}

	c.JSON(http.StatusOK, models.Response{})
}
