package ghstats

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

const DBName = "./test.gbstats.db"

type Dummy struct {
	Id   int
	Name string
}

func (d *Dummy) Table() Table {
	return Table{"dummy", []Column{
		Column{"id", "INTEGER"},
		Column{"name", "TEXT"},
	}}
}

func (d *Dummy) Values() []interface{} {
	return []interface{}{1, "antani", 2, "vicesindaco"}
}

func TestCreateTables(t *testing.T) {
	assert := assert.New(t)
	os.Remove(DBName)
	db, err := sql.Open("sqlite3", DBName)

	b := GetBackend(db)
	d := &Dummy{}
	table := d.Table()

	b.CreateTables(d)

	var rows *sql.Rows
	var sql string
	rows, err = db.Query("SELECT sql FROM sqlite_master WHERE type='table' AND name=?;", table.Name)
	defer rows.Close()

	rows.Next()
	err = rows.Scan(&sql)
	assert.Nil(err)
	expect := fmt.Sprintf("CREATE TABLE %s (%s %s,%s %s)",
		table.Name,
		table.Columns[0].Name,
		table.Columns[0].Type,
		table.Columns[1].Name,
		table.Columns[1].Type)
	assert.Equal(expect, sql)

	os.Remove(DBName)
	db.Close()
}

func TestInsertValue(t *testing.T) {
	assert := assert.New(t)
	os.Remove(DBName)
	db, err := sql.Open("sqlite3", DBName)

	b := GetBackend(db)
	d := Dummy{}
	//table := d.Table()

	b.CreateTables(&d)
	err = b.Insert(&d)
	assert.Nil(err)

	var rows *sql.Rows
	rows, err = db.Query("SELECT id, name FROM dummy;")
	// rows, err = db.Query("SELECT ?, ? FROM ?;", table.Columns[0].Name, table.Columns[1].Name, table.Name)
	assert.Nil(err)
	defer rows.Close()
	e := Dummy{1, "antani"}
	r := Dummy{}

	rows.Next()
	err = rows.Scan(&r.Id, &r.Name)
	assert.Equal(e, r)

	rows.Next()
	e = Dummy{2, "vicesindaco"}
	r = Dummy{}
	err = rows.Scan(&r.Id, &r.Name)
	assert.Equal(e, r)

	os.Remove(DBName)
	db.Close()
}
