package prep

import (
	"fmt"
	"path/filepath"
	"strings"

	"pgxgenerate/pgtools/writer"

	"github.com/jackc/pgx"
	"github.com/jackc/pgx/pgtype"
	"github.com/pkg/errors"
)

type prepTyp struct {
	aname  string
	bname  string
	schema string
	aoid   pgtype.OID
	boid   pgtype.OID
}

func Gener(dir string, importer string, prepStmt *map[string]*pgx.PreparedStatement) error {
	var dirname = filepath.Join(dir, importer)
	fmt.Println("starte generierung phase1")
	f := writer.Init(dirname)
	if w, err := f.Create("initfunc1.go"); err != nil {
		return err
	} else {

		defer f.Close()

		writeHeader(w, "")

		fmt.Fprintf(w, "func init(){\n\ndb.SQLListe=[]db.SQLInter{")

		for k, stmt := range *prepStmt {

			var x string
			for _, field := range stmt.FieldDescriptions {
				if strings.Contains(field.DataTypeName, "generprep.") {
					x = fmt.Sprintf(`import 	%q
								`, importer+"prep")
					break
				}
			}
			if err := writeStruct(k, writer.Init(dirname), false, stmt.FieldDescriptions, "", "", x, stmt.SQL); err != nil {
				return errors.Wrapf(err, "fehler gen bei %s", k)
			}
			writeInit1(w, k)
		}
		fmt.Fprintf(w, "}}\n\n")
		if err := w.Flush(); err != nil {
			return err
		}
	}
	fmt.Println("ende generierung phase1")
	return nil
}
