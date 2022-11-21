package time

import (
	"instasafe/env"
	"strconv"
	"time"
)

var TIME_LAG int

func init() {
	var s string = env.Env["TIME_LAG"]
	TIME_LAG, _ = strconv.Atoi(s)
}
func TmNow() time.Time {
	return time.Now().UTC()
}

var DateLayout = "2006-01-02 15:04:05.000000000"
