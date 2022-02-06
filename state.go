package main

import "sync"

type State struct {
	opts       []string
	currentOpt int
	mu         sync.RWMutex
}

func (s *State) selectPrevious() {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.currentOpt > 0 {
		s.currentOpt -= 1
	}
}

func (s *State) selectNext() {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.currentOpt < len(s.opts)-1 {
		s.currentOpt += 1
	}
}

func (s *State) getSelected() string {
	s.mu.RLock()
	defer s.mu.RUnlock()

	out := s.opts[s.currentOpt]
	return out
}
