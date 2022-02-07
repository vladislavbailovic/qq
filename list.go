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
		tm = tm / 1000
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

func StaticList() map[string]string {
	return map[string]string{
		"static longish text": "So this is my text\nwithsome newlines\n\tand tabs<-",
		"static long text": `Within the format string, the underscores in "_2" and "__2" represent spaces that may be replaced by digits if the following number has multiple digits, for compatibility with fixed-width Unix time formats. A leading zero represents a zero-padded value.

		The formats and 002 are space-padded and zero-padded three-character day of year; there is no unpadded day of year format.

		A comma or decimal point followed by one or more zeros represents a fractional second, printed to the given number of decimal places. A comma or decimal point followed by one or more nines represents a fractional second, printed to the given number of decimal places, with trailing zeros removed. For example "15:04:05,000" or "15:04:05.000" formats or parses with millisecond precision.

		Some valid layouts are invalid time values for time.Parse, due to formats such as _ for space padding and Z for zone information.`,
	}
}
