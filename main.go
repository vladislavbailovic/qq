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

	update(&state, ui)
}
