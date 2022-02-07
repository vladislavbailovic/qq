package main

import (
	"fmt"
	"testing"
	"time"
)

func TestCurrentTimeList(t *testing.T) {
	testTimeList(t, "now", time.Now(), CurrentTimeList())
}

func TestPastHourTimeList(t *testing.T) {
	testTimeList(t, "hour ago", time.Now().Add(-1*3600*time.Second), PastHourTimeList())
}

func TestNextHourTimeList(t *testing.T) {
	testTimeList(t, "hour from now", time.Now().Add(1*3600*time.Second), NextHourTimeList())
}

func testTimeList(t *testing.T, name string, ts time.Time, actual map[string]string) {
	expectedUnix := fmt.Sprintf("%d", ts.Unix())
	expectedRFC339Local := ts.Local().Format(time.RFC3339)
	expectedRFC339UTC := ts.UTC().Format(time.RFC3339)

	keyUnix := fmt.Sprintf("UNIX time (%s)", name)
	keyRFC339Local := fmt.Sprintf("timestamp (%s, local)", name)
	keyRFC339UTC := fmt.Sprintf("timestamp (%s, UTC)", name)

	if len(actual) != 3 {
		t.Log(actual)
		t.Fatalf("expected 3 list members, got %d", len(actual))
	}

	actualUnix, ok := actual[keyUnix]
	if !ok {
		t.Fatalf("expected key to be present: %s", keyUnix)
	}
	if actualUnix != expectedUnix {
		t.Fatalf("expected [%s] but got [%s] for: [%s]", actualUnix, expectedUnix, keyUnix)
	}

	actualRFC339Local, ok := actual[keyRFC339Local]
	if !ok {
		t.Fatalf("expected key to be present: %s", keyRFC339Local)
	}
	if actualRFC339Local != expectedRFC339Local {
		t.Fatalf("expected [%s] but got [%s] for: [%s]", actualRFC339Local, expectedRFC339Local, keyRFC339Local)
	}

	actualRFC339UTC, ok := actual[keyRFC339UTC]
	if !ok {
		t.Fatalf("expected key to be present: %s", keyRFC339UTC)
	}
	if actualRFC339UTC != expectedRFC339UTC {
		t.Fatalf("expected [%s] but got [%s] for: [%s]", actualRFC339UTC, expectedRFC339UTC, keyRFC339UTC)
	}
}
