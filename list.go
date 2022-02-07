package main

import (
	"encoding/base64"
	"fmt"
	"strconv"
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

func ClipboardTimeList() map[string]string {
	result := map[string]string{}
	what := systemGetClipboard()
	if len(what) == 0 {
		return result
	}

	tm, err := strconv.Atoi(what)
	if err != nil {
		return result
	}

	if tm/int(time.Millisecond) > 10000 {
		tm = tm / (1000 * int(time.Millisecond))
	}

	return timeList("clipboard", time.Unix(int64(tm), 0))
}

func ClipboardBase64List() map[string]string {
	result := map[string]string{}
	what := systemGetClipboard()
	if len(what) == 0 {
		return result
	}

	if val := base64.StdEncoding.EncodeToString([]byte(what)); len(val) > 0 {
		result["clipboard base64 encode"] = val
	}

	if val, err := base64.StdEncoding.DecodeString(what); err == nil {
		result["clipboard base64 decode"] = string(val)
	}
	return result
}
