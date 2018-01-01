package gener
			import "pgxgenerate/pgtools/db"
			
			const InsplantName="insplant" 


type Insplant struct{
Fuplantid db.Int4 `json:"fuplantid"`
Fulocno db.Int4 `json:"fulocno"`
Fuplantno db.Int2 `json:"fuplantno"`
Fuplantcode db.Int2 `json:"fuplantcode"`
Fuspecialid db.Int2 `json:"fuspecialid"`
Fuserieno db.Int4 `json:"fuserieno"`
Fdcomdate db.Date `json:"fdcomdate"`
Fdturnoff db.Date `json:"fdturnoff"`
Fcplantalias db.Varchar `json:"fcplantalias"`
Fcplanthwtype db.Varchar `json:"fcplanthwtype"`
Fuplantpower db.Int4 `json:"fuplantpower"`
Foincurrentlist db.Int2 `json:"foincurrentlist"`
Foserialchanged db.Int2 `json:"foserialchanged"`
}

var InsplantColumns=[]string{
"fuplantid",
"fulocno",
"fuplantno",
"fuplantcode",
"fuspecialid",
"fuserieno",
"fdcomdate",
"fdturnoff",
"fcplantalias",
"fcplanthwtype",
"fuplantpower",
"foincurrentlist",
"foserialchanged",
}
type InsplantArray []Insplant

				func (x *InsplantArray) Scanner() []interface{} {
					*x = append(*x, Insplant{})

					return (*x)[len(*x)-1].Scanner()
				}

		func (x *Insplant) Scanner() []interface {}  {
 return []interface {} {
&x.Fuplantid,
&x.Fulocno,
&x.Fuplantno,
&x.Fuplantcode,
&x.Fuspecialid,
&x.Fuserieno,
&x.Fdcomdate,
&x.Fdturnoff,
&x.Fcplantalias,
&x.Fcplanthwtype,
&x.Fuplantpower,
&x.Foincurrentlist,
&x.Foserialchanged,
}
}
