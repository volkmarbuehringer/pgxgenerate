package generprep
			import 	"fmt"
import 	"github.com/jackc/pgx/pgtype"
import 	"github.com/jackc/pgx"
 
import "prounix.de/pgtools/db"
const Agg_eeaopcserverName="agg_eeaopcserver" 


type Agg_eeaopcserver struct{
Eea_opcid db.Int4 `json:"eea_opcid"`
Eea_sernr db.Varchar `json:"eea_sernr"`
Eea_plantnr db.Int4 `json:"eea_plantnr"`
Eea_typnr db.Varchar `json:"eea_typnr"`
Eea_typstr db.Varchar `json:"eea_typstr"`
Eea_hersteller db.Varchar `json:"eea_hersteller"`
Eea_nennleistung db.Numeric `json:"eea_nennleistung"`
Eea_anzminuten db.Int4 `json:"eea_anzminuten"`
Eea_aktiv db.Bool `json:"eea_aktiv"`
}

var Agg_eeaopcserverColumns=[]string{
"eea_opcid",
"eea_sernr",
"eea_plantnr",
"eea_typnr",
"eea_typstr",
"eea_hersteller",
"eea_nennleistung",
"eea_anzminuten",
"eea_aktiv",
}
func (x *Agg_eeaopcserver) Scanner() []pgtype.BinaryDecoder  {
 return []pgtype.BinaryDecoder {
&x.Eea_opcid,
&x.Eea_sernr,
&x.Eea_plantnr,
&x.Eea_typnr,
&x.Eea_typstr,
&x.Eea_hersteller,
&x.Eea_nennleistung,
&x.Eea_anzminuten,
&x.Eea_aktiv,
}
}

	func (src *Agg_eeaopcserver) AssignTo(dst interface{}) error {
		if src != nil {
			ttt, ok := dst.(*Agg_eeaopcserver)
			if !ok {
					return fmt.Errorf("cannot assig Agg_eeaopcserver ")
			}
			*ttt = *src
		}
		return nil
	}

	func (dst *Agg_eeaopcserver) Set(src interface{}) error {
		return fmt.Errorf("cannot convert to Agg_eeaopcserver")
	}

	func (dst *Agg_eeaopcserver) Get() interface{} {
	return dst
	}

	func (dst *Agg_eeaopcserver) DecodeBinary(ci *pgtype.ConnInfo, src []byte) error {
		if src == nil {
				return nil
		}

		struT := new(Agg_eeaopcserver)
		if	err := db.DecodeBinary(ci,struT.Scanner(),src);err != nil {
			return err
		}
			*dst = *struT

		return nil
	}

	
		func init(){

db.InitOIDMap["Agg_eeaopcserver"]=func(con *pgx.Conn){
db.Register(con , &Agg_eeaopcserver{},"Agg_eeaopcserver", "agg_eeaopcserver" ,"eeacollector")
}
	}

	