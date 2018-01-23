package db

import (
	"fmt"
	"os"
	"strings"

	"github.com/jackc/pgx"
	"github.com/jackc/pgx/pgtype"
	"github.com/pkg/errors"
)

var typsql string

var InitOIDMap = map[string]func(con *pgx.Conn){}

type SQLInterB interface {
	Name() string
	Columns() []string
	SQL() string
}
type SQLInter interface {
	SQLInterB
	Scanner() []interface{}
}

var SQLListe []SQLInter

var SQLListeAgg []SQLInter

//var InitScanMap = map[string]func() ([]interface{}, interface{}, []string){}
//var PreqSQLMap = map[string]string{}
//var CheckerCalls = []func(*pgx.Conn) error{}

func Register(con *pgx.Conn, stru pgtype.Value, name string, typname string, schema string) {

	var oid pgtype.OID

	if err := con.QueryRow(`select t.oid from pg_namespace n
		inner join pg_type t on t.typnamespace= n.oid
		where n.nspname =any($1) and t.typname =$2`, "{"+typsql+"}", typname).Scan(&oid); err != nil {
		panic(errors.Wrapf(err, "aggregat in schema %s nicht gefunden: %s", typsql, typname))
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

	//fmt.Println("lade oid:", len(InitOIDMap))
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
	//fmt.Println("prepare sql:", len(InitScanMap))
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

	//fmt.Println("prüfen aggregate:", len(CheckerCalls))

	for _, x := range SQLListeAgg {
		sql := x.SQL()

		sql = sql[strings.Index(sql, `"`)+1:]
		pos1 := strings.Index(sql, `"`)
		sql1 := sql[pos1+3:]

		pos3 := strings.Index(sql1, `"`)
		if err := Checkaggview(con, sql1[:pos3], sql[:pos1], x.Columns(), x.Scanner()); err != nil {
			return err
		}
	}

	//fmt.Println("prüfen sql:", len(PreqSQLMap))
	for _, x := range SQLListe {
		sql := x.SQL()
		if len(sql) > 0 {
			if stmt, err := pool.Prepare(x.Name(), sql); err != nil {

				return err

			} else {

				if strings.Contains(x.Name(), "Return") {

					if err := checker(x.Name(), stmt.FieldDescriptions, x.Scanner(), x.Columns()); err != nil {
						return err
					}
					if err := pool.Deallocate(x.Name()); err != nil {
						return err
					}
				} else {

					erg := CheckSQL(x.SQL())
					var fd []pgx.FieldDescription
					if len(erg) > 0 {

						//	fmt.Println("testupd", x, len(erg), len(stmt.ParameterOIDs))
						var err error
						if fd, err = CheckOIDs(con, erg, stmt.ParameterOIDs); err != nil {
							return err
						}
						//	fmt.Println(fd)

					} else {
						fd = stmt.FieldDescriptions

					}
					if err := checker(x.Name(), fd, x.Scanner(), x.Columns()); err != nil {
						return err
					}
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
