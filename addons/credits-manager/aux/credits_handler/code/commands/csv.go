package commands

import (
	local "credits_manager/error"
	"credits_manager/infra"
	"encoding/json"
	"fmt"
)

const fail = "fail"

var csvHeader = []string{
	"work",
	"author",
	"link",
	"licence",
	"description",
}

type csv_query struct {
	Path          string `json:"path"`
	Attribuitions bool   `json:"attribuitions"`
	Licences      bool   `json:"licences"`
}

func ExportCSV(query string) string {

	var queryData csv_query
	if err := json.Unmarshal([]byte(query), &queryData); err != nil {
		local.HandleErrorMessage("Cant unmarshal csv query", err)
		return fail
	}

	list := infra.ListCredits("asc", "")
	filepath := fmt.Sprintf("%s/works.csv", queryData.Path)

	data := [][]string{
		csvHeader,
	}

	for _, item := range list {
		data = append(data, []string{
			item.Name,
			item.Author,
			item.Link,
			item.Licence,
			item.Type,
		})
	}
	if err := infra.SaveCSV(filepath, data); err != nil {
		local.HandleErrorMessage("Cant save csv file", err)
		return fail
	}

	return fmt.Sprintf("saved")
}
