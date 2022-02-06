package main

import (
	"strings"
	"sync"
)

type State struct {
	raw        map[string]string
	opts       []string
	currentOpt int
	filter     string
	mu         sync.RWMutex
}

func NewState(raw map[string]string) *State {
	state := State{raw: map[string]string{}, opts: []string{}}
	state.with(raw)
	return &state
}

func (s *State) with(raw map[string]string) *State {
	s.mu.Lock()
	defer s.mu.Unlock()

	for key, value := range raw {
		if _, found := s.raw[key]; found {
			continue
		}
		s.raw[key] = value
		s.opts = append(s.opts, key)
	}
	return s
}

func (s *State) getFilter() string {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.filter
}

func (s *State) getFiltered() []string {
	s.mu.RLock()
	defer s.mu.RUnlock()

	out := []string{}
	for _, item := range s.opts {
		if strings.Contains(item, s.filter) {
			out = append(out, item)
		}
	}
	return out
}

func (s *State) updateFilter(flt string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.filter = flt
}

func (s *State) pushToFilter(flt string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.filter += flt
}

func (s *State) popFromFilter() {
	s.mu.Lock()
	defer s.mu.Unlock()

	if len(s.filter) > 0 {
		s.filter = s.filter[:len(s.filter)-1]
	}
}

func (s *State) selectPrevious() {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.currentOpt > 0 {
		s.currentOpt -= 1
	} else {
		s.currentOpt = 0
	}
}

func (s *State) selectNext() {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.currentOpt < len(s.opts)-1 {
		s.currentOpt += 1
	} else {
		s.currentOpt = len(s.opts) - 1
	}
}

func (s *State) getCurrentOpt() int {
	s.mu.RLock()
	defer s.mu.RUnlock()

	max := len(s.getFiltered()) - 1
	if s.currentOpt > max {
		return max
	}
	return s.currentOpt
}

func (s *State) getSelected() string {
	s.mu.RLock()
	defer s.mu.RUnlock()

	flt := s.getFiltered()
	key := flt[s.getCurrentOpt()]
	out := s.raw[key]
	return out
}
