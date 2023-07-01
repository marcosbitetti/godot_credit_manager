package commands

import (
	local "credits_manager/error"
	"credits_manager/infra"
	"strconv"
)

func DeleteType(query string) string {
	id, err := strconv.Atoi(query)
	if err != nil {
		local.HandleErrorMessage("Cant recognize id", err)
	}
	associatedLicences := infra.CountAssociatedCredits("type", int64(id))
	if associatedLicences > 0 {
		infra.UpdateType(int64(id), "[deleted]")
	} else {
		infra.DeleteType(int64(id))
	}
	return "[]"
}
