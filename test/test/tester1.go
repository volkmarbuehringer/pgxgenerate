package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"

	"github.com/jackc/pgx"
	"github.com/jackc/pgx/pgtype"
	"prounix.de/test/gener"
	"prounix.de/test/generprep"

	"prounix.de/pgtools/db"
)

func test555() {
	pool := db.GetPool()
	rows, err := pool.Query(gener.BoreasName)
	if err != nil {
		panic(err)
	}
	start := time.Now()
	storer := make(gener.BoreasArray, 0, 10000)
	for rows.Next() {
		if err := rows.Err(); err != nil {
			panic(err)
		}
		if err := rows.Scan(storer.Scanner()...); err != nil {
			panic(err)
		}

	}
	store := make(map[int32]*gener.Boreas)
	for idx := range storer {
		store[storer[idx].Opc_scadanr.Int] = &storer[idx]
	}

	fmt.Println("eingelesen boreas:", time.Since(start), len(storer), len(store))
}

/*
func testblubber() {
	pool := db.GetPool()
	tx, err := pool.Begin()
	if err != nil {
		panic(err)
	}
	defer tx.Rollback()
	obj, err := tx.LargeObjects()
	if err != nil {
		panic(err)
	}

	rows, err := pool.Query(gener.BlubberName)
	if err != nil {
		panic(err)
	}
	start := time.Now()
	var x gener.Blubber

	for rows.Next() {
		if err := rows.Err(); err != nil {
			panic(err)
		}
		if err := rows.Scan(x.Scanner()...); err != nil {
			panic(err)
		}
		fmt.Println(x)
		blubber, err := obj.Open(x.Raster, pgx.LargeObjectModeRead)
		if err != nil {
			panic(err)
		}
		if f, err := os.Create(x.Name.String); err != nil {
			panic(err)
		} else {
			w := bufio.NewWriter(f)
			len, err := io.Copy(w, blubber)
			if err != nil {
				panic(err)
			}
			if err := w.Flush(); err != nil {
				panic(err)
			}
			if err := f.Close(); err != nil {
				panic(err)
			}

			fmt.Println(len, x)

		}
		if err := blubber.Close(); err != nil {
			panic(err)
		}
	}
	fmt.Println("eingelesen:", time.Since(start))
}
*/

/*
func testa() {
	pool := db.GetPool()
	rows, err := pool.Query(gener.TestaName)
	if err != nil {
		panic(err)
	}
	start := time.Now()
	store := make(map[int32]gener.Testa)
	var x gener.Testa

	for rows.Next() {
		if err := rows.Err(); err != nil {
			panic(err)
		}
		if err := rows.Scan(x.Scanner()...); err != nil {
			panic(err)
		}
		store[x.Hop_id.Int] = x


	}
	fmt.Println("eingelesen:", time.Since(start), len(store))
}

func test55() {
	pool := db.GetPool()
	rows, err := pool.Query(gener.Test55Name)
	if err != nil {
		panic(err)
	}
	start := time.Now()
	store := make(map[int32]gener.Test55)
	var x gener.Test55

	for rows.Next() {
		if err := rows.Err(); err != nil {
			panic(err)
		}
		if err := rows.Scan(x.Scanner()...); err != nil {
			panic(err)
		}
		store[x.Opl_id.Int] = x


	}
	fmt.Println("eingelesen:", time.Since(start), len(store))
}

*/
type CopyTest struct {
	syncer chan<- *gener.Location
	rows   *pgx.Rows
	err    error
	dat    gener.Location
	out    gener.Tlocation
}

type CopyTest1 struct {
	syncer <-chan *gener.Location
	err    error
	dat    *gener.Location
	out    gener.Insplant
	pos    int
}

func cleanup() {
	if r := recover(); r != nil {
		fmt.Println("recover", r)
	}
}
func (x *CopyTest) Next() bool {

	defer cleanup()
	if x.err != nil {
		return false
	}

	for x.rows.Next() {
		if x.err = x.rows.Err(); x.err != nil {
			close(x.syncer)
			return false
		}

		if x.err = x.rows.Scan(x.dat.Scanner()...); x.err != nil {
			close(x.syncer)
			return false
		}

		if len(x.dat.Plants) > 0 {

			return true
		}

	}
	close(x.syncer)
	return false
}

func (x *CopyTest1) Next() bool {
	defer cleanup()
	if x.err != nil {
		return false
	}
	if x.pos >= 0 && x.pos < len(x.dat.Plants)-1 {
		x.pos++
		return true
	}
	var ok bool
	x.dat, ok = <-x.syncer
	if !ok {
		return false
	}
	x.pos = 0
	return true

}

func (x *CopyTest) Values() ([]interface{}, error) {
	defer cleanup()
	if x.err != nil {
		return nil, x.err
	}
	d := x.dat
	x.syncer <- &d
	x.out = gener.Tlocation{
		d.Fulocno,
		d.Fuservicestatid,
		d.Fclocname,
		d.Fclocation,
		d.Funoplants,
		d.Fccountry,
		d.Fctelno,
		d.Fosettime,
		d.Fftimeoffset,
		d.Fubaudrate,
		d.Fudatablocksize,
		d.Fuminreq,
		d.Fudailyreq,
		d.Fumonthlyreq,
		d.Fustatereq,
		d.Fuweekreq,
		d.Fuavailreq,
		d.Fopostcode,
		d.Fodatarequest,
		d.Foshortdial,
		d.Fcip1,
		d.Foincurrentlist,
		d.Foishdcfileassigned,
	}

	return x.out.Scanner(), x.err

}

func (x *CopyTest1) Values() ([]interface{}, error) {
	defer cleanup()
	if x.err != nil {
		return nil, x.err
	}

	d := x.dat.Plants[x.pos]
	x.out = gener.Insplant{
		d.Fuplantid,
		d.Fulocno,
		d.Fuplantno,
		d.Fuplantcode,
		d.Fuspecialid,
		d.Fuserieno,
		d.Fdcomdate,
		d.Fdturnoff,
		d.Fcplantalias,
		d.Fcplanthwtype,
		d.Fuplantpower,
		d.Foincurrentlist,
		d.Foserialchanged,
	}
	x.out.Fuplantid.Int += 7000000
	return x.out.Scanner(), x.err

}
func (x *CopyTest) Err() error {
	return x.err
}

func (x *CopyTest1) Err() error {
	return x.err
}

func test4() {
	pool := db.GetPool()

	rows, err := pool.Query(gener.LocationName)
	if err != nil {
		panic(err)
	}
	tx1, err := pool.Begin()
	if err != nil {
		panic(err)
	}
	tx2, err := pool.Begin()
	if err != nil {
		panic(err)
	}
	syncer := make(chan *gener.Location, 500)
	var g errgroup.Group
	var ttt = CopyTest{rows: rows, syncer: syncer}
	var ttt1 = CopyTest1{pos: -1, syncer: syncer}
	fmt.Println("starte copy")
	start := time.Now()

	g.Go(func() error {
		copyCount, err := tx1.CopyFrom(
			pgx.Identifier{"gaga"},
			gener.InsplantColumns,
			pgx.CopyFromSource(&ttt1),
		)
		fmt.Println("fe", copyCount, err)
		return err
	})
	g.Go(func() error {
		copyCount, err := tx2.CopyFrom(
			pgx.Identifier{"gaga1"},
			gener.TlocationColumns,
			pgx.CopyFromSource(&ttt),
		)
		fmt.Println(copyCount, err)
		return err
	})
	err = g.Wait()
	if err != nil {
		tx1.Rollback()
		tx2.Rollback()
		panic(err)
	}
	tx1.Commit()
	tx2.Commit()
	fmt.Println(err, time.Since(start))

}

