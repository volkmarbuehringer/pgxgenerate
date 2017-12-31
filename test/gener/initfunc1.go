package gener

		import "prounix.de/pgtools/db"
			
func init(){

db.InitScanMap[BoreasName]= func ()( []interface {}, interface{},[]string ) 	{
 var x Boreas
return x.Scanner(), &x,BoreasColumns

	}
db.PreqSQLMap[BoreasName]= `select *
from opcserver
left outer join (select eea_opcid,array_agg(t order by eea_plantnr) eeaopcserver
from agg_eeaopcserver t
group by eea_opcid ) t1 on eea_opcid = opc_id
left outer join (select str_opcid , array_agg( t order by str_plantnr) steuerungseinheit
from agg_steuerungseinheit t
group by str_opcid) t2 on str_opcid = opc_id
`

db.InitScanMap[LocationName]= func ()( []interface {}, interface{},[]string ) 	{
 var x Location
return x.Scanner(), &x,LocationColumns

	}
db.PreqSQLMap[LocationName]= `select x.*, plants from sdbms.tlocation x
left outer join
(
select fulocno,array_agg(t) plants
from sdbms.agg_tplants t
group by fulocno
) t on t.fulocno =x.fulocno
`

db.InitScanMap[Test1Name]= func ()( []interface {}, interface{},[]string ) 	{
 var x Test1
return x.Scanner(), &x,Test1Columns

	}
db.PreqSQLMap[Test1Name]= `select * from opcserver
`

db.InitScanMap[TestinsName]= func ()( []interface {}, interface{},[]string ) 	{
 var x Testins
return x.Scanner(), &x,TestinsColumns

	}
db.PreqSQLMap[TestinsName]= `insert  into  opcserver   ( opc_name,opc_aktiv,opc_uri,opc_scadanr,opc_letzteanfrage)
  values  ( $1,false,$2,$3,now())
`

db.InitScanMap[TestinsrReturnName]= func ()( []interface {}, interface{},[]string ) 	{
 var x TestinsrReturn
return x.Scanner(), &x,TestinsrReturnColumns

	}
db.InitScanMap[TlocationName]= func ()( []interface {}, interface{},[]string ) 	{
 var x Tlocation
return x.Scanner(), &x,TlocationColumns

	}
db.PreqSQLMap[TlocationName]= `select * from sdbms.tlocation
`

db.InitScanMap[InsplantName]= func ()( []interface {}, interface{},[]string ) 	{
 var x Insplant
return x.Scanner(), &x,InsplantColumns

	}
db.PreqSQLMap[InsplantName]= `insert   into  sdbms.tplants (
fuplantid   ,
 fulocno     ,
 fuplantno     ,
 fuplantcode     ,
 fuspecialid     ,
 fuserieno       ,
 fdcomdate       ,
 fdturnoff      ,
 fcplantalias    ,
 fcplanthwtype   ,
 fuplantpower    ,
 foincurrentlist ,
 foserialchanged
)   values  ($1,
  $2,
  $3,
  $4,
  $5,
  $6,
  $7,
  $8,
  $9,
  $10,
  $11,
  $12,
  $13
  )
`

db.InitScanMap[Test0Name]= func ()( []interface {}, interface{},[]string ) 	{
 var x Test0
return x.Scanner(), &x,Test0Columns

	}
db.PreqSQLMap[Test0Name]= `update   opcserver   set opc_tcptim=$1,opc_preverror=$2,opc_crdate=now(),opc_error=null   where opc_id=$4 and substr($3,1,1) = substr(opc_name,1,1)
`

db.InitScanMap[TestinsrName]= func ()( []interface {}, interface{},[]string ) 	{
 var x Testinsr
return x.Scanner(), &x,TestinsrColumns

	}
db.PreqSQLMap[TestinsrName]= `insert  into
  opcserver ( opc_name,opc_aktiv,opc_uri,opc_scadanr,opc_letzteanfrage) values( $1,false,$2,$3,now())   returning *
`

}

