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

	if data.Id == 0 {
		infra.AddCredit(data.Name, data.FileName, data.Author, data.Link, data.Type, data.Licence)
	} else {
		infra.UpdateCredit(data.Id, data.Name, data.FileName, data.Author, data.Link, data.Type, data.Licence)
	}

	return "[]"
}