func test3() {
	pool := db.GetPool()
	rows, err := pool.Query(gener.LocationName)
	if err != nil {
		panic(err)
	}

	arr := make(gener.LocationArray, 0, 10000)
	start := time.Now()

	for rows.Next() {
		if err := rows.Err(); err != nil {
			panic(err)
		}
		if err := rows.Scan(arr.Scanner()...); err != nil {
			panic(err)
		}

	}
	store := make(map[int32]*gener.Location, len(arr))
	stor1 := make(map[struct {
		loc   int32
		plant int32
	}]*generprep.Agg_tplants)
	for idx := range arr {
		store[arr[idx].Fulocno.Int] = &arr[idx]
		for i := range arr[idx].Plants {
			stor1[struct {
				loc   int32
				plant int32
			}{arr[idx].Plants[i].Fulocno.Int,
				arr[idx].Plants[i].Fuplantid.Int,
			}] = &arr[idx].Plants[i]

		}

	}
	fmt.Println("eingelesen:", len(stor1), len(store), len(arr))
	batch := pool.BeginBatch()
	for _, d := range stor1 {
		var inser = gener.Insplant{
			d.Fuplantid,
			d.Fulocno,
			d.Fuplantno,
			d.Fuplantcode,
			d.Fuspecialid,
			d.Fuserieno,
			d.Fdcomdate,
			d.Fdturnoff,
			d.Fcplantalias,
			d.Fcplanthwtype,
			d.Fuplantpower,
			d.Foincurrentlist,
			d.Foserialchanged,
		}
		inser.Fuplantid.Int += 7000000
		batch.Queue(gener.InsplantName, inser.Scanner(), []pgtype.OID{}, []int16{})

	}

	ctx := context.Background()
	if err := batch.Send(ctx, nil); err != nil {
		fmt.Println("hier fehler")
		panic(err)
	}

	for i := 0; i < 10000; i++ {
		flag, err := batch.ExecResults()
		if err != nil {
			panic(err)
		}
		if flag.RowsAffected() != 1 {
			panic("alles kaputt")
		}

	}

	fmt.Println("nach ins:", time.Since(start))
	if err := batch.Close(); err != nil {
		fmt.Println("hier fehler")
		panic(err)
	}

}

func testinsr() {
	pool := db.GetPool()
	var input gener.Testinsr
	var output gener.TestinsrReturn
	input.Opc_name.Set("serverda")
	//input.Addtime.Set(time.Now())
	input.Opc_scadanr.Set(3453334)
	input.Opc_uri.Set("gagaga")
	err1 := pool.QueryRow(gener.TestinsrName, input.Scanner()...).Scan(output.Scanner()...)
	if err1 != nil {
		panic(err1)
	}
	fmt.Println("fertig", output)
}

func testins() {
	con, err := db.GetPool().Acquire()
	if err != nil {
		panic(err)
	}

	batch := con.BeginBatch()

	for i := 0; i < 10000; i++ {
		//cmd, err1 := pool.Exec(gener.TestinsName, input.Scanner()...)
		var input gener.Testins
		input.Opc_name.Set("serverda")
		//input.Addtime.Set(time.Now())
		input.Opc_scadanr.Set(i + 1000000)
		input.Opc_uri.Set("gagaga")
		batch.Queue(gener.TestinsName, input.Scanner(), []pgtype.OID{}, []int16{})

	}

	ctx := context.Background()
	if err := batch.Send(ctx, nil); err != nil {
		fmt.Println("hier fehler")
		panic(err)
	}

	start := time.Now()
	for i := 0; i < 10000; i++ {
		flag, err := batch.ExecResults()
		if err != nil {
			panic(err)
		}
		if flag.RowsAffected() != 1 {
			panic("alles kaputt")
		}

	}

	fmt.Println("nach ins:", time.Since(start))
	if err := batch.Close(); err != nil {
		fmt.Println("hier fehler")
		panic(err)
	}

}

