package main

import (
	"credits_manager/app"
	"testing"
)

func TestList(t *testing.T) {
	t.Run("return empty list", func(t *testing.T) {
		app.Start([]string{"_executable", "list", "asc"})
	})
}

func TestFilteredListByNameOrAuthor(t *testing.T) {
	t.Run("return empty list", func(t *testing.T) {
		app.Start([]string{"_executable", "list", "asc", "wood"})
	})
}

func TestListLicences(t *testing.T) {
	t.Run("return empty list", func(t *testing.T) {
		app.Start([]string{"_executable", "licences", "asc"})
	})
}

func TestListTypes(t *testing.T) {
	t.Run("return empty list", func(t *testing.T) {
		app.Start([]string{"_executable", "types", "asc"})
	})
}
