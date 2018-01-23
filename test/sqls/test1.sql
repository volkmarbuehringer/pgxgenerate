select * from opcserver
where opc_name like $1 and opc_scadanr > $2 and opc_aktiv=$3
