select x.*, plants from sdbms.tlocation x
left outer join
(
select fulocno,array_agg(t) plants
from sdbms.agg_tplants t
group by fulocno
) t on t.fulocno =x.fulocno
