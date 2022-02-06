package main

import (
	"fmt"
	"log"
	"sync"

	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
)

type State struct {
	opts       []string
	currentOpt int
	mu         sync.RWMutex
}

func main() {
	state := State{
		opts: []string{
			"",
			":this is my list",
			":of selected strings",
			":whatever the fuck goes here",
			":is what goes here",
			":yay",
		},
	}
	ui := NewUi()

	s := hook.Start()

	isCtrl := false
	for {
		i := <-s

		if i.Rawcode == KeyCtrl {
			isCtrl = i.Kind == hook.KeyDown
		}

		if ui.hasWindowOpen() {
			if i.Rawcode == KeyEnter || i.Rawcode == KeyTab {
				ui.toggleWindow(&state)
				out := state.opts[state.currentOpt]
				if state.currentOpt > 0 && len(out) > 0 {
					robotgo.TypeStr(fmt.Sprintf(" [%s]", out), 25, 10)
					log.Println(out)
				}
			}
		}

		if i.Kind == hook.KeyDown {
			if i.Rawcode == KeyUp {
				if state.currentOpt > 0 {
					state.currentOpt -= 1
				}
			}
			if i.Rawcode == KeyDown {
				if state.currentOpt < len(state.opts)-1 {
					state.currentOpt += 1
				}
			}
		}

		if i.Kind == hook.KeyDown && i.Rawcode < 255 {
			key := string(int32(i.Rawcode))
			if isCtrl {
				if key == "q" {
					ui.toggleWindow(&state)
				}
			}
			// log.Printf("evt: %v\n", string(i.Rawcode))
		}
		// log.Printf("evt: %v\n", i)
	}
}
