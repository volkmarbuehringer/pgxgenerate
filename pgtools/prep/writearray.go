package prep

import (
	"fmt"
	"strings"

	"prounix.de/pgtools/writer"
)

func writeArrayCode(name string, f *writer.Writer, bname string, schema string) error {

	if w, err := f.Create(name + "array.go"); err != nil {
		return err
	} else {

		defer f.Close()

		writeHeader(w, true)
		fmt.Fprintf(w, `type %[1]sArray []%[1]s

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

		`, strings.Title(name))

		writeInitArray(w, name, bname, schema)

		if err := w.Flush(); err != nil {
			return err
		}
	}
	return nil
}
