package main

import (
	"fmt"
	"time"
)

func timeList(title string, ts time.Time) map[string]string {
	return map[string]string{
		fmt.Sprintf("UNIX time (%s)", title):        fmt.Sprintf("%d", ts.Unix()),
		fmt.Sprintf("timestamp (%s, UTC)", title):   ts.UTC().Format(time.RFC3339),
		fmt.Sprintf("timestamp (%s, local)", title): ts.Local().Format(time.RFC3339),
	}
}

func CurrentTimeList() map[string]string {
	return timeList("now", time.Now())
}

func PastHourTimeList() map[string]string {
	return timeList("hour ago", time.Now().Add(-1*3600*time.Second))
}

func NextHourTimeList() map[string]string {
	return timeList("hour from now", time.Now().Add(3600*time.Second))
}
