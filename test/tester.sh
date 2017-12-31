export PGHOST=localhost
export PGPORT=5432
export PGDATABASE="boreas"
export PGUSER=postgres
export PGPASSWORD=postgres
export PGSSLMODE="allow"
export PGAPPNAME="goofer"

export PU_SCHEMA="eeacollector,sdbms"

export PU_DB_LOG="warn" ##loglevel von pgx

go run test/*.go
