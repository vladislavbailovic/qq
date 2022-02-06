package main

import (
	"sync"
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

	// hook.Register(hook.KeyDown, []string{"ctrl", "q"}, func(e hook.Event) {
	// 	ui.toggleWindow(&state)
	// })

	// hook.Register(hook.KeyDown, []string{"up"}, func(e hook.Event) {
	// 	if state.currentOpt > 0 {
	// 		state.currentOpt -= 1
	// 	}
	// })

	// hook.Register(hook.KeyDown, []string{"down"}, func(e hook.Event) {
	// 	if state.currentOpt < len(state.opts)-1 {
	// 		state.currentOpt += 1
	// 	}
	// })

	// s := hook.Start()
	// <-hook.Process(s)

	update(&state, ui)
}
