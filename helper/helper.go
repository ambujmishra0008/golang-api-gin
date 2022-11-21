package helper

import (
	"errors"
	"instasafe/logs"
	"instasafe/models"
	tm "instasafe/time"
	"net/http"
	"strconv"
)

func Err(msg, uuid string) models.Response {
	logs.Error("error", msg, uuid)
	res := models.Response{
		TransId: uuid,
		Status:  http.StatusBadRequest,
		Message: msg,
	}
	return res
}

func ActualReq(req models.Request) (models.CorrReq, error, int) {
	val, err := StringToFloat(req.Amount)
	if err != nil {
		return models.CorrReq{}, err, http.StatusBadRequest
	}
	if req.Timestamp.After(tm.TmNow()) {
		return models.CorrReq{}, errors.New("timestamp is in future"), 422
	}
	return models.CorrReq{
		Amount:    val,
		Timestamp: req.Timestamp,
	}, nil, 200
}

func StringToFloat(str string) (float64, error) {
	if s, err := strconv.ParseFloat(str, 64); err == nil {
		return s, nil
	} else {
		return 0, err
	}
}
