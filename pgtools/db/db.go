// Package db ist ein Singleton für den intitaliserten DB-Pool
package db

import (
	"fmt"
	"net"
	"os"
	"strconv"

	"github.com/jackc/pgx"
	"github.com/jackc/pgx/log/logrusadapter"
	"github.com/sirupsen/logrus"
)

var dbx *pgx.ConnPool

// GetPool liefert den Pool und intitialisiert ihn beim ersten Aufruf
func GetPool() *pgx.ConnPool {
	if dbx == nil {
		initDB()
		fmt.Println("stat", dbx.Stat())
	}
	return dbx
}

//End schliesst den pool
func End() {
	if dbx != nil {
		fmt.Println("stat", dbx.Stat())
		dbx.Close()
	}
}

func initDB() {

	//defer db.Close()

	log := logrus.New()
	log.Out = os.Stdout
	connConfig, err := pgx.ParseEnvLibpq()
	if err != nil {
		panic(fmt.Errorf("DBparse %v", err))
	}
	if g, err := pgx.LogLevelFromString(os.Getenv("PU_DB_LOG")); err != nil {
		panic(fmt.Errorf("DBparse %v", err))
	} else {
		connConfig.LogLevel = int(g)
	}

	logger := logrusadapter.NewLogger(log)
	connConfig.Logger = logger
	poolsize, err := strconv.Atoi(os.Getenv("PU_PG_MAXPOOL"))
	if err != nil {
		poolsize = 10
	}
	config := pgx.ConnPoolConfig{ConnConfig: connConfig, MaxConnections: poolsize, AfterConnect: SetTyp}

	if dbx, err = pgx.NewConnPool(config); err != nil {
		fmt.Printf("connect %T %+v\n", err, err)

		if perr, ok := err.(*net.OpError); ok {
			panic(fmt.Errorf("connect %+v", *perr))
		}

		if perr, ok := err.(pgx.PgError); ok {
			panic(fmt.Errorf("connect %+v", perr))
		}
		panic("abbruch")
	}

}
