package gener
import "pgxgenerate/pgtools/db"
		const TestinsName="testins" 


type Testins struct{
Opc_name db.Varchar `json:"opc_name"`
Opc_uri db.Varchar `json:"opc_uri"`
Opc_scadanr db.Int4 `json:"opc_scadanr"`
}

var TestinsColumns=[]string{
"opc_name",
"opc_uri",
"opc_scadanr",
}
type TestinsArray []Testins

				func (x *TestinsArray) Scanner() []interface{} {
					*x = append(*x, Testins{})

					return (*x)[len(*x)-1].Scanner()
				}

		func (x *Testins) Scanner() []interface {}  {
 return []interface {} {
&x.Opc_name,
&x.Opc_uri,
&x.Opc_scadanr,
}
}
