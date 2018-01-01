package generprep
			import "pgxgenerate/pgtools/db"
			import 	"fmt"
			import 	"github.com/jackc/pgx/pgtype"
			import 	"github.com/jackc/pgx"
						
			const Agg_tplantsName="agg_tplants" 


type Agg_tplants struct{
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

var Agg_tplantsColumns=[]string{
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
func (x *Agg_tplants) Scanner() []pgtype.BinaryDecoder  {
 return []pgtype.BinaryDecoder {
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

	func (src *Agg_tplants) AssignTo(dst interface{}) error {
		if src != nil {
			ttt, ok := dst.(*Agg_tplants)
			if !ok {
					return fmt.Errorf("cannot assig Agg_tplants ")
			}
			*ttt = *src
		}
		return nil
	}

	func (dst *Agg_tplants) Set(src interface{}) error {
		return fmt.Errorf("cannot convert to Agg_tplants")
	}

	func (dst *Agg_tplants) Get() interface{} {
	return dst
	}

	func (dst *Agg_tplants) DecodeBinary(ci *pgtype.ConnInfo, src []byte) error {
		if src == nil {
				return nil
		}

		struT := new(Agg_tplants)
		if	err := db.DecodeBinary(ci,struT.Scanner(),src);err != nil {
			return err
		}
			*dst = *struT

		return nil
	}

	
		func init(){

db.InitOIDMap["Agg_tplants"]=func(con *pgx.Conn){
db.Register(con , &Agg_tplants{},"Agg_tplants", "agg_tplants" ,"sdbms")
}
	}

	