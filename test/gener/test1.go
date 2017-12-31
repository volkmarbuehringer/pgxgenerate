package gener
import "prounix.de/pgtools/db"
const Test1Name="test1" 


type Test1 struct{
Opc_id db.Int4 `json:"opc_id"`
Opc_name db.Varchar `json:"opc_name"`
Opc_uri db.Varchar `json:"opc_uri"`
Opc_scadanr db.Int4 `json:"opc_scadanr"`
Opc_letzteanfrage db.Timestamptz `json:"opc_letzteanfrage"`
Opc_letzteresponse db.Timestamptz `json:"opc_letzteresponse"`
Opc_aktiv db.Bool `json:"opc_aktiv"`
Opc_crdate db.Timestamptz `json:"opc_crdate"`
Opc_upddate db.Timestamptz `json:"opc_upddate"`
Opc_error db.Text `json:"opc_error"`
Opc_preverror db.Text `json:"opc_preverror"`
Opc_tcptim db.Int4 `json:"opc_tcptim"`
Opc_proctim db.Int4 `json:"opc_proctim"`
Opc_conttim db.Int4 `json:"opc_conttim"`
Opc_tottim db.Int4 `json:"opc_tottim"`
Opc_letzteresponseok db.Timestamptz `json:"opc_letzteresponseok"`
Opc_runs db.Int4 `json:"opc_runs"`
Opc_timediff db.Int4 `json:"opc_timediff"`
Opc_timestamp db.Int8 `json:"opc_timestamp"`
}

var Test1Columns=[]string{
"opc_id",
"opc_name",
"opc_uri",
"opc_scadanr",
"opc_letzteanfrage",
"opc_letzteresponse",
"opc_aktiv",
"opc_crdate",
"opc_upddate",
"opc_error",
"opc_preverror",
"opc_tcptim",
"opc_proctim",
"opc_conttim",
"opc_tottim",
"opc_letzteresponseok",
"opc_runs",
"opc_timediff",
"opc_timestamp",
}
type Test1Array []Test1

				func (x *Test1Array) Scanner() []interface{} {
					*x = append(*x, Test1{})

					return (*x)[len(*x)-1].Scanner()
				}

		func (x *Test1) Scanner() []interface {}  {
 return []interface {} {
&x.Opc_id,
&x.Opc_name,
&x.Opc_uri,
&x.Opc_scadanr,
&x.Opc_letzteanfrage,
&x.Opc_letzteresponse,
&x.Opc_aktiv,
&x.Opc_crdate,
&x.Opc_upddate,
&x.Opc_error,
&x.Opc_preverror,
&x.Opc_tcptim,
&x.Opc_proctim,
&x.Opc_conttim,
&x.Opc_tottim,
&x.Opc_letzteresponseok,
&x.Opc_runs,
&x.Opc_timediff,
&x.Opc_timestamp,
}
}
