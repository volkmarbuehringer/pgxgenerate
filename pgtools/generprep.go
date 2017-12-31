// The following directive is necessary to make the package coherent:

package main

import (
	"fmt"
	"os"

	"pgxgenerate/pgtools/db"
	"pgxgenerate/pgtools/prep"
)

func main() {

	defer db.End()

	if len(os.Args) < 6 {
		panic(fmt.Errorf("zuwenig argumente"))
	}
	prefix := os.Args[1]
	importpre := os.Args[2]
	generpath := os.Args[3]
	basis := os.Args[4]
	schema := os.Args[5]
	err := prep.GesamtPrep(prefix, generpath, importpre, basis, schema)
	if err != nil {
		panic(err)
	}

}
