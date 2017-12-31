create or replace view eeacollector.agg_eeaopcserver  as
select
 eea_opcid    ,
 eea_sernr     ,
 eea_plantnr    ,
 eea_typnr      ,
 eea_typstr      ,
 eea_hersteller   ,
 eea_nennleistung ,
 eea_anzminuten   ,
 eea_aktiv
from eeacollector.eeaopcserver

create view sdbms.agg_tplants as select * from sdbms.tplants


create view eeacollector.agg_steuerungseinheit as select * from eeacollector.steuerungseinheit

  
