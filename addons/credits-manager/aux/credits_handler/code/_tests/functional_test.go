package _tests

import (
	"bytes"
	"credits_manager/app"
	"credits_manager/infra"
	"credits_manager/model"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"strings"
	"testing"
)

func TestListCreditsForFirstTime(t *testing.T) {
	prepareTest(t)
	prepareDatabase(t)
	t.Run("return empty list", func(t *testing.T) {
		app.Start([]string{getExecPath(), "list", "asc"})
		result := freePrint()
		if result != "[]" {
			t.Errorf("wrong empty result: %s", result)
		}
	})
}

func TestAddCredit(t *testing.T) {
	prepareTest(t)
	prepareDatabase(t)
	t.Run("return list with new record", func(t *testing.T) {
		app.Start([]string{getExecPath(), "add",
			`{"name":"Work","filename":"res://file","author":"Joe","link":"http://...","type":"Music","licence":"MIT"}`,
		})
		result := freePrint()
		if result != `{"status":"added"}` {
			t.Errorf("wrong result: %s", result)
		}
		app.Start([]string{getExecPath(), "list", "asc"})
		result = freePrint()
		if result != `[{"_id":1,"name":"Work","filename":"res://file","type":"Music","author":"Joe","link":"http://...","licence":"MIT","licenceUrl":"https://opensource.org/license/mit/"}]` {
			t.Errorf("wrong result: %s", result)
		}
	})
}

func TestIfFileReferenceAlreadyExists(t *testing.T) {
	prepareTest(t)
	prepareDatabase(t)
	firstStep := func() {
		app.Start([]string{getExecPath(), "add",
			`{"name":"Work","filename":"file_test","author":"Joe","link":"http://...","type":"Music","licence":"MIT"}`,
		})
		result := freePrint()
		if result != `{"status":"added"}` {
			t.Errorf("wrong result: %s", result)
		}
	}
	t.Run("return true if file already credited", func(t *testing.T) {
		firstStep()
		app.Start([]string{getExecPath(), "file-exists", "file_test"})
		result := freePrint()
		if result != `{"exists":true}` {
			t.Errorf("wrong result: %s", result)
		}
	})
	t.Run("return false if file not credited", func(t *testing.T) {
		firstStep()
		app.Start([]string{getExecPath(), "file-exists", "file_new_file"})
		result := freePrint()
		if result != `{"exists":false}` {
			t.Errorf("wrong result: %s", result)
		}
	})
}

func TestFilteredListCreditsByNameOrAuthor(t *testing.T) {
	prepareTest(t)
	db := prepareDatabase(t)
	dumpRegisters(t, db)
	t.Run("return jose list", func(t *testing.T) {
		app.Start([]string{getExecPath(), "list", "asc", "jose"})
		var list []model.Credit
		if err := json.Unmarshal([]byte(freePrint()), &list); err != nil {
			t.Error(err.Error())
		}
		if len(list) != 2 {
			t.Error("cant match regis")
		}
	})
}

func TestDeleteCredit(t *testing.T) {
	prepareTest(t)
	db := prepareDatabase(t)
	dumpRegisters(t, db)
	t.Run("return a list with minus one register", func(t *testing.T) {
		var list []model.Credit
		app.Start([]string{getExecPath(), "list", "asc"})
		if err := json.Unmarshal([]byte(freePrint()), &list); err != nil {
			t.Error(err.Error())
		}
		if len(list) != 3 {
			t.Error("failed to load list")
		}
		app.Start([]string{getExecPath(), "delete", "1"})
		freePrint()
		app.Start([]string{getExecPath(), "list", "asc"})
		if err := json.Unmarshal([]byte(freePrint()), &list); err != nil {
			t.Error(err.Error())
		}
		if len(list) != 2 {
			t.Error("failed to delete")
		}
	})
}

func TestListLicences(t *testing.T) {
	prepareTest(t)
	prepareDatabase(t)
	t.Run("return full list", func(t *testing.T) {
		var list []model.Licence
		app.Start([]string{getExecPath(), "licences", "asc"})
		if err := json.Unmarshal([]byte(freePrint()), &list); err != nil {
			t.Error(err.Error())
		}
		if len(list) == 0 {
			t.Error("failed to load list")
		}
	})
}

