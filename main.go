package main

func main() {
	state := NewState(
		map[string]string{
			"key that is being shown": "expanded value",
			"key2":                    "another expanded value",
			"something":               "nothing",
			"else entirely":           "whatever",
		}).
		with(TimeList())
	ui := NewUi()

	update(state, ui)
}
