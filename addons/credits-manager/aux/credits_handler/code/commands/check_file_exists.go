package commands

import (
	"credits_manager/infra"
	"strings"
)

func FileExists(query string) string {
	name := strings.TrimSpace(query)
	if infra.FileExists(name) {
		return `{"exists":true}`
	}
	return `{"exists":false}`
}
