package commands

import (
	local "credits_manager/error"
	"credits_manager/infra"
	"credits_manager/model"
	"encoding/json"
)

func AddOrUpdate(query string) string {
	data := model.Credit{}
	if err := json.Unmarshal([]byte(query), &data); err != nil {
		local.HandleErrorMessage("Cant unmarshal list", err)
	}

	var response string
	var err error
	if data.Id == 0 {
		response = "added"
		err = infra.AddCredit(data.Name, data.FileName, data.Author, data.Link, data.Type, data.Licence)
	} else {
		response = "updated"
		err = infra.UpdateCredit(data.Id, data.Name, data.FileName, data.Author, data.Link, data.Type, data.Licence)
	}
	if err != nil {
		response = "error: " + err.Error()
	}

	return `{"status":"` + response + `"}`
}
