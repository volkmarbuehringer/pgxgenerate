package gener
			import "pgxgenerate/pgtools/db"
			
			const Test0Name="test0" 


type Test0 struct{
Opc_tcptim db.Int4 `json:"opc_tcptim"`
Opc_preverror db.Text `json:"opc_preverror"`
Opc_name db.Text `json:"opc_name"`
Opc_id db.Int4 `json:"opc_id"`
}

var Test0Columns=[]string{
"opc_tcptim",
"opc_preverror",
"opc_name",
"opc_id",
}
type Test0Array []Test0

				func (x *Test0Array) Scanner() []interface{} {
					*x = append(*x, Test0{})

					return (*x)[len(*x)-1].Scanner()
				}

		func (x *Test0) Scanner() []interface {}  {
 return []interface {} {
&x.Opc_tcptim,
&x.Opc_preverror,
&x.Opc_name,
&x.Opc_id,
}
}
