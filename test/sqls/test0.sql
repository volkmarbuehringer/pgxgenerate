update   opcserver   set opc_tcptim=$1,opc_preverror=$2,opc_crdate=now(),opc_error=null   where opc_id=$4 and substr($3,1,1) = substr(opc_name,1,1)