func test0() {
	pool := db.GetPool()
	var input gener.Test0
	input.Opc_id.Set(530)
	input.Opc_name.Set("gaga")
	//input.Addtime.Set(time.Now())
	input.Opc_preverror.Set("gagaga")
	input.Opc_tcptim.Set(555)

	cmd, err1 := pool.Exec(gener.Test0Name, input.Scanner()...)

	if err1 != nil {
		panic(err1)
	}
	fmt.Println("nach upd", cmd.RowsAffected())
}

func test1() {
	pool := db.GetPool()
	rows, err := pool.Query(gener.Test1Name)
	if err != nil {
		panic(err)
	}
	var x gener.Test1
	for rows.Next() {
		if err := rows.Err(); err != nil {
			panic(err)
		}

		if err := rows.Scan(x.Scanner()...); err != nil {
			panic(err)
		}
		ga, _ := json.Marshal(x)
		fmt.Println(string(ga))

	}

}

/*
func test2() {
	pool := db.GetPool()
	rows, err := pool.Query(gener.Test2Name)
	if err != nil {
		panic(err)
	}
	store := make(map[int32]gener.Test2)
	for rows.Next() {
		if err := rows.Err(); err != nil {
			panic(err)
		}
		var x gener.Test2
		if err := rows.Scan(x.Scanner()...); err != nil {
			panic(err)
		}
		if len(x.Agg) != int(x.Sizer.Int) {
			fmt.Println("falsche länge", x.Opc_id.Int, len(x.Agg), int(x.Sizer.Int))
			panic("kaputt")
		}
		fmt.Println(len(x.Agg), x.Sizer.Int, x.Opc_id.Int, x.Opc_name.String)
		store[x.Opc_id.Int] = x
		/*
			for i, d := range x.Agg {
				fmt.Println("hier row", i, d.Wop_id.Int, d.Wop_plantnr.String)
			}
*/
//	}
//	fmt.Println("eingelesen:", len(store))
//}

/*
func test2b() {
	pool := db.GetPool()
	var input gener.Test2b
	input.Url.Set("willi")
	//input.Addtime.Set(time.Now())
	input.Addtime.Status = pgtype.Null
	input.Id.Set(6)
	input.Error.Set("la")
	cmd, err1 := pool.Exec(gener.Test2bName, input.Scanner()...)

	if err1 != nil {
		panic(err1)
	}
	fmt.Println("nach upd", cmd.RowsAffected())
}

func test2a() {
	pool := db.GetPool()
	var input gener.Test2a
	input.Url.Set("lalala111lxxl")
	//input.Addtime.Set(time.Now())
	input.Addtime.Status = pgtype.Null
	gag, err := pool.Exec(gener.Test2aName, input.Scanner()...)

	if err != nil {
		panic(err)
	}

	fmt.Println(gag)
}

func test4() {
	pool := db.GetPool()
	rows, err := pool.Query(gener.Test4Name)
	if err != nil {
		panic(err)
	}
	var x gener.Test4
	for rows.Next() {
		if err := rows.Err(); err != nil {
			panic(err)
		}

		if err := rows.Scan(x.Scanner()...); err != nil {
			panic(err)
		}
		json, err := json.Marshal(x)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(json))
		fmt.Println("daten", len(x.X), x.Url.String)
		for i, p := range x.X {
			fmt.Println("row", i, p.Url.String, p.Error.String)
		}

	}

}

func test2() {
	pool := db.GetPool()
	rows, err := pool.Query(gener.Test2Name)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		if err := rows.Err(); err != nil {
			panic(err)
		}
		var x gener.Test2
		if err := rows.Scan(x.Scanner()...); err != nil {
			panic(err)
		}
		if x.Duration.Status == pgtype.Present {
			fmt.Println("is not null")
		}
		for i, g := range x.Gaga.Elements {
			fmt.Println("ga", i, g.Int)
		}
		//fmt.Println(x.Url.String, x.Id.Int, x.Error.String)
	}

}
*/
