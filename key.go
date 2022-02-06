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

type counterState int

const (
	StateReady counterState = iota
	StatePending
	StateDone
)

type ctrlCounter struct {
	hit   int
	state counterState
}

func (cc *ctrlCounter) update(in hook.Event) {
	if in.Kind != hook.KeyDown {
		return
	}
	if in.Rawcode != KeyCtrl {
		cc.reset()
		return
	}

	if !cc.isDone() {
		cc.state = StatePending
		cc.hit += 1
		if cc.hit > 2 {
			cc.state = StateDone
		}
	}
}

func (cc *ctrlCounter) reset() {
	cc.state = StateReady
	cc.hit = 0
}

func (cc ctrlCounter) isDone() bool {
	return cc.state == StateDone
}

func update(state *State, ui *Ui) {
	s := hook.Start()
	cc := ctrlCounter{}
	for {
		i := <-s

		cc.update(i)
		if cc.isDone() {
			ui.toggleWindow(state)
			cc.reset()
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

		// if i.Kind == hook.KeyDown && i.Rawcode < 255 {
		// 	key := string(int32(i.Rawcode))
		// 	if isCtrl {
		// 		if key == "q" {
		// 			ui.toggleWindow(state)
		// 		}
		// 	}
		// }
		// log.Printf("evt: %v\n", i)
	}
}
