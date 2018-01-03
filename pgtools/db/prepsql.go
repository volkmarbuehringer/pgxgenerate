package db

import (
	"fmt"
	"os"
	"strings"

	"github.com/jackc/pgx"
	"github.com/jackc/pgx/pgtype"
)

var typsql string

var InitOIDMap = map[string]func(con *pgx.Conn){}

var InitScanMap = map[string]func() ([]interface{}, interface{}, []string){}
var PreqSQLMap = map[string]string{}
var CheckerCalls = []func(*pgx.Conn) error{}

func Register(con *pgx.Conn, stru pgtype.Value, name string, typname string, schema string) {

	var oid pgtype.OID

	if err := con.QueryRow(`select t.oid from pg_namespace n
		inner join pg_type t on t.typnamespace= n.oid
		where n.nspname =any($1) and t.typname =$2`, "{"+typsql+"}", typname).Scan(&oid); err != nil {
		fmt.Println("oids", err, typname, typsql)
		panic(err)
	}
	con.ConnInfo.RegisterDataType(pgtype.DataType{
		Value: stru,
		Name:  "generprep." + name,
		OID:   oid,
	})

	/*
		if x, ok := con.ConnInfo.DataTypeForOID(oid); !ok {
			fmt.Println("nicht da", typname)

		} else {

			fmt.Println("schon da", typname, x)
		}
	*/
}

func SetTyp(con *pgx.Conn) error {
	_, err := con.Exec("set search_path to " + typsql)
	if err != nil {
		panic(err)
	}

	fmt.Println("lade oid:", len(InitOIDMap))
	for _, f := range InitOIDMap {
		f(con)
	}
	return nil
}

func TypCon(ftyp string) string {

	switch ftyp {
	case "time":
		ftyp = "Text"
	case "oid":
		ftyp = "OID"
	}
	if !strings.Contains(ftyp, "generprep.") {
		ftyp = "db." + strings.Title(ftyp)
	}
	return ftyp
}

func checker(k string, fd []pgx.FieldDescription, inter []interface{}, cols []string) error {
	if len(fd) != len(inter) {
		return fmt.Errorf("falsche anzahl felder %s %d %d", k, len(inter), len(fd))
	}

	for i, d := range inter {

		typ := strings.Replace(fmt.Sprintf("%T", d), "*", "", 1)
		typ = strings.Replace(typ, "pgtype.", "db.", 1)
		vtyp := TypCon(fd[i].DataTypeName)
		if typ != vtyp {
			return fmt.Errorf("struct %s und sql stimmen nicht überein bei Feld %s %d %s: Typ %s %s", cols[i], k, i, fd[i].Name, typ, vtyp)

		}
		if cols[i] != fd[i].Name {
			return fmt.Errorf("struct %s und sql stimmen nicht überein bei Feldname %q %d %q", k, cols[i], i, fd[i].Name)

		}
	}
	return nil
}

func Prep() error {
	fmt.Println("prepare sql:", len(InitScanMap))
	pool := GetPool()
	con, err := pool.Acquire()
	if err != nil {
		panic(err)
	}
	defer con.Close()

	if _, err := con.Prepare("xxxaggviewxxx", `select attname, typname
	from pg_attribute
	inner join pg_type t on oid = atttypid
	WHERE  attrelid = $1::regclass  -- table name, optionally schema-qualified
	AND    attnum > 0
	AND    NOT attisdropped
	ORDER  BY attnum`); err != nil {
		return err
	}

	fmt.Println("prüfen aggregate:", len(CheckerCalls))
	for _, x := range CheckerCalls {
		if err := x(con); err != nil {
			return err
		}
	}
	fmt.Println("prüfen sql:", len(PreqSQLMap))
	for k, x := range PreqSQLMap {

		if stmt, err := pool.Prepare(k, x); err != nil {

			fmt.Println(err)
			return err

		} else {

			if ga, ok := InitScanMap[k]; !ok {
				return fmt.Errorf("falsches sql %s", k)
			} else {

				inter, _, cols := ga()
				//		sql := CheckSQLReturn(x)
				erg := CheckSQL(x)
				var fd []pgx.FieldDescription
				if len(erg) > 0 {
					if len(stmt.FieldDescriptions) > 0 {
						var kk = k + "Return"
						if ga, ok := InitScanMap[kk]; !ok {
							return fmt.Errorf("falsches sql %s", kk)
						} else {

							inter, _, cols := ga()
							if err := checker(kk, stmt.FieldDescriptions, inter, cols); err != nil {
								return err
							}
						}
					}
					//	fmt.Println("testupd", x, len(erg), len(stmt.ParameterOIDs))
					var err error
					if fd, err = CheckOIDs(con, erg, stmt.ParameterOIDs); err != nil {
						return err
					}
					//	fmt.Println(fd)

				} else {
					fd = stmt.FieldDescriptions

				}
				if err := checker(k, fd, inter, cols); err != nil {
					return err
				}

			}

		}
	}

	return nil
}

func Schemas() string {

	ttt, ok := os.LookupEnv("PU_SCHEMA")
	var schemaname []string

	if !ok {
		schemaname = append(schemaname, "public")
	} else {
		t := strings.Split(ttt, ",")
		schemaname = append(schemaname, t...)

	}
	return strings.Join(schemaname, ",")

}

func init() {

	typsql = Schemas()

}
