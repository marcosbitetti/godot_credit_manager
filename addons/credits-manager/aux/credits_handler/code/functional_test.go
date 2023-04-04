package main

import (
	"credits_manager/app"
	"testing"
)

func TestEmptyList(t *testing.T) {
	t.Run("return empty list", func(t *testing.T) {
		app.Start([]string{"list", "asc"})
	})
}
