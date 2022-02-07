package main

func reloadState(state *State) *State {
	// TODO: templates and more lists
	state.clear().
		with(CurrentTimeList()).
		with(PastHourTimeList()).
		with(NextHourTimeList()).
		with(ClipboardTimeList()).
		with(ClipboardBase64List()).
		with(StaticList())
	return state
}

func main() {
	state := reloadState(
		NewState(map[string]string{}))
	ui := NewUi()

	update(state, ui)
}
