package actions

import (
	"instasafe/models"
	"instasafe/mp"
	tm "instasafe/time"
	"math"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	avg   float64
	sum   float64
	mn    float64
	mx    float64
	count int
)

func Stat(c *gin.Context) {
	avg = 0
	sum = 0
	mn = 0
	mx = 0
	count = 0
	transId := c.GetString("transId")
	for k, v := range mp.M {
		if v.Timestamp.Before(tm.TmNow().Add(-time.Minute * time.Duration(tm.TIME_LAG))) {
			delete(mp.M, k)
		} else {
			sum += v.Amount
			count++
			mn = math.Min(v.Amount, mn)
			mx = math.Max(v.Amount, mx)
			avg = sum / float64(count)
		}
	}
	res := models.Response{
		Status:  http.StatusOK,
		Message: "success",
		TransId: transId,
		Sum:     sum,
		Avg:     avg,
		Max:     mx,
		Min:     mn,
		Count:   count}

	c.JSON(http.StatusOK, res)
}
