#export PGHOST=localhost
#export PGPORT=5432
#export PGDATABASE="boreas"
#export PGUSER=postgres
#export PGPASSWORD=postgres
#export PGSSLMODE="allow"
#export PGAPPNAME=goofer


#lokales postgres
export PGHOST=localhost
export PGPORT=5432
export PGDATABASE="boreas"
export PGUSER=postgres
export PGPASSWORD=postgres
export PGSSLMODE="allow"
export PGAPPNAME="goofer"

export PU_SCHEMA="eeacollector,sdbms"

export PU_DB_LOG="debug" ##loglevel von pgx

rm generprep/*.go
rm gener/*.go

go run ../pgtools/generprep.go "agg%,aggv%" "prounix.de/test" "../test/generprep" "eeacollector,sdbms"

go run generate.go



if [ $? -eq 0 ]
then
  echo "Successfully created file"
    exit 0
else
  echo "!!!!!!!!!!!!!!!!!!!!!!!!"

fi
