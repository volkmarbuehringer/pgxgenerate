// The following directive is necessary to make the package coherent:

package main

import (
	"fmt"
	"os"

	_ "prounix.de/test/generprep"

	"prounix.de/pgtools/db"
	"prounix.de/pgtools/prep"
)

func main() {

	defer db.End()
	flag, err := prep.Gesamt("test", "prounix.de")

	if err != nil {
		panic(err)
	}
	if flag {
		fmt.Println("programmende ok mit code")

		os.Exit(1)
	}
}
