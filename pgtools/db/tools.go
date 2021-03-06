package db

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/jackc/pgx"
	"github.com/jackc/pgx/pgtype"
)

var r1 *regexp.Regexp
var r2 *regexp.Regexp
var r3 *regexp.Regexp
var r4 *regexp.Regexp
var r5 *regexp.Regexp
var r6 *regexp.Regexp
var r7 *regexp.Regexp
var r8 *regexp.Regexp
var r9 *regexp.Regexp

func init() {
	const columnname = `([a-z]{3}_[a-z,0-9]+)`

	r1 = regexp.MustCompile(`(?is)insert\s+into\s+(.+)\s*\((.+)\)\s*values\s*\((.+)\)`)
	r2 = regexp.MustCompile(`(?is)update\s+(.+)\s+set\s+(.+)\s+where\s+(.+)`)

	r3 = regexp.MustCompile(`(?is)(\s+returning\s+.+)`)

	r4 = regexp.MustCompile(`(?is)\s*` + columnname + `\s*=.*\$(\d{1,2}).*`)
	r5 = regexp.MustCompile(`(?is).*\$(\d{1,2}).*`)
	r6 = regexp.MustCompile(`(?is)(\(|\)|\+|\-|\*|\/|,|substr)`)
	r7 = regexp.MustCompile(`(?is)(\s+and\s+|\s+or\s+)`)
	r8 = regexp.MustCompile(`(?is)\s*` + columnname + `\s*(?:=|<|>|<>|\s+in\s+|<=|>=|\slike).*\$(\d{1,2}).*`)
	r9 = regexp.MustCompile(`(?is)\s*\$(\d{1,2}).*(?:=|<|>|<>|\s+in\s+|<=|>=|\slike)\s*` + columnname)
}

func splitter(input string) []string {
	erg := make([]string, 0)
	var brazahl int
	var striflag bool
	var pos int
	for idx, w := range strings.Split(input, "") {
		switch w {
		case "'":
			striflag = !striflag
		case "(":
			brazahl++
		case ")":
			brazahl--
		case ",":
			if brazahl == 0 && !striflag {
				erg = append(erg, input[pos:idx])
				pos = idx + 1
			}
		}

	}
	if pos < len(input) {
		erg = append(erg, input[pos:])
	}

	return erg
}

func CheckOIDs(con *pgx.Conn, namen []string, binder []pgtype.OID) ([]pgx.FieldDescription, error) {
	if len(binder) != len(namen) {
		return nil, fmt.Errorf("interner fehler namen: %d binder: %d", len(namen), len(binder))
	}

	liste := make([]pgx.FieldDescription, len(binder))
	for idx, n := range binder {
		if v, ok := con.ConnInfo.DataTypeForOID(n); !ok {
			return nil, fmt.Errorf("oid nicht gefunden %d %d", n, idx)
		} else {
			if len(namen[idx]) == 0 {
				namen[idx] = fmt.Sprintf("pos%d", idx)
			}
			liste[idx].Name = namen[idx]
			liste[idx].DataType = n
			liste[idx].DataTypeName = v.Name
		}
	}
	return liste, nil
}

func Checkaggview(con *pgx.Conn, table string, schema string, columns []string, intertypes []interface{}) error {
	if rows, err := con.Query("xxxaggviewxxx", schema+"."+table); err != nil {
		return err
	} else {
		var pos int
		defer rows.Close()
		for rows.Next() {
			var name string
			var typ string

			if err := rows.Scan(&name, &typ); err != nil {
				return err
			}
			if len(columns) <= pos {
				return fmt.Errorf("falsche anzahl spalten%s:%s %d", schema, table, pos)
			}
			if columns[pos] != name {
				return fmt.Errorf("falscher spaltename %s %s %s:%s %d", table, schema, name, columns[pos], pos)
			}
			typnamei := strings.Replace(fmt.Sprintf("%T", intertypes[pos]), "*", "", 1)
			if typnamei != "db."+strings.Title(typ) {
				return fmt.Errorf("falscher spaltentyp %s %s %s<> %s", table, columns[pos], typnamei, typ)
			}
			pos++
		}
		if len(columns) != pos {
			return fmt.Errorf("falsche anzahl spalten%s:%s %d", schema, table, pos)
		}
	}
	return nil
}
