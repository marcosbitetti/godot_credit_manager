package commands

import (
	local "credits_manager/error"
	"credits_manager/infra"
	"encoding/json"
)

func List(query string) string {
	list := infra.ListCredits()
	out, err := json.Marshal(list)
	if err != nil {
		local.HandleErrorMessage("Cant marshal list", err)
	}
	return string(out)
}
