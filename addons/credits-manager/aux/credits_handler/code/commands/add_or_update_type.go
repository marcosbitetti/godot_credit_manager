package commands

import (
	local "credits_manager/error"
	"credits_manager/infra"
	"credits_manager/model"
	"encoding/json"
)

func AddOrUpdateType(query string) string {
	data := model.Type{}
	if err := json.Unmarshal([]byte(query), &data); err != nil {
		local.HandleErrorMessage("Cant unmarshal list", err)
	}

	if data.Id == 0 {
		infra.AddType(data.Name)
	} else {
		infra.UpdateType(data.Id, data.Name)
	}

	return "[]"
}
