package commands

import (
	local "credits_manager/error"
	"credits_manager/infra"
	"encoding/json"
	"strings"
)

func List(query string) string {
	cmds := strings.SplitN(query, " ", 2)
	if len(cmds) < 2 {
		cmds = append(cmds, "")
	}
	ascDesc := cmds[0]
	search := cmds[1]
	list := infra.ListCredits(ascDesc, search)
	out, err := json.Marshal(list)
	if err != nil {
		local.HandleErrorMessage("Cant marshal list", err)
	}
	return string(out)
}
