package gener
import "prounix.de/pgtools/db"
import 	"prounix.de/test/generprep"
					const LocationName="location" 


type Location struct{
Fulocno db.Int4 `json:"fulocno"`
Fuservicestatid db.Int8 `json:"fuservicestatid"`
Fclocname db.Varchar `json:"fclocname"`
Fclocation db.Varchar `json:"fclocation"`
Funoplants db.Int2 `json:"funoplants"`
Fccountry db.Varchar `json:"fccountry"`
Fctelno db.Varchar `json:"fctelno"`
Fosettime db.Int2 `json:"fosettime"`
Fftimeoffset db.Float4 `json:"fftimeoffset"`
Fubaudrate db.Int4 `json:"fubaudrate"`
Fudatablocksize db.Int2 `json:"fudatablocksize"`
Fuminreq db.Int2 `json:"fuminreq"`
Fudailyreq db.Int2 `json:"fudailyreq"`
Fumonthlyreq db.Int2 `json:"fumonthlyreq"`
Fustatereq db.Int2 `json:"fustatereq"`
Fuweekreq db.Int2 `json:"fuweekreq"`
Fuavailreq db.Int2 `json:"fuavailreq"`
Fopostcode db.Varchar `json:"fopostcode"`
Fodatarequest db.Int2 `json:"fodatarequest"`
Foshortdial db.Int2 `json:"foshortdial"`
Fcip1 db.Varchar `json:"fcip1"`
Foincurrentlist db.Int2 `json:"foincurrentlist"`
Foishdcfileassigned db.Int2 `json:"foishdcfileassigned"`
Plants generprep.Agg_tplantsArray `json:"plants"`
}

var LocationColumns=[]string{
"fulocno",
"fuservicestatid",
"fclocname",
"fclocation",
"funoplants",
"fccountry",
"fctelno",
"fosettime",
"fftimeoffset",
"fubaudrate",
"fudatablocksize",
"fuminreq",
"fudailyreq",
"fumonthlyreq",
"fustatereq",
"fuweekreq",
"fuavailreq",
"fopostcode",
"fodatarequest",
"foshortdial",
"fcip1",
"foincurrentlist",
"foishdcfileassigned",
"plants",
}
type LocationArray []Location

				func (x *LocationArray) Scanner() []interface{} {
					*x = append(*x, Location{})

					return (*x)[len(*x)-1].Scanner()
				}

		func (x *Location) Scanner() []interface {}  {
 return []interface {} {
&x.Fulocno,
&x.Fuservicestatid,
&x.Fclocname,
&x.Fclocation,
&x.Funoplants,
&x.Fccountry,
&x.Fctelno,
&x.Fosettime,
&x.Fftimeoffset,
&x.Fubaudrate,
&x.Fudatablocksize,
&x.Fuminreq,
&x.Fudailyreq,
&x.Fumonthlyreq,
&x.Fustatereq,
&x.Fuweekreq,
&x.Fuavailreq,
&x.Fopostcode,
&x.Fodatarequest,
&x.Foshortdial,
&x.Fcip1,
&x.Foincurrentlist,
&x.Foishdcfileassigned,
&x.Plants,
}
}
