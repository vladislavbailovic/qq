package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/go-vgo/robotgo"
)

type kbdDelay int

const (
	DelayTyping  kbdDelay = 5
	DelayInitial kbdDelay = 50 // Not sure what this does
	DelayTap     kbdDelay = 10
)

func systemInteractOut(raw string) {
	if len(raw) == 0 {
		return
	}
	lines := strings.Split(raw, "\n")

	for idx, out := range lines {
		if idx == 0 {
			// Idk why we strip first 1? chars on first line
			// Issue: https://github.com/go-vgo/robotgo/issues/315
			out = fmt.Sprintf("  %s", out)
		}
		robotgo.TypeStr(out, float64(DelayTyping), float64(DelayInitial))
		if len(lines) > 1 {
			robotgo.KeyTap("\n")
			time.Sleep(time.Duration(DelayTap) * time.Millisecond)
		}
	}
}

func systemGetClipboard() string {
	str, err := robotgo.ReadAll()
	if err != nil {
		log.Printf("unable to read from clipboard: %v\n", err)
	}
	return str
}

func systemSetClipboard(what string) {
	err := robotgo.WriteAll(what)
	if err != nil {
		log.Printf("unable to write to clipboard: %v\n", err)
	}
}
