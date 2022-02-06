package main

import (
	"fmt"

	"github.com/go-vgo/robotgo"
)

type kbdDelay int

const (
	DelayInitial kbdDelay = 25
	DelayTyping  kbdDelay = 10
)

func systemInteractOut(out string) {
	if len(out) == 0 {
		return
	}
	robotgo.TypeStr(fmt.Sprintf(" [%s]", out), float64(DelayInitial), float64(DelayTyping))
}
