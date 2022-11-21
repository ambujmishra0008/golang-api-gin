package id

import (
	"strconv"
	"time"
)

var RowPtr int64

func init() {
	RowPtr = 0
}

func GenerateRow() string {
	hex_str_unix_time := strconv.FormatInt(time.Now().Unix(), 16)
	str_iterator := strconv.FormatInt(RowPtr%100000, 10)
	final_uuid := "T" + hex_str_unix_time + str_iterator
	RowPtr++
	return final_uuid
}
