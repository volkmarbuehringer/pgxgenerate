package generprep
			import "pgxgenerate/pgtools/db"
			import (
	"fmt"
	"github.com/jackc/pgx/pgtype"
  "github.com/jackc/pgx"
)
		
			const Agg_steuerungseinheitName="agg_steuerungseinheit" 


type Agg_steuerungseinheit struct{
Str_id db.Int4 `json:"str_id"`
Str_opcid db.Int4 `json:"str_opcid"`
Str_plantnr db.Int4 `json:"str_plantnr"`
Str_typstr db.Varchar `json:"str_typstr"`
Str_sernr db.Varchar `json:"str_sernr"`
Str_crdate db.Timestamptz `json:"str_crdate"`
Str_upddate db.Timestamptz `json:"str_upddate"`
Str_aktiv db.Bool `json:"str_aktiv"`
}

var Agg_steuerungseinheitColumns=[]string{
"str_id",
"str_opcid",
"str_plantnr",
"str_typstr",
"str_sernr",
"str_crdate",
"str_upddate",
"str_aktiv",
}
func (x *Agg_steuerungseinheit) Scanner() []pgtype.BinaryDecoder  {
 return []pgtype.BinaryDecoder {
&x.Str_id,
&x.Str_opcid,
&x.Str_plantnr,
&x.Str_typstr,
&x.Str_sernr,
&x.Str_crdate,
&x.Str_upddate,
&x.Str_aktiv,
}
}

func (src *Agg_steuerungseinheit) AssignTo(dst interface{}) error {
	if src != nil {
		ttt, ok := dst.(*Agg_steuerungseinheit)
		if !ok {
				return fmt.Errorf("cannot assig Agg_steuerungseinheit ")
		}
		*ttt = *src
	}
	return nil
}

func (dst *Agg_steuerungseinheit) Set(src interface{}) error {
	return fmt.Errorf("cannot convert to Agg_steuerungseinheit")
}

func (dst *Agg_steuerungseinheit) Get() interface{} {
return dst
}

func (dst *Agg_steuerungseinheit) DecodeBinary(ci *pgtype.ConnInfo, src []byte) error {
	if src == nil {
			return nil
	}

	struT := new(Agg_steuerungseinheit)
	if	err := db.DecodeBinary(ci,struT.Scanner(),src);err != nil {
		return err
	}
		*dst = *struT

	return nil
}


		func init(){

db.InitOIDMap["Agg_steuerungseinheit"]=func(con *pgx.Conn){
db.Register(con , &Agg_steuerungseinheit{},"Agg_steuerungseinheit", "agg_steuerungseinheit" ,"eeacollector")
}
	}

	