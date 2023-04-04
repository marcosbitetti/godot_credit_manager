package commands

import (
	local "credits_manager/error"
	"credits_manager/infra"
	"strconv"
)

func Delete(query string) string {
	id, err := strconv.Atoi(query)
	if err != nil {
		local.HandleErrorMessage("Cant recognize id", err)
	}
	infra.DeleteCredit(int64(id))

	return "[]"
}
