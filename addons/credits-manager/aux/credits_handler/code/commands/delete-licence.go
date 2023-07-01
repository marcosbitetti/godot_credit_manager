package commands

import (
	local "credits_manager/error"
	"credits_manager/infra"
	"strconv"
)

func DeleteLicence(query string) string {
	id, err := strconv.Atoi(query)
	if err != nil {
		local.HandleErrorMessage("Cant recognize id", err)
	}
	associatedLicences := infra.CountAssociatedCredits("licence", int64(id))
	if associatedLicences > 0 {
		infra.UpdateLicence(int64(id), "[deleted]", "")
	} else {
		infra.DeleteLicence(int64(id))
	}
	return "[]"
}
