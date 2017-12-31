package gener
import "prounix.de/pgtools/db"
const TestinsrName="testinsr" 


type Testinsr struct{
Opc_name db.Varchar `json:"opc_name"`
Opc_uri db.Varchar `json:"opc_uri"`
Opc_scadanr db.Int4 `json:"opc_scadanr"`
}

var TestinsrColumns=[]string{
"opc_name",
"opc_uri",
"opc_scadanr",
}
type TestinsrArray []Testinsr

				func (x *TestinsrArray) Scanner() []interface{} {
					*x = append(*x, Testinsr{})

					return (*x)[len(*x)-1].Scanner()
				}

		func (x *Testinsr) Scanner() []interface {}  {
 return []interface {} {
&x.Opc_name,
&x.Opc_uri,
&x.Opc_scadanr,
}
}
