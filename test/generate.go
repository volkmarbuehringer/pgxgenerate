// The following directive is necessary to make the package coherent:

package main

import (
	"fmt"
	"os"

	_ "pgxgenerate/test/generprep"

	"pgxgenerate/pgtools/db"
	"pgxgenerate/pgtools/prep"
)

func main() {

	defer db.End()
	flag, err := prep.Gesamt("test", "pgxgenerate")

	if err != nil {
		panic(err)
	}
	if flag {
		fmt.Println("programmende ok mit code")

		os.Exit(1)
	}
}
