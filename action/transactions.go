package actions

import (
	"encoding/json"
	"fmt"
	"instasafe/helper"
	"instasafe/logs"
	"instasafe/models"
	"instasafe/mp"
	tm "instasafe/time"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Transactions(context *gin.Context) {
	var res models.Response
	var req models.Request
	//user, password, hasAuth := context.Request.BasicAuth()
	//flag := auth.Validate(user, password, hasAuth)
	uuid := context.GetString("transId")
	ReqJsonBytes, err := context.GetRawData()
	if err != nil {
		res = helper.Err("Error while getting request body"+err.Error(), uuid)
		context.JSON(http.StatusBadRequest, res)
		return
	}
	isValid, err := helper.ValidateRequest(ReqJsonBytes)
	if !isValid {
		logs.Error("", "invalid json"+err.Error(), uuid)
		res = helper.Err("invalid json"+err.Error(), uuid)
		context.JSON(http.StatusBadRequest, res)
		return
	}
	err = json.Unmarshal(ReqJsonBytes, &req)
	if err != nil { //Error in unmarshal into struct
		res = helper.Err("msg unmarshal failed :"+err.Error(), uuid)
		context.JSON(http.StatusBadRequest, res)
		return
	} else { // Bind Successfully
		logs.Info("request", fmt.Sprintf("Request : %+v", req), uuid)
		logs.Debug("time", "input time : "+req.Timestamp.Format(tm.DateLayout)+" current time : "+tm.TmNow().Format(tm.DateLayout), uuid)
		corrreq, err, status := helper.ActualReq(req)
		if err != nil {
			res = helper.Err(" "+err.Error(), uuid)
			context.JSON(status, res)
			return
		}
		x := mp.Insert(uuid, corrreq)
		context.JSON(x, res)
	}

}
