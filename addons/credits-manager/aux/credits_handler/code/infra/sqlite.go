package infra

import (
	"context"
	local "credits_manager/error"
	"credits_manager/model"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

const AtribuitionModelPath = "ATRIBUITION_HANDLER_PATH"

var db *sql.DB

func getExecPath() string {
	ex, err := os.Getwd() //os.Executable()
	if err != nil {
		local.HandleErrorMessage("cant get executable path", err)
	}
	return ex //filepath.Dir(ex)
}

func getDatabasePath() string {
	path := fmt.Sprintf("%s/../source-database.json", getExecPath())
	// check 1st path
	_, err := os.Stat(path)
	if err != nil {
		if !os.IsNotExist(err) {
			panic(" path " + getExecPath() + " : " + err.Error())
		}
		// check 2sd path
		path = fmt.Sprintf("%s/source-database.json", getExecPath())
		_, err = os.Stat(path)
		if err != nil {
			if !os.IsNotExist(err) {
				panic("path " + getExecPath() + " : " + err.Error())
			}
			path = os.Getenv(AtribuitionModelPath)
			if path == "" {
				panic("path cant be reachable")
			}
			path = fmt.Sprintf("%s/source-database.json", path)
			_, err = os.Stat(path)
			if err != nil {
				panic("path cant be reachable: " + err.Error())
			}
		}
	}

	bytes, err := os.ReadFile(path)
	if err != nil {
		local.HandleErrorMessage("cant load data from source-database.json file", err)
	}
	var pathData struct {
		Database string `json:"database"`
	}
	err = json.Unmarshal(bytes, &pathData)
	if err != nil {
		local.HandleErrorMessage("malformated source-database.json file", err)
	}
	return pathData.Database
}

func OpenDatabase() error {
	var err error
	needDump := false
	if _, err := os.Stat(getDatabasePath()); err != nil {
		handler, err := os.Create(getDatabasePath())
		if err != nil {
			local.HandleErrorMessage("cant create database file", err)
		}
		defer func() {
			if err := handler.Close(); err != nil {
				local.HandleErrorMessage("closing db file", err)
			}
		}()
		needDump = true
	}

	db, err = sql.Open("sqlite3", getDatabasePath())
	if err != nil {
		local.HandleErrorMessage("cant open db file", err)
	}

	if err := createBaseTable(context.Background()); err != nil {
		local.HandleErrorMessage("cant create table", err)
	}

	if needDump {
		dumpFirstTypes()
		dumpFirstLicences()
	}

	return nil
}

func CloseDatabase() {
	if db == nil {
		return
	}
	err := db.Close()
	if err != nil {
		local.HandleErrorMessage("cant close database", err)
	}
}

func createBaseTable(ctx context.Context) error {
	var err error

	// types
	_, err = db.ExecContext(ctx, `
		CREATE TABLE IF NOT EXISTS types (
			_id 	INTEGER PRIMARY KEY NOT NULL,
			name	TEXT
		);
	`)
	if err != nil {
		local.HandleErrorMessage("cant run sql create table types", err)
	}

	// licences
	_, err = db.ExecContext(ctx, `
		CREATE TABLE IF NOT EXISTS licences (
			_id 	INTEGER PRIMARY KEY NOT NULL,
			name	TEXT DEFAULT "Attribution 4.0 International (CC BY 4.0)",
			link	TEXT DEFAULT "https://creativecommons.org/licenses/by/4.0/"
		);
	`)
	if err != nil {
		local.HandleErrorMessage("cant run sql create table types", err)
	}

	// credits
	_, err = db.ExecContext(ctx, `
		CREATE TABLE IF NOT EXISTS credits (
			_id 		INTEGER PRIMARY KEY NOT NULL,
			name		TEXT,
			filename	TEXT,
			type_id		INTEGER NOT NULL DEFAULT 1,
			author 		TEXT,
			link 		TEXT,
			licence_id 	INTEGER NOT NULL DEFAULT 1,
			FOREIGN KEY (type_id)
				REFERENCES types (_id)
					ON DELETE CASCADE
					ON UPDATE NO ACTION,
			FOREIGN KEY (licence_id)
				REFERENCES licences (_id)
					ON DELETE CASCADE
					ON UPDATE NO ACTION
		);
	`)
	if err != nil {
		local.HandleErrorMessage("cant run sql create table credits", err)
	}
	return nil
}

func dumpFirstTypes() {
	AddType("3D Model")
	AddType("Music")
	AddType("Plugin")
	AddType("Project")
	AddType("Sound Effect")
	AddType("Texture")
	AddType("Shader")
	AddType("Photo")
	AddType("Dubbing/Narration")
	AddType("Font")
	AddType("Code Snippet")
}

func dumpFirstLicences() {
	AddLicence("Attribution 4.0 International (CC BY 4.0)", "https://creativecommons.org/licenses/by/4.0/")
	AddLicence("Attribution-ShareAlike 4.0 International (CC BY-SA 4.0)", "https://creativecommons.org/licenses/by-sa/4.0/")
	AddLicence("Attribution-NonCommercial 4.0 International (CC BY-NC 4.0)", "https://creativecommons.org/licenses/by-nc/4.0/")
	AddLicence("Attribution-NonCommercial-ShareAlike 4.0 International (CC BY-NC-SA 4.0)", "https://creativecommons.org/licenses/by-nc-sa/4.0/")
	AddLicence("Attribution-NoDerivatives 4.0 International (CC BY-ND 4.0)", "https://creativecommons.org/licenses/by-nd/4.0/")
	AddLicence("Attribution-NonCommercial-NoDerivatives 4.0 International (CC BY-NC-ND 4.0)", "https://creativecommons.org/licenses/by-nc-nd/4.0/")
	AddLicence("CC0 1.0 Universal (CC0 1.0) - Public Domain Dedication", "https://creativecommons.org/publicdomain/zero/1.0/")
	AddLicence("MIT", "https://opensource.org/license/mit/")
	AddLicence("GNU General Public Licence", "https://www.gnu.org/licenses/gpl-3.0.html")
	AddLicence("Attribution-NonCommercial-ShareAlike 3.0 Unported (CC BY-NC-SA 3.0)", "https://creativecommons.org/licenses/by-nc-sa/3.0/")
	AddLicence("Attribution-NonCommercial-NoDerivs 3.0 Unported (CC BY-NC-ND 3.0)", "https://creativecommons.org/licenses/by-nc-nd/3.0/")
	AddLicence("Attribution-ShareAlike 3.0 Unported (CC BY-SA 3.0)", "https://creativecommons.org/licenses/by-sa/3.0/")
	AddLicence("Attribution-NoDerivs 3.0 Unported (CC BY-ND 3.0)", "https://creativecommons.org/licenses/by-nd/3.0/")
	AddLicence("Attribution 3.0 Unported (CC BY 3.0)", "https://creativecommons.org/licenses/by/3.0/")
	AddLicence("GNU Lesser General Public License (LGPL)", "https://www.gnu.org/licenses/lgpl-3.0.html")
	AddLicence("Apache License 2.0", "https://www.apache.org/licenses/LICENSE-2.0")
	AddLicence("Mozilla Public License 2.0", "https://www.mozilla.org/en-US/MPL/2.0/")
	AddLicence("Beerware", "https://fedoraproject.org/wiki/Licensing/Beerware")
	AddLicence("Royalty Free", "https://en.wikipedia.org/wiki/Royalty-free")
	AddLicence("Open Font License (OFL)", "https://openfontlicense.org/")
	AddLicence("OGA-BY 3.0 (Open Game Art)", "https://static.opengameart.org/OGA-BY-3.0.txt")
	AddLicence("Free Standard (Sketchfab)", "https://www.youtube.com/watch?v=M2bKt1oZsi4")
}

func ListCredits(ascDesc string, search string) []model.Credit {
	list := make([]model.Credit, 0)
	rows, err := db.Query(fmt.Sprintf(`
		SELECT c._id, c.name, filename, author, c.link,
			t.name as type,
			l.name as licence,
			l.link as licence_link
		FROM credits c
		LEFT JOIN types t ON t._id == c.type_id
		LEFT JOIN licences l ON l._id == c.licence_id
		%s
		ORDER BY c.name  COLLATE NOCASE %s
	`, mountQuery(search), ascDesc))
	if err != nil {
		local.HandleErrorMessage("cant read rows from credits", err)
	}
	defer func() {
		if err := rows.Close(); err != nil {
			local.HandleErrorMessage("cant close rows from credits", err)
		}
	}()
	for rows.Next() {
		data := model.Credit{}
		if err := rows.Scan(&data.Id, &data.Name, &data.FileName, &data.Author, &data.Link, &data.Type, &data.Licence, &data.LicenceUrl); err != nil {
			local.HandleErrorMessage("cant read row from credits", err)
		}
		list = append(list, data)
	}
	return list
}

func AddCredit(name string, fileame string, author string, link string, ctype string, licence string) error {
	stmt, err := db.Prepare(`
		INSERT InTO credits
		(name, filename, author, link, type_id, licence_id)
		VALUES
		(?, ?, ?, ?,
			(SELECT _id FROM types WHERE name=?),
			(SELECT _id FROM licences WHERE name=?)
		)
	`)
	if err != nil {
		local.HandleErrorMessage("cant prepare to add credit", err)
	}
	defer func() {
		if err := stmt.Close(); err != nil {
			local.HandleErrorMessage("cant close prepare to add credit", err)
		}
	}()
	_, err = stmt.Exec(name, fileame, author, link, ctype, licence)
	return err
}

func UpdateCredit(id int64, name string, fileame string, author string, link string, ctype string, licence string) error {
	stmt, err := db.Prepare(`
		UPDATE credits SET
			name=?,
			filename=?,
			author=?,
			link=?,
			type_id=(SELECT _id FROM types WHERE name=?),
			licence_id=(SELECT _id FROM licences WHERE name=?)
		WHERE _id = ?
	`)
	if err != nil {
		local.HandleErrorMessage("cant prepare to add credit", err)
	}
	defer func() {
		if err := stmt.Close(); err != nil {
			local.HandleErrorMessage("cant close prepare to add credit", err)
		}
	}()
	_, err = stmt.Exec(name, fileame, author, link, ctype, licence, id)
	return err
}

func DeleteCredit(id int64) error {
	stmt, err := db.Prepare(`DELETE FROM credits WHERE _id = ?`)
	if err != nil {
		local.HandleErrorMessage("cant prepare to delete credit", err)
	}
	defer func() {
		if err := stmt.Close(); err != nil {
			local.HandleErrorMessage("cant close prepare to delete credit", err)
		}
	}()
	_, err = stmt.Exec(id)
	return err
}

func ListTypes() []model.Type {
	list := make([]model.Type, 0)
	rows, err := db.Query(`
		SELECT _id, name FROM types ORDER BY name COLLATE NOCASE ASC
	`)
	if err != nil {
		local.HandleErrorMessage("cant read rows from typess", err)
	}
	defer func() {
		if err := rows.Close(); err != nil {
			local.HandleErrorMessage("cant close rows from typess", err)
		}
	}()
	for rows.Next() {
		data := model.Type{}
		if err := rows.Scan(&data.Id, &data.Name); err != nil {
			local.HandleErrorMessage("cant read row from types", err)
		}
		list = append(list, data)
	}
	return list
}

func AddType(name string) error {
	stmt, err := db.Prepare(`
		INSERT INTO types(name) VALUES(?);
	`)
	if err != nil {
		local.HandleErrorMessage("cant prepare to add type", err)
	}

	defer func() {
		if err := stmt.Close(); err != nil {
			local.HandleErrorMessage("cant close prepare to add type", err)
		}
	}()

	_, err = stmt.Exec(name)
	return err
}

func UpdateType(id int64, name string) error {
	stmt, err := db.Prepare(`
		UPDATE types SET name=? WHERE _id=?
	`)
	if err != nil {
		local.HandleErrorMessage("cant prepare to update type", err)
	}

	defer func() {
		if err := stmt.Close(); err != nil {
			local.HandleErrorMessage("cant close prepare to add type", err)
		}
	}()

	_, err = stmt.Exec(name, id)
	return err
}

func DeleteType(id int64) error {
	stmt, err := db.Prepare(`DELETE FROM types WHERE _id = ?`)
	if err != nil {
		local.HandleErrorMessage("cant prepare to delete type", err)
	}
	defer func() {
		if err := stmt.Close(); err != nil {
			local.HandleErrorMessage("cant close prepare to delete type", err)
		}
	}()
	_, err = stmt.Exec(id)
	return err
}

func ListLicences() []model.Licence {
	list := make([]model.Licence, 0)
	rows, err := db.Query(`
		SELECT _id, name, link FROM licences ORDER BY name COLLATE NOCASE ASC
	`)
	if err != nil {
		local.HandleErrorMessage("cant read rows from licences", err)
	}
	defer func() {
		if err := rows.Close(); err != nil {
			local.HandleErrorMessage("cant close rows from licences", err)
		}
	}()
	for rows.Next() {
		data := model.Licence{}
		if err := rows.Scan(&data.Id, &data.Name, &data.Link); err != nil {
			local.HandleErrorMessage("cant read row from licences", err)
		}
		list = append(list, data)
	}
	return list
}

func AddLicence(name string, link string) error {
	stmt, err := db.Prepare(`
		INSERT INTO licences(name, link) VALUES(?, ?);
	`)
	if err != nil {
		local.HandleErrorMessage("cant prepare to add licence", err)
	}

	defer func() {
		if err := stmt.Close(); err != nil {
			local.HandleErrorMessage("cant close prepare to add licence", err)
		}
	}()

	_, err = stmt.Exec(name, link)
	return err
}

func UpdateLicence(id int64, name string, link string) error {
	stmt, err := db.Prepare(`
		UPDATE licences SET name=?, link=? WHERE _id=?;
	`)
	if err != nil {
		local.HandleErrorMessage("cant prepare to update licence", err)
	}

	defer func() {
		if err := stmt.Close(); err != nil {
			local.HandleErrorMessage("cant close prepare to update licence", err)
		}
	}()

	_, err = stmt.Exec(name, link, id)
	return err
}

func DeleteLicence(id int64) error {
	stmt, err := db.Prepare(`DELETE FROM licences WHERE _id = ?`)
	if err != nil {
		local.HandleErrorMessage("cant prepare to delete licence", err)
	}
	defer func() {
		if err := stmt.Close(); err != nil {
			local.HandleErrorMessage("cant close prepare to delete licences", err)
		}
	}()
	_, err = stmt.Exec(id)
	return err
}

func mountQuery(q string) string {
	if q == "" {
		return ""
	}
	tokens := strings.Fields(q)
	joined := "%" + strings.Join(tokens, "%") + "%"
	return fmt.Sprintf("WHERE c.name LIKE '%s' OR c.author LIKE '%s'", joined, joined)
}

func CountAssociatedCredits(table string, id int64) int64 {
	column := table + "_id"
	rows, err := db.Query(fmt.Sprintf(`SELECT COUNT(_id) AS total FROM credits WHERE %s = %d`, column, id))
	if err != nil {
		local.HandleErrorMessage("cant read rows from count credits", err)
	}
	defer func() {
		if err := rows.Close(); err != nil {
			local.HandleErrorMessage("cant close rows from count credits", err)
		}
	}()
	var total int64
	for rows.Next() {
		if err := rows.Scan(&total); err != nil {
			local.HandleErrorMessage("cant read row from count credits", err)
		}
	}
	return total
}

func ListField(search string, field string) []string {
	list := make([]string, 0)
	search = search + "%"
	stmt, err := db.Prepare(fmt.Sprintf(`SELECT %s FROM credits WHERE %s LIKE ?`, field, field))
	if err != nil {
		local.HandleErrorMessage("cant prepare to list field", err)
	}
	defer func() {
		if err := stmt.Close(); err != nil {
			local.HandleErrorMessage("cant close prepare list field", err)
		}
	}()
	rows, err := stmt.Query(search)
	if err != nil {
		local.HandleErrorMessage("cant read rows from list fields", err)
	}
	defer func() {
		if err := rows.Close(); err != nil {
			local.HandleErrorMessage("cant close rows from list fields", err)
		}
	}()
	for rows.Next() {
		data := ""
		if err := rows.Scan(&data); err != nil {
			local.HandleErrorMessage("cant read row from list fields", err)
		}
		list = append(list, data)
	}
	return list
}

func CreateEmptyDatabase(path string) (*sql.DB, error) {
	var err error
	db, err = sql.Open("sqlite3", path)
	if err != nil {
		return nil, errors.Join(err, errors.New("cant open/create file"))
	}
	if err := createBaseTable(context.Background()); err != nil {
		return nil, errors.Join(err, errors.New("cant create table"))
	}
	dumpFirstTypes()
	dumpFirstLicences()
	return db, nil
}
