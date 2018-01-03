package prep

import (
	"fmt"
	"io"
	"strings"

	"github.com/jackc/pgx"
)

func writeInit1(w io.Writer, name string, prepSql string) {
	fmt.Fprintf(w, "db.InitScanMap[%[1]sName]= func ()( []interface {}, interface{},[]string ) 	{\n var x %[1]s\nreturn x.Scanner(), &x,%[1]sColumns\n", strings.Title(name))

	fmt.Fprintf(w, "\n	}\n")

	if len(prepSql) > 0 {
		fmt.Fprintf(w, "db.PreqSQLMap[%[1]sName]= `%[2]s`\n\n", strings.Title(name), prepSql)
	}

}

func writeInit(w io.Writer, name, aname string, schema string, fields []pgx.FieldDescription) {

	var zus string

	if len(fields) > 0 {
		helper := make([]string, 0, len(fields))
		for _, field := range fields {
			helper = append(helper, fmt.Sprintf("&x.%s", strings.Title(field.Name)))
		}
		zus = `
	   db.CheckerCalls=append(db.CheckerCalls,func(con *pgx.Conn)error{
	   	var x %[1]s
	   return db.Checkaggview(con , %[2]q, %[3]q ,%[1]sColumns, []interface {} {` +
			strings.Join(helper, ",\n") + `})
	   })`
	}

	fmt.Fprintf(w, `
		func init(){

db.InitOIDMap[%[1]q]=func(con *pgx.Conn){
db.Register(con , &%[1]s{},%[1]q, %[2]q ,%[3]q)
}`+zus+`

}	`, strings.Title(name), aname, schema)

}
