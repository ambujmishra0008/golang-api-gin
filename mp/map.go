package mp

import (
	"instasafe/models"
	tm "instasafe/time"
	"time"
)

var M = make(map[string]models.CorrReq)

func Insert(key string, value models.CorrReq) int {
	now := tm.TmNow().Add(-time.Minute * time.Duration(tm.TIME_LAG))
	if value.Timestamp.After(now) {
		M[key] = value
		return 200
	} else {
		return 204
	}
}
