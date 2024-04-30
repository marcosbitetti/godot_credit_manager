package commands

import (
	local "credits_manager/error"
	"credits_manager/infra"
	"credits_manager/model"
	"encoding/json"
	"strings"
)

func AutoCompleteAuthor(query string) string {
	data := model.Query{}
	if err := json.Unmarshal([]byte(query), &data); err != nil {
		local.HandleErrorMessage("Cant unmarshal list", err)
	}
	search := strings.TrimSpace(strings.ToLower(data.Text))
	out, err := json.Marshal(infra.ListField(search, "author"))
	if err != nil {
		local.HandleErrorMessage("Cant marshal licences list", err)
	}
	return string(out)
}
