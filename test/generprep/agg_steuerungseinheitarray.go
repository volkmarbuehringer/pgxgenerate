package generprep
			import 	"fmt"
import 	"github.com/jackc/pgx/pgtype"
import 	"github.com/jackc/pgx"
 
import "prounix.de/pgtools/db"
type Agg_steuerungseinheitArray []Agg_steuerungseinheit

		func (src *Agg_steuerungseinheitArray) AssignTo(dst interface{}) error {

			if src != nil {
				ttt, ok := dst.(*Agg_steuerungseinheitArray)
				if !ok {
						return fmt.Errorf("cannot assign Agg_steuerungseinheit")
				}
				*ttt = *src
			}
			return nil
		}
		func (dst *Agg_steuerungseinheitArray) Set(src interface{}) error {
			return fmt.Errorf("cannot convert to Agg_steuerungseinheitArray")
		}

		func (dst *Agg_steuerungseinheitArray) Get() interface{} {
			return dst
		}

						func (dst *Agg_steuerungseinheitArray) DecodeBinary(ci *pgtype.ConnInfo, src []byte) error {

							elements := make(Agg_steuerungseinheitArray, 0)
							if src == nil {
									*dst = elements
								return nil
							}
							funcer := func (result *Agg_steuerungseinheitArray) func()[]pgtype.BinaryDecoder {
								return func() []pgtype.BinaryDecoder{
									pos := len(*result)
									*result = append(*result,Agg_steuerungseinheit{})
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

db.InitOIDMap["Agg_steuerungseinheitArray"]=func(con *pgx.Conn){
db.Register(con , &Agg_steuerungseinheitArray{},"Agg_steuerungseinheitArray","_agg_steuerungseinheit","eeacollector" )
}


	}

	