package date_utils

import (
	"time"
)

const (
	APIDATELAYOUT = "2006-01-02T15:04:05Z"
	APIDBFORMAT   = "2006-01-02 15:04:05"
)

func GetNow() time.Time {
	return time.Now()
}

func GetNowString() string {
	return GetNow().Format(APIDATELAYOUT)
}

func GetNowDBFormat() string {
	return GetNow().Format(APIDBFORMAT)
}
