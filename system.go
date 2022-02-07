package main

import (
	"fmt"
	"log"

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
