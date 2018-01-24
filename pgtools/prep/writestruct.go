package prep

import (
	"fmt"
	"strings"

	"github.com/jackc/pgx"

	"pgxgenerate/pgtools/db"
	"pgxgenerate/pgtools/writer"
)

func writeArrayCode(name string, f *writer.Writer, bname string, schema string) error {

	if w, err := f.Create(name + "array.go"); err != nil {
		return err
	} else {

		defer f.Close()

		writeHeader(w, importheader)
		fmt.Fprintf(w, pgxarraymethoden, strings.Title(name))

		writeInit2(w, name+"Array", bname, schema)

		if err := w.Flush(); err != nil {
			return err
		}
	}
	return nil
}

func writeStruct(name string, f *writer.Writer, flag bool, fields []pgx.FieldDescription, aname string, schema string, ext string, prepSql string) error {

	if w, err := f.Create(name + ".go"); err != nil {
		return err
	} else {

		defer f.Close()

		writeHeader(w, ext)

		fmt.Fprintf(w, "func (_ *%[1]s) SQL() string {\n return `%s` }\n\n", strings.Title(name), prepSql)

		fmt.Fprintf(w, "func (_ *%[1]s) Name() string {\n return %[2]q }", strings.Title(name), name)

		fmt.Fprintf(w, "\n\n\ntype %[1]s struct{\n", strings.Title(name))

		for _, field := range fields {
			fmt.Println(field)

			var helper = func() string {
				if field.DataTypeName[0:1] == "_" {
					return db.TypCon(field.DataTypeName[1:]) + "Array"
				} else {
					return db.TypCon(field.DataTypeName)
				}
			}

			fmt.Fprintf(w, "%s %s `json:%q`\n", strings.Title(field.Name), helper(), field.Name)

		}
		fmt.Fprintln(w, "}")

		fmt.Fprintf(w, "func (_ *%[1]s) Columns() []string {\n return []string{", strings.Title(name))

		for _, field := range fields {

			fmt.Fprintf(w, "%q,\n", field.Name)

		}
		fmt.Fprintln(w, "}}")

		var helper = func(namer string, typ string) {

			fmt.Fprintf(w, "func (x *%[1]s) %[3]s() %[2]s  {\n return %[2]s {\n", strings.Title(name), typ, namer)

			for _, field := range fields {

				fmt.Fprintf(w, "&x.%s,\n", strings.Title(field.Name))
			}
			fmt.Fprintf(w, "}\n}\n")
		}

		if flag {

			helper("Scanner1", "[]pgtype.BinaryDecoder")
		} else {
			fmt.Fprintf(w, pgxscanner, strings.Title(name))
		}
		helper("Scanner", "[]interface {}")
		if flag {
			fmt.Fprintf(w, pgxmethoden, strings.Title(name))

			writeInit(w, name, aname, schema)

		} else {
			fmt.Fprintf(w, "\nfunc (x *%s)String()[]string{\nreturn []string{\n", strings.Title(name))

			for _, field := range fields {
				if !strings.Contains(field.DataTypeName, "generprep.") {
					fmt.Fprintf(w, "x.%s.Stringer(),\n", strings.Title(field.Name))
				}
			}
			fmt.Fprintln(w, "\n}\n}")
		}

		if err := w.Flush(); err != nil {
			return err
		}
	}
	return nil
}
