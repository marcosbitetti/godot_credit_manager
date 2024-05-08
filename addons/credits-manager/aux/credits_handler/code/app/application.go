package app

import (
	"credits_manager/commands"
	local "credits_manager/error"
	"credits_manager/infra"
	"strings"
)

var commandList map[string]func(query string) string = map[string]func(query string) string{
	"list":   commands.List,
	"add":    commands.AddOrUpdate,
	"update": commands.AddOrUpdate,
	"delete": commands.Delete,

	"licences":       commands.ListLicences,
	"add-licence":    commands.AddOrUpdateLicence,
	"update-licence": commands.AddOrUpdateLicence,
	"delete-licence": commands.DeleteLicence,

	"types":       commands.ListTypes,
	"add-type":    commands.AddOrUpdateType,
	"update-type": commands.AddOrUpdateType,
	"delete-type": commands.DeleteType,

	"auto-complete-author": commands.AutoCompleteAuthor,

	"csv": commands.ExportCSV,
}

func Start(com []string) {
	if len(com) < 3 {
		local.Message("command error")
		return
	}
	infra.OpenDatabase()
	defer infra.CloseDatabase()
	callable, has := commandList[com[1]]
	if !has {
		local.Message("command not found")
		return
	}

	commandStr := make([]string, 0)
	for i := 2; i < len(com); i++ {
		commandStr = append(commandStr, com[i])
	}

	out := callable(strings.Join(commandStr, " "))
	if out == "" {
		Print(`{"error":"error"}`)
		return
	}
	Print(out)
}
