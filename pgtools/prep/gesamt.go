package prep

import (
	"fmt"
	"path/filepath"

	"github.com/jackc/pgx"
	"github.com/pkg/errors"

	"pgxgenerate/pgtools/db"
)

var importDB string
var importGenerPre string

func dirSetter(importpre string, basis string) {
	importDB = filepath.Join(basis, "pgtools/db")
	importGenerPre = filepath.Join(importpre, "generprep")

}

func GesamtPrep(prefix string, dir string, importpre string, basis string, schema string) error {
	prepTypes := make(map[string]prepTyp)
	dirSetter(importpre, basis)
	if err := PrepList(prefix, &prepTypes, schema); err != nil {
		return errors.Wrapf(err, "Views nicht gefunden")
	}
	fmt.Println("vor generierung weiter", len(prepTypes))
	if err := GenerPrep(filepath.Join(dir, "generprep"), &prepTypes); err != nil {
		return errors.Wrapf(err, "generierungsfehler")
	}
	return nil
}

func Gesamt(dir string, importpre string, basis string) (bool, error) {
	fmt.Println("vor generierung prep")
	sql, err := ReadSQL(filepath.Join(dir, "sqls"))
	if err != nil {
		return false, err
	}
	prepTypes := make(map[string]prepTyp)
	prepStmt := make(map[string]*pgx.PreparedStatement)
	dirSetter(importpre, basis)
	if err := Prep(sql, &prepTypes, &prepStmt); err != nil {
		return false, err
	}
	fmt.Println("vor generierung start")
	if err = Gener(filepath.Join(dir, "gener"), &prepStmt); err != nil {
		return false, err
	}

	fmt.Println("vor generierung weiter")
	if err = GenerPrep(filepath.Join(dir, "generprep"), &prepTypes); err != nil {
		return false, err
	}
	if len(prepTypes) > 0 {
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
