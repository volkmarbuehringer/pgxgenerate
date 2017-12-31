package prep

import (
	"fmt"

	"github.com/pkg/errors"
	"prounix.de/pgtools/db"
	"prounix.de/pgtools/writer"
)

func GenerPrep(dirname string, importpre string, prepTypes *map[string]prepTyp) error {
	fmt.Println("starte generierung phase2")
	con, err := db.GetPool().Acquire()
	if err != nil {
		return err
	}

	if len(*prepTypes) == 0 {
		f := writer.Init(dirname)

		if w, err := f.Create("dummy.go"); err != nil {
			return err
		} else {

			fmt.Fprintf(w, `package generprep

			type Dummy int

				`)
			w.Flush()
			f.Close()

		}

	}

	fmt.Println("vor generinit", len(*prepTypes))
	for k, t := range *prepTypes {

		if stmt, err := con.Prepare(k, fmt.Sprintf("select * from %q.%q", t.schema, k)); err != nil {

			return fmt.Errorf("fehler bei prepare %s %s", t.schema, k)
		} else {
			if err := writeStruct(k, writer.Init(dirname), true, stmt.FieldDescriptions, t.aname, importpre, t.schema); err != nil {

				return errors.Wrapf(err, "fehler gen bei %s", k)
			}
			if err := writeArrayCode(k, writer.Init(dirname), t.bname, t.schema); err != nil {
				return errors.Wrapf(err, "fehler gen bei %s", k)
			}
			fmt.Println("schreibe init", k, t.aname, t.bname)

		}

	}
	fmt.Println("ende generierung phase2")
	return nil
}
