package main

import (
	"sync"

	g "github.com/AllenDang/giu"
)

type Ui struct {
	wnd *g.MasterWindow
	mu  sync.RWMutex
}

func (ui *Ui) hasWindowOpen() bool {
	ui.mu.RLock()
	defer ui.mu.RUnlock()

	return ui.wnd != nil
}

func (ui *Ui) toggleWindow(state *State) {
	ui.mu.Lock()
	defer ui.mu.Unlock()

	if ui.wnd != nil {
		ui.wnd.Close()
		ui.wnd = nil
	} else {
		go ui.renderWindow(state)
	}
}

func (ui *Ui) renderWindow(state *State) {

	loop := func() {
		input := g.Label(state.getFilter())
		selectables := []g.Widget{
			input,
		}
		for idx, opt := range state.getFiltered() {
			sel := g.Selectable(opt)
			if idx == state.currentOpt {
				sel.Selected(true)
			}
			selectables = append(selectables, g.Row(sel))
		}

		g.SingleWindow().Layout(
			selectables...,
		)
	}

	flags := g.MasterWindowFlagsFloating | g.MasterWindowFlagsFrameless | g.MasterWindowFlagsNotResizable
	ui.wnd = g.NewMasterWindow("Hello world", 400, 200, flags)
	ui.wnd.Run(loop)
}

func NewUi() *Ui {
	ui := Ui{}
	return &ui
}