func TestAddLicence(t *testing.T) {
	prepareTest(t)
	prepareDatabase(t)
	t.Run("return new item", func(t *testing.T) {
		var list []model.Licence
		app.Start([]string{getExecPath(), "licences", "asc"})
		if err := json.Unmarshal([]byte(freePrint()), &list); err != nil {
			t.Error(err.Error())
		}
		expectedListSize := len(list) + 1
		app.Start([]string{getExecPath(), "add-licence", `{"name":"_teste", "link":"http://test.tst"}`})
		freePrint()
		app.Start([]string{getExecPath(), "licences", "asc"})
		if err := json.Unmarshal([]byte(freePrint()), &list); err != nil {
			t.Error(err.Error())
		}
		if len(list) != expectedListSize {
			t.Error("failed to add licence")
		}
	})
}

func TestUpdateLicence(t *testing.T) {
	prepareTest(t)
	prepareDatabase(t)
	t.Run("return empty list", func(t *testing.T) {
		app.Start([]string{getExecPath(), "update-licence", `{"_id":1,"name":"_teste_changed", "link":"http://test.tst"}`})
		freePrint()
		app.Start([]string{getExecPath(), "licences", "asc"})
		if !strings.Contains(freePrint(), "_teste_changed") {
			t.Error("cant update licence")
		}
	})
}

func TestDeleteLicenceSoft(t *testing.T) {
	t.Skip("not implemented yet")
}

func TestDeleteLicenceHard(t *testing.T) {
	prepareTest(t)
	prepareDatabase(t)
	t.Run("return list minus one", func(t *testing.T) {
		var list []model.Licence
		app.Start([]string{getExecPath(), "licences", "asc"})
		if err := json.Unmarshal([]byte(freePrint()), &list); err != nil {
			t.Error(err.Error())
		}
		expectedListSize := len(list) - 1
		app.Start([]string{getExecPath(), "delete-licence", "1"})
		freePrint()
		app.Start([]string{getExecPath(), "licences", "asc"})
		if err := json.Unmarshal([]byte(freePrint()), &list); err != nil {
			t.Error(err.Error())
		}
		if len(list) != expectedListSize {
			t.Error("failed to delete licence")
		}
	})
}

func TestListTypes(t *testing.T) {
	prepareTest(t)
	prepareDatabase(t)
	t.Run("return full list", func(t *testing.T) {
		var list []model.Type
		app.Start([]string{getExecPath(), "types", "asc"})
		if err := json.Unmarshal([]byte(freePrint()), &list); err != nil {
			t.Error(err.Error())
		}
		if len(list) == 0 {
			t.Error("cant read list of types")
		}
	})
}

func TestAddType(t *testing.T) {
	prepareTest(t)
	prepareDatabase(t)
	t.Run("return new item", func(t *testing.T) {
		var list []model.Type
		app.Start([]string{getExecPath(), "types", "asc"})
		if err := json.Unmarshal([]byte(freePrint()), &list); err != nil {
			t.Error(err.Error())
		}
		expectedListSize := len(list) + 1
		app.Start([]string{getExecPath(), "add-type", `{"name":"_teste"}`})
		freePrint()
		app.Start([]string{getExecPath(), "types", "asc"})
		if err := json.Unmarshal([]byte(freePrint()), &list); err != nil {
			t.Error(err.Error())
		}
		if len(list) != expectedListSize {
			t.Error("cant add type")
		}
	})
}

func TestUpdateType(t *testing.T) {
	prepareTest(t)
	prepareDatabase(t)
	t.Run("return empty list", func(t *testing.T) {
		app.Start([]string{getExecPath(), "update-type", `{"_id":1,"name":"_teste_changed"}`})
		freePrint()
		app.Start([]string{getExecPath(), "types", "asc"})
		if !strings.Contains(freePrint(), "_teste_changed") {
			t.Error("cant update type")
		}
	})
}

func TestDeleteTypeSoft(t *testing.T) {
	t.Skip("not implemented yet")
}

func TestDeleteTypeHard(t *testing.T) {
	prepareTest(t)
	prepareDatabase(t)
	t.Run("return empty list", func(t *testing.T) {
		var list []model.Type
		app.Start([]string{getExecPath(), "types", "asc"})
		if err := json.Unmarshal([]byte(freePrint()), &list); err != nil {
			t.Error(err.Error())
		}
		expectedListSize := len(list) - 1
		app.Start([]string{getExecPath(), "delete-type", "1"})
		freePrint()
		app.Start([]string{getExecPath(), "types", "asc"})
		if err := json.Unmarshal([]byte(freePrint()), &list); err != nil {
			t.Error(err.Error())
		}
		if len(list) != expectedListSize {
			t.Error("cant delete type")
		}
	})
}

