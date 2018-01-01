package prep

const importheader = `import (
	"fmt"
	"github.com/jackc/pgx/pgtype"
  "github.com/jackc/pgx"
)
		`

const pgxscanner = `type %[1]sArray []%[1]s

	func (x *%[1]sArray) Scanner() []interface{} {
		*x = append(*x, %[1]s{})

		return (*x)[len(*x)-1].Scanner()
	}

`

const pgxarraymethoden = `type %[1]sArray []%[1]s

func (src *%[1]sArray) AssignTo(dst interface{}) error {

	if src != nil {
		ttt, ok := dst.(*%[1]sArray)
		if !ok {
				return fmt.Errorf("cannot assign %[1]s")
		}
		*ttt = *src
	}
	return nil
}
func (dst *%[1]sArray) Set(src interface{}) error {
	return fmt.Errorf("cannot convert to %[1]sArray")
}

func (dst *%[1]sArray) Get() interface{} {
	return dst
}

				func (dst *%[1]sArray) DecodeBinary(ci *pgtype.ConnInfo, src []byte) error {

					elements := make(%[1]sArray, 0)
					if src == nil {
							*dst = elements
						return nil
					}
					funcer := func (result *%[1]sArray) func()[]pgtype.BinaryDecoder {
						return func() []pgtype.BinaryDecoder{
							pos := len(*result)
							*result = append(*result,%[1]s{})
							return (*result)[pos].Scanner()
						}
					}

					helperfun := funcer(&elements)
					if err := db.Helper(ci, src, helperfun);err != nil {
					 return err
				 }
					*dst = elements
					return nil
				}

`

const pgxmethoden = `
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

`
