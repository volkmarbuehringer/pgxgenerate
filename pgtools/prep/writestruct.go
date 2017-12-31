package prep

import (
	"fmt"
	"strings"

	"github.com/jackc/pgx"

	"pgxgenerate/pgtools/db"
	"pgxgenerate/pgtools/writer"
)

func writeStruct(name string, f *writer.Writer, flag bool, fields []pgx.FieldDescription, aname string, schema string) error {

	if w, err := f.Create(name + ".go"); err != nil {
		return err
	} else {

		defer f.Close()

		writeHeader(w, flag)

		for _, field := range fields {
			if strings.Contains(field.DataTypeName, "generprep.") {
				fmt.Fprintf(w, `import 	%q
					`, importGenerPre)
				break
			}
		}

		fmt.Fprintf(w, "const %[1]sName=%[2]q \n\n\ntype %[1]s struct{\n", strings.Title(name), name)

		for _, field := range fields {
			fmt.Println(field)
			var ftyp string
			if field.DataTypeName[0:1] == "_" {
				ftyp = db.TypCon(field.DataTypeName[1:]) + "Array"
			} else {
				ftyp = db.TypCon(field.DataTypeName)
			}

			fmt.Fprintf(w, "%s %s `json:%q`\n", strings.Title(field.Name), ftyp, field.Name)

		}
		fmt.Fprintln(w, "}")

		fmt.Fprintf(w, "\nvar %sColumns=[]string{\n", strings.Title(name))

		for _, field := range fields {

			fmt.Fprintf(w, "%q,\n", field.Name)

		}
		fmt.Fprintln(w, "}")

		if !flag {
			fmt.Fprintf(w, `type %[1]sArray []%[1]s

				func (x *%[1]sArray) Scanner() []interface{} {
					*x = append(*x, %[1]s{})

					return (*x)[len(*x)-1].Scanner()
				}

		`, strings.Title(name))

		}

		/*

		 */
		var typer string
		if flag {
			typer = "[]pgtype.BinaryDecoder"
		} else {
			typer = "[]interface {}"
		}
		fmt.Fprintf(w, "func (x *%[1]s) Scanner() %[2]s  {\n return %[2]s {\n", strings.Title(name), typer)

		for _, field := range fields {

			fmt.Fprintf(w, "&x.%s,\n", strings.Title(field.Name))
		}
		fmt.Fprintf(w, "}\n}\n")

		if flag {
			writeMethoden(w, name)

			writeInit(w, name, aname, schema)

		}

		if err := w.Flush(); err != nil {
			return err
		}
	}
	return nil
}
