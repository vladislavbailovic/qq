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

		if i.Kind == hook.KeyDown && ui.hasWindowOpen() {
			if i.Rawcode == KeyEnter || i.Rawcode == KeyTab {
				ui.toggleWindow(state)
				ui.systemInteractOut(state.getSelected())
			}
			if i.Rawcode == KeyUp {
				state.selectPrevious()
			}
			if i.Rawcode == KeyDown {
				state.selectNext()
			}
		}
		// log.Printf("evt: %v\n", i)
	}
}
