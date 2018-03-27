package ghstats

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"
)

type Table struct {
	Name    string
	Columns []Column
}

type Column struct {
	Name string
	Type string
}

type SQLable interface {
	Table() Table
	Values() []interface{}
}

type Backend struct {
	db *sql.DB
}

func tableToNames(t Table) string {
	v := make([]string, len(t.Columns))
	for i, c := range t.Columns {
		v[i] = c.Name
	}
	return strings.Join(v, ",")
}

func tableToNamesAndTypes(t Table) string {
	v := make([]string, len(t.Columns))
	for i, c := range t.Columns {
		v[i] = c.Name + " " + c.Type
	}
	return strings.Join(v, ",")
}

func GetBackend(db *sql.DB) *Backend {
	return &Backend{db}
}
func (b *Backend) Store(s SQLable) error {
	values := s.Values()
	columns := len(s.Table().Columns)
	items := len(values) / columns

	stmt := fmt.Sprintf("INSERT INTO %s (%s) VALUES", s.Table().Name, tableToNames(s.Table()))

	placeholder := fmt.Sprintf("(%s)", strings.TrimRight(strings.Repeat("?,", columns), ","))
	placeholder = strings.TrimRight(strings.Repeat(placeholder+",", items), ",")
	stmt += placeholder

	_, err := b.db.Exec(stmt, values...)
	return err
}

func (b *Backend) CreateTables(args ...SQLable) error {
	for _, s := range args {
		stmt := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (%s);", s.Table().Name, tableToNamesAndTypes(s.Table()))
		if _, err := b.db.Exec(stmt); err != nil {
			return errors.New(fmt.Sprintf(`error while executing "%v": %v`, stmt, err))
		}
	}
	return nil
}
