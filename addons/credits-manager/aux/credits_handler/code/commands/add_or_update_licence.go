package commands

import (
	local "credits_manager/error"
	"credits_manager/infra"
	"credits_manager/model"
	"encoding/json"
)

func AddOrUpdateLicence(query string) string {
	data := model.Licence{}
	if err := json.Unmarshal([]byte(query), &data); err != nil {
		local.HandleErrorMessage("Cant unmarshal list", err)
	}

	if data.Id == 0 {
		infra.AddLicence(data.Name, data.Link)
	} else {
		infra.UpdateLicence(data.Id, data.Name, data.Link)
	}

	return "[]"
}
