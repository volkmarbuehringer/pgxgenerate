package prep

import (
	"fmt"
	"io"
	"strings"
)

func writeHeader(w io.Writer, flag bool) {
	if flag {
		fmt.Fprintln(w,
			`package generprep
			import 	"fmt"
import 	"github.com/jackc/pgx/pgtype"
import 	"github.com/jackc/pgx"
 `)
	} else {
		fmt.Fprintf(w, "package gener\n")
	}
	fmt.Fprintln(w, `import "prounix.de/pgtools/db"`)
}

func writeMethoden(w io.Writer, name string) {

	fmt.Fprintf(w, `
	func (src *%[1]s) AssignTo(dst interface{}) error {
		if src != nil {
			ttt, ok := dst.(*%[1]s)
			if !ok {
					return fmt.Errorf("cannot assig %[1]s ")
			}
			*ttt = *src
		}
		return nil
	}

	func (dst *%[1]s) Set(src interface{}) error {
		return fmt.Errorf("cannot convert to %[1]s")
	}

	func (dst *%[1]s) Get() interface{} {
	return dst
	}

	func (dst *%[1]s) DecodeBinary(ci *pgtype.ConnInfo, src []byte) error {
		if src == nil {
				return nil
		}

		struT := new(%[1]s)
		if	err := db.DecodeBinary(ci,struT.Scanner(),src);err != nil {
			return err
		}
			*dst = *struT

		return nil
	}

	`, strings.Title(name))

}
