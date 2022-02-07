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
	DelayInitial kbdDelay = 25
	DelayTyping  kbdDelay = 10
)

func systemInteractOut(raw string) {
	if len(raw) == 0 {
		return
	}
	lines := strings.Split(raw, "\n")

	for idx, out := range lines {
		if idx == 0 {
			// Idk why we strip first 1? chars on first line
			out = fmt.Sprintf(" %s", out)
		}
		robotgo.TypeStr(out, float64(DelayInitial), float64(DelayTyping))
		if len(lines) > 1 {
			robotgo.KeyTap("\n")
			time.Sleep(50 * time.Millisecond)
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
