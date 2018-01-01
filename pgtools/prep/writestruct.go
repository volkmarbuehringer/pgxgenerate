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

		writeInit(w, name+"Array", bname, schema)

		if err := w.Flush(); err != nil {
			return err
		}
	}
	return nil
}

func writeStruct(name string, f *writer.Writer, flag bool, fields []pgx.FieldDescription, aname string, schema string, ext string) error {

	if w, err := f.Create(name + ".go"); err != nil {
		return err
	} else {

		defer f.Close()

		writeHeader(w, ext)

		fmt.Fprintf(w, "const %[1]sName=%[2]q \n\n\ntype %[1]s struct{\n", strings.Title(name), name)

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

		fmt.Fprintf(w, "\nvar %sColumns=[]string{\n", strings.Title(name))

		for _, field := range fields {

			fmt.Fprintf(w, "%q,\n", field.Name)

		}
		fmt.Fprintln(w, "}")

		var helper = func() string {
			if flag {
				return "[]pgtype.BinaryDecoder"
			} else {
				fmt.Fprintf(w, pgxscanner, strings.Title(name))
				return "[]interface {}"
			}
		}

		fmt.Fprintf(w, "func (x *%[1]s) Scanner() %[2]s  {\n return %[2]s {\n", strings.Title(name), helper())

		for _, field := range fields {

			fmt.Fprintf(w, "&x.%s,\n", strings.Title(field.Name))
		}
		fmt.Fprintf(w, "}\n}\n")

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
