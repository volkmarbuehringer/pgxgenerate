package prep

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/jackc/pgx"
	"github.com/jackc/pgx/pgtype"
	"github.com/pkg/errors"

	"prounix.de/pgtools/db"
)

func Prep(prepSql map[string]string, prepTypes *map[string]prepTyp, prepStmt *map[string]*pgx.PreparedStatement) error {

	con, err := db.GetPool().Acquire()
	if err != nil {
		return err
	}

	r := regexp.MustCompile(`unknown oid:\s+([0-9]{1,5})`)

	for k, x := range prepSql {
		fmt.Println("sqlprep", x)

		if db.CheckSQLReturn(x) {
			if stmt, err := con.Prepare(k+"Return", x); err != nil {

				return errors.Wrapf(err, "prepare")

			} else {
				stmt.SQL = ""
				(*prepStmt)[k+"Return"] = stmt
			}

		}
		erg := db.CheckSQL(x)
		if stmt, err := con.Prepare(k, x); err != nil {

			errs := err.Error()

			if x := r.FindStringSubmatch(errs); len(x) == 2 {

				if oid, err := strconv.Atoi(x[1]); err != nil {
					return err
				} else if err = search(con, pgtype.OID(oid), "", prepTypes, ""); err != nil {
					return errors.Wrapf(err, "oid nicht gefunden %d %s", oid, k)
				}
			} else {

				return errors.Wrapf(err, "prepare %d %s", len(x), k)
			}

		} else if len(erg) > 0 {

			if fd, err := db.CheckOIDs(con, erg, stmt.ParameterOIDs); err != nil {
				return err
			} else {
				stmt.FieldDescriptions = fd
			}

			(*prepStmt)[k] = stmt
		} else {
			(*prepStmt)[k] = stmt
		}
	}

	return nil
}
