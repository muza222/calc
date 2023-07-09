package main

import (
	"encoding/json"
	"fmt"
)

func printStats() {
	var writer stats
	switch format {
	case "tabular":
		writer = formatTabular()
	case "csv":
		writer = formatCSV()
	case "json":
		s, _ := json.Marshal(groupFinalUser)
		fmt.Println(string(s))
		return
	case "json-lines":
		writer = jsonWriter{}
	default:
		attention()
	}
	for _, item := range groupFinalUser {
		writer.write(item)
	}
	writer.flush()
}
