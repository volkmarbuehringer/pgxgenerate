package generprep
			import 	"fmt"
import 	"github.com/jackc/pgx/pgtype"
import 	"github.com/jackc/pgx"
 
import "prounix.de/pgtools/db"
type Agg_eeaopcserverArray []Agg_eeaopcserver

		func (src *Agg_eeaopcserverArray) AssignTo(dst interface{}) error {

			if src != nil {
				ttt, ok := dst.(*Agg_eeaopcserverArray)
				if !ok {
						return fmt.Errorf("cannot assign Agg_eeaopcserver")
				}
				*ttt = *src
			}
			return nil
		}
		func (dst *Agg_eeaopcserverArray) Set(src interface{}) error {
			return fmt.Errorf("cannot convert to Agg_eeaopcserverArray")
		}

		func (dst *Agg_eeaopcserverArray) Get() interface{} {
			return dst
		}

						func (dst *Agg_eeaopcserverArray) DecodeBinary(ci *pgtype.ConnInfo, src []byte) error {

							elements := make(Agg_eeaopcserverArray, 0)
							if src == nil {
									*dst = elements
								return nil
							}
							funcer := func (result *Agg_eeaopcserverArray) func()[]pgtype.BinaryDecoder {
								return func() []pgtype.BinaryDecoder{
									pos := len(*result)
									*result = append(*result,Agg_eeaopcserver{})
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

		
		func init(){

db.InitOIDMap["Agg_eeaopcserverArray"]=func(con *pgx.Conn){
db.Register(con , &Agg_eeaopcserverArray{},"Agg_eeaopcserverArray","_agg_eeaopcserver","eeacollector" )
}


	}

	