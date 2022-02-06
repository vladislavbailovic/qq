package main

func TimeList() map[string]string {
	return map[string]string{
		"UNIX time":         "time()",
		"timestamp (UTC)":   "Y-m-dTH:i:s00",
		"timestamp (local)": "Y-m-dTH:i:sZ",
	}
}
