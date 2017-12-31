package generprep
			import 	"fmt"
import 	"github.com/jackc/pgx/pgtype"
import 	"github.com/jackc/pgx"
 
import "prounix.de/pgtools/db"
type Agg_tplantsArray []Agg_tplants

		func (src *Agg_tplantsArray) AssignTo(dst interface{}) error {

			if src != nil {
				ttt, ok := dst.(*Agg_tplantsArray)
				if !ok {
						return fmt.Errorf("cannot assign Agg_tplants")
				}
				*ttt = *src
			}
			return nil
		}
		func (dst *Agg_tplantsArray) Set(src interface{}) error {
			return fmt.Errorf("cannot convert to Agg_tplantsArray")
		}

		func (dst *Agg_tplantsArray) Get() interface{} {
			return dst
		}

						func (dst *Agg_tplantsArray) DecodeBinary(ci *pgtype.ConnInfo, src []byte) error {

							elements := make(Agg_tplantsArray, 0)
							if src == nil {
									*dst = elements
								return nil
							}
							funcer := func (result *Agg_tplantsArray) func()[]pgtype.BinaryDecoder {
								return func() []pgtype.BinaryDecoder{
									pos := len(*result)
									*result = append(*result,Agg_tplants{})
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

db.InitOIDMap["Agg_tplantsArray"]=func(con *pgx.Conn){
db.Register(con , &Agg_tplantsArray{},"Agg_tplantsArray","_agg_tplants","sdbms" )
}


	}

	