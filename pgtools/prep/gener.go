package prep

import (
	"fmt"

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

func Gener(dirname string, prepStmt *map[string]*pgx.PreparedStatement) error {
	fmt.Println("starte generierung phase1")
	f := writer.Init(dirname)
	if w, err := f.Create("initfunc1.go"); err != nil {
		return err
	} else {

		defer f.Close()

		fmt.Fprintf(w, `package gener

		import %q

			`, importDB)

		fmt.Fprintf(w, "func init(){\n\n")

		for k, stmt := range *prepStmt {

			if err := writeStruct(k, writer.Init(dirname), false, stmt.FieldDescriptions, "", ""); err != nil {
				return errors.Wrapf(err, "fehler gen bei %s", k)
			}
			writeInit1(w, k, stmt.SQL)
		}
		fmt.Fprintf(w, "}\n\n")
		if err := w.Flush(); err != nil {
			return err
		}
	}
	fmt.Println("ende generierung phase1")
	return nil
}
