package commands

import (
	local "credits_manager/error"
	"credits_manager/infra"
	"encoding/json"
)

func ListTypes(query string) string {
	list := infra.ListTypes()
	out, err := json.Marshal(list)
	if err != nil {
		local.HandleErrorMessage("Cant marshal types list", err)
	}
	return string(out)
}
