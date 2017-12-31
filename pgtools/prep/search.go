package prep

import (
	"github.com/jackc/pgx"
	"github.com/jackc/pgx/pgtype"
	"github.com/pkg/errors"

	"pgxgenerate/pgtools/db"
)

func search(con *pgx.Conn, oid pgtype.OID, name string, prepTypes *map[string]prepTyp, schema string) error {
	var typ prepTyp
	var namer *string

	if len(name) > 0 {
		namer = &name
	}
	x := db.Schemas()
	if len(schema) > 0 {
		x = schema
	}
	sql := `select
b.typname,c.typname,
b.oid,c.oid,n.nspname
from pg_namespace n
inner join				 pg_type a on a.typnamespace= n.oid
inner join pg_type b on b.oid = case when a.typelem > 0 then a.typelem else a.oid end
inner join pg_type c on c.oid = case when a.typarray > 0 then a.typarray else a.oid end
where a.oid= case when $1 > 0 then $1 else a.oid end
and n.nspname =any($2)  and
a.typname=coalesce($3,a.typname)`

	if err := con.QueryRow(sql, oid, "{"+x+"}", &namer).Scan(&typ.aname,
		&typ.bname, &typ.aoid, &typ.boid, &typ.schema); err != nil {

		return errors.Wrapf(err, "fehler suche %d %s", oid, x)
	}

	(*prepTypes)[typ.aname] = typ

	return nil
}
