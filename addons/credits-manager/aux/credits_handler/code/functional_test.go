package main

import (
	"credits_manager/app"
	"testing"
)

func TestEmptyList(t *testing.T) {
	t.Run("return empty list", func(t *testing.T) {
		app.Start([]string{"_executable", "list", "asc"})
	})
}

func TestFilteredListByNameOrAuthor(t *testing.T) {
	t.Run("return empty list", func(t *testing.T) {
		app.Start([]string{"_executable", "list", "asc", "wood"})
	})
}
