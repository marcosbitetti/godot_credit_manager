package app

import (
	"credits_manager/commands"
	local "credits_manager/error"
	"credits_manager/infra"
	"strings"
)

var commandList map[string]func(query string) string = map[string]func(query string) string{
	"list":     commands.List,
	"add":      commands.AddOrUpdate,
	"update":   commands.AddOrUpdate,
	"delete":   commands.Delete,
	"types":    commands.ListTypes,
	"licences": commands.ListLicences,
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
		print(`{"error":"error"}`)
		return
	}
	print(out)
}