func TestAutoCompleteAuthor(t *testing.T) {
	prepareTest(t)
	db := prepareDatabase(t)
	dumpRegisters(t, db)
	t.Run("return 3 entries", func(t *testing.T) {
		app.Start([]string{getExecPath(), "auto-complete-author", `{"text":"j"}`})
		var list []string
		if err := json.Unmarshal([]byte(freePrint()), &list); err != nil {
			t.Error(err.Error())
		}
		if len(list) != 3 {
			t.Error("cant match regis")
		}
	})
	t.Run("return 1 entries", func(t *testing.T) {
		app.Start([]string{getExecPath(), "auto-complete-author", `{"text":"jA"}`})
		var list []string
		if err := json.Unmarshal([]byte(freePrint()), &list); err != nil {
			t.Error(err.Error())
		}
		if len(list) != 1 {
			t.Error("cant match regis")
		}
	})
}

var expectedCSV = `work,author,link,licence,description
thing 1,jose,https://link,Attribution 4.0 International (CC BY 4.0),3D Model
thing 2,jose,https://link,Attribution 4.0 International (CC BY 4.0),3D Model
thing 3,jack,https://link,Attribution 4.0 International (CC BY 4.0),3D Model
`

func TestExportCSV(t *testing.T) {
	prepareTest(t)
	db := prepareDatabase(t)
	dumpRegisters(t, db)
	t.Run("return jose list", func(t *testing.T) {
		app.Start([]string{getExecPath(), "csv", `{
			"path" : "/tmp/",
			"attribuitions" : true,
			"licences" : true
		}`})
	})
	if freePrint() != "saved" {
		t.Error("cant write csv")
	}
	file, err := os.ReadFile("/tmp/works.csv")
	if err != nil {
		t.Error(err.Error())
	}
	if string(file) != expectedCSV {
		t.Error(err.Error())
	}
}

// ********

var output bytes.Buffer
var sqlite_path string

func customPrint(text string) {
	output.WriteString(text)
}

func freePrint() string {
	out := output.String()
	output.Reset()
	return out
}

func prepareTest(t *testing.T) {
	sqlite_path = fmt.Sprintf("%s/credits_tst.sqlitedb", getExecPath())
	raw, err := json.Marshal(map[string]string{
		"database": sqlite_path,
	})
	if err != nil {
		t.Error(err.Error())
	}
	filename := getExecPath() + "/source-database.json"
	if err := os.WriteFile(filename, raw, fs.FileMode(os.O_CREATE|os.O_RDWR)); err != nil {
		t.Error(err.Error())
	}
	if err := os.Chmod(filename, os.FileMode(0666)); err != nil {
		t.Error(err.Error())
	}

	app.SetCaller(customPrint)
	freePrint()
}

func getExecPath() string {
	ex, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return ex
}

func dumpRegisters(t *testing.T, db *sql.DB) {
	if err := addCredit(db, "thing 1", "jose"); err != nil {
		t.Error(err.Error())
	}
	if err := addCredit(db, "thing 2", "jose"); err != nil {
		t.Error(err.Error())
	}
	if err := addCredit(db, "thing 3", "jack"); err != nil {
		t.Error(err.Error())
	}
}

func addCredit(db *sql.DB, name string, author string) error {
	stmt, err := db.Prepare(`
		INSERT InTO credits
		(name, filename, author, link, type_id, licence_id)
		VALUES
		(?, ?, ?, ?, 1, 1)
	`)
	if err != nil {
		return err
	}
	defer func() {
		if err := stmt.Close(); err != nil {
			println(err.Error())
		}
	}()
	_, err = stmt.Exec(name, "fileame", author, "https://link")
	return err
}

func prepareDatabase(t *testing.T) *sql.DB {
	_, err := os.Stat(sqlite_path)
	if err != nil {
		if !os.IsNotExist(err) {
			t.Error("erros creating database")
		}
	} else {
		err = os.Remove(sqlite_path)
		if err != nil {
			t.Error("erros creating database")
		}
	}
	db, err := infra.CreateEmptyDatabase(sqlite_path)
	if err != nil {
		t.Error(err.Error())
	}
	return db
}
