package main

func main() {
	state := NewState(
		// TODO: scroll lists
		map[string]string{
			"key that is being shown": "expanded value",
			"key2":                    "another expanded value",
			"something":               "nothing",
			"else entirely":           "whatever",
			"more":                    "whatever",
			"stuff":                   "whatever",
			"here":                    "whatever",
			"to":                      "whatever",
			"see":                     "whatever",
			"if it will":              "whatever",
			"scroll":                  "whatever",
		}).
		with(CurrentTimeList()).
		with(PastHourTimeList()).
		with(NextHourTimeList())
	// TODO: templates and more lists
	ui := NewUi()

	update(state, ui)
}
