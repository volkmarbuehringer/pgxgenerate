select *
from opcserver
left outer join (select eea_opcid,array_agg(t order by eea_plantnr) eeaopcserver
from agg_eeaopcserver t
group by eea_opcid ) t1 on eea_opcid = opc_id
left outer join (select str_opcid , array_agg( t order by str_plantnr) steuerungseinheit
from agg_steuerungseinheit t
group by str_opcid) t2 on str_opcid = opc_id
