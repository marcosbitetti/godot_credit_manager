package commands

import (
	local "credits_manager/error"
	"credits_manager/infra"
	"encoding/json"
)

func ListLicences(query string) string {
	list := infra.ListLicences()
	out, err := json.Marshal(list)
	if err != nil {
		local.HandleErrorMessage("Cant marshal licences list", err)
	}
	return string(out)
}
