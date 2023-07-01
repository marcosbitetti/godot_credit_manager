package main

import (
	"credits_manager/app"
	"io/ioutil"
	"os"
	"testing"
)

func TestListCredits(t *testing.T) {
	t.Run("return empty list", func(t *testing.T) {
		app.Start([]string{"_executable", "list", "asc"})
	})
}

func TestFilteredListCreditsByNameOrAuthor(t *testing.T) {
	t.Run("return empty list", func(t *testing.T) {
		app.Start([]string{"_executable", "list", "asc", "wood"})
	})
}

func TestDeleteCredit(t *testing.T) {
	t.Run("return empty list", func(t *testing.T) {
		app.Start([]string{"_executable", "list", "asc"})
		app.Start([]string{"_executable", "delete", "19"})
		app.Start([]string{"_executable", "list", "asc"})
	})
}

func TestListLicences(t *testing.T) {
	t.Run("return empty list", func(t *testing.T) {
		app.Start([]string{"_executable", "licences", "asc"})
	})
}

func TestAddLicence(t *testing.T) {
	t.Run("return empty list", func(t *testing.T) {
		app.Start([]string{"_executable", "add-licence", `{"name":"_teste", "link":"http://test.tst"}`})
		app.Start([]string{"_executable", "licences", "asc"})
	})
}

func TestUpdateLicence(t *testing.T) {
	t.Run("return empty list", func(t *testing.T) {
		app.Start([]string{"_executable", "update-licence", `{"_id":19,"name":"_teste_changed", "link":"http://test.tst"}`})
		app.Start([]string{"_executable", "licences", "asc"})
	})
}

func TestDeleteLicenceSoft(t *testing.T) {
	t.Run("return empty list", func(t *testing.T) {
		app.Start([]string{"_executable", "delete-licence", "1"})
		app.Start([]string{"_executable", "licences", "asc"})
	})
}

func TestDeleteLicenceHard(t *testing.T) {
	t.Run("return empty list", func(t *testing.T) {
		app.Start([]string{"_executable", "delete-licence", "19"})
		app.Start([]string{"_executable", "licences", "asc"})
	})
}

func TestListTypes(t *testing.T) {
	t.Run("return empty list", func(t *testing.T) {
		app.Start([]string{"_executable", "types", "asc"})
	})
}

func TestAddType(t *testing.T) {
	t.Run("return empty list", func(t *testing.T) {
		app.Start([]string{"_executable", "add-type", `{"name":"_teste"}`})
		app.Start([]string{"_executable", "types", "asc"})
	})
}

func TestUpdateType(t *testing.T) {
	t.Run("return empty list", func(t *testing.T) {
		app.Start([]string{"_executable", "update-type", `{"_id":10,"name":"_teste_changed"}`})
		app.Start([]string{"_executable", "types", "asc"})
	})
}

func TestDeleteTypeSoft(t *testing.T) {
	t.Run("return empty list", func(t *testing.T) {
		app.Start([]string{"_executable", "delete-type", "6"})
		app.Start([]string{"_executable", "types", "asc"})
	})
}

func TestDeleteTypeHard(t *testing.T) {
	t.Run("return empty list", func(t *testing.T) {
		app.Start([]string{"_executable", "delete-type", "10"})
		app.Start([]string{"_executable", "types", "asc"})
	})
}

// ********

var stdoutBuffer os.File //io.WriteSeeker //  bytes.Buffer
var oldStream *os.File

func catchBuffer() {
	oldStream = os.Stdout
	os.Stdout = &stdoutBuffer
}

func releaseBuffer() {
	os.Stdout = oldStream
	println("OUT")
	//stdoutBuffer.Close()
	content, err := ioutil.ReadAll(&stdoutBuffer)
	println(err)
	println(string(content))
}
