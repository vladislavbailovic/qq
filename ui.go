package main

import (
	"log"

	g "github.com/AllenDang/giu"
)

func renderWindow(state *State) {

	loop := func() {
		selectables := []g.Widget{}
		for idx, opt := range state.opts {
			sel := g.Selectable(opt)
			if idx == state.currentOpt {
				sel.Selected(true)
			}
			selectables = append(selectables, g.Row(sel))
		}

		g.SingleWindow().Layout(
			selectables...,
		)
		log.Printf("still being updated: %v", state.wnd)
	}

	flags := g.MasterWindowFlagsFloating | g.MasterWindowFlagsFrameless | g.MasterWindowFlagsNotResizable
	state.wnd = g.NewMasterWindow("Hello world", 400, 200, flags)
	state.wnd.Run(loop)
}
