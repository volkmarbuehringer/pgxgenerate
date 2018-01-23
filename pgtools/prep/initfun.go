package prep

import (
	"fmt"
	"io"
	"strings"
)

func writeInit1(w io.Writer, name string) {

	fmt.Fprintf(w, "new(%[1]s),\n", strings.Title(name))

}

func writeInit(w io.Writer, name, aname string, schema string) {

	fmt.Fprintf(w, `
		func init(){

db.InitOIDMap[%[1]q]=func(con *pgx.Conn){
db.Register(con , &%[1]s{},%[1]q, %[2]q ,%[3]q)
}
db.SQLListeAgg=append(db.SQLListeAgg,new(%[1]s))

}	`, strings.Title(name), aname, schema)

}

func writeInit2(w io.Writer, name, aname string, schema string) {

	fmt.Fprintf(w, `
		func init(){

db.InitOIDMap[%[1]q]=func(con *pgx.Conn){
db.Register(con , &%[1]s{},%[1]q, %[2]q ,%[3]q)
}

}	`, strings.Title(name), aname, schema)

}
