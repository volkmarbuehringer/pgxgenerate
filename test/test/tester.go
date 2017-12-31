package main

import (
	_ "pgxgenerate/test/gener"
	_ "pgxgenerate/test/generprep"

	"pgxgenerate/pgtools/db"
)

func main() {

	if err := db.Prep(); err != nil {
		panic(err)
	}

	defer db.End()

	/*
		test2()
		test3()
		test55()

		testa()
	*/
	test1()

	test0()
	//testinsr()
	//testins()
	//test3()
	test4()
	test555()

}
