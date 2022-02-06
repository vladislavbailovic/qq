package main

import (
	hook "github.com/robotn/gohook"
)

const (
	KeyEnter     uint16 = 65293
	KeyCtrl      uint16 = 65507
	KeyTab       uint16 = 65289
	KeyBackspace uint16 = 65288
	KeyUp        uint16 = 65362
	KeyDown      uint16 = 65364
)

func update(state *State, ui *Ui) {
	s := hook.Start()
	isCtrl := false
	for {
		i := <-s

		if i.Rawcode == KeyCtrl {
			isCtrl = i.Kind == hook.KeyDown
		}

		if ui.hasWindowOpen() {
			if i.Rawcode == KeyEnter || i.Rawcode == KeyTab {
				ui.toggleWindow(state)
				out := state.opts[state.currentOpt]
				if state.currentOpt > 0 && len(out) > 0 {
					ui.systemInteractOut(out)
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
					ui.toggleWindow(state)
				}
			}
		}
		// log.Printf("evt: %v\n", i)
	}
}
