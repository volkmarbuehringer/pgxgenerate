// The following directive is necessary to make the package coherent:

package main

import (
	"fmt"
	"os"

	"prounix.de/pgtools/db"
	"prounix.de/pgtools/prep"
)

//"agg%" "prounix.de/test" "../generprep"

func main() {

	defer db.End()

	if len(os.Args) < 5 {
		panic(fmt.Errorf("zuwenig argumente"))
	}
	prefix := os.Args[1]
	importpre := os.Args[2]
	generpath := os.Args[3]
	schema := os.Args[4]
	err := prep.GesamtPrep(prefix, generpath, importpre, schema)
	if err != nil {
		panic(err)
	}

}
