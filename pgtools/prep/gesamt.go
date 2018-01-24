package prep

import (
	"fmt"
	"io"
	"path/filepath"

	"github.com/jackc/pgx"
	"github.com/pkg/errors"

	"pgxgenerate/pgtools/db"
)

var writeHeader func(w io.Writer, header string)

func dirSetter(basis string, packname string) {
	importDB := filepath.Join(basis, "pgtools/db")

	writeHeader = func(w io.Writer, header string) {

		fmt.Fprintf(w, `package %s
			import %q
			%s
			`, packname, importDB, header)
	}

}

func GesamtPrep(prefix string, importpre string, basis string, schema string) error {
	prepTypes := make(map[string]prepTyp)
	var dir = "../.."
	if err := PrepList(prefix, &prepTypes, schema); err != nil {
		return errors.Wrapf(err, "Views nicht gefunden")
	}
	fmt.Println("vor generierung weiter", len(prepTypes))
	dirSetter(basis, "generprep")

	if err := GenerPrep(dir, filepath.Join(basis, importpre, "generprep"), &prepTypes); err != nil {
		return errors.Wrapf(err, "generierungsfehler")
	}
	return nil
}

func Gesamt(importpre string, basis string) (bool, error) {
	fmt.Println("vor generierung prep")
	var dir = "../.."
	sql, err := ReadSQL(filepath.Join(dir, basis, importpre, "sqls"))
	if err != nil {
		return false, err
	}
	prepTypes := make(map[string]prepTyp)
	prepStmt := make(map[string]*pgx.PreparedStatement)
	dirSetter(basis, "gener")
	if err := Prep(sql, &prepTypes, &prepStmt); err != nil {
		return false, err
	}
	fmt.Println("vor generierung start")
	if err = Gener(dir, filepath.Join(basis, importpre, "gener"), &prepStmt); err != nil {
		return false, err
	}
	if len(prepTypes) > 0 {
		dirSetter(basis, "generprep")
		fmt.Println("vor generierung weiter")
		if err = GenerPrep(dir, filepath.Join(basis, importpre, "generprep"), &prepTypes); err != nil {
			return false, err
		}

		return true, nil
	} else {
		return false, nil
	}

}

func PrepList(prefix string, prepTypes *map[string]prepTyp, schema string) error {

	pool := db.GetPool()
	con, err := pool.Acquire()
	if err != nil {
		return err
	}
	defer pool.Release(con)
	rows, err := pool.Query(`select viewname ,schemaname
		 from pg_catalog.pg_views
		  where viewname like any($1) and schemaname like any($2)`,
		"{"+prefix+"}", "{"+schema+"}")
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var name string
		var schema string
		if err := rows.Scan(&name, &schema); err != nil {
			return err
		}
		if err := search(con, 0, name, prepTypes, schema); err != nil {
			return errors.Wrapf(err, "suche name %s %s", name, schema)
		}
	}
	return nil
}
