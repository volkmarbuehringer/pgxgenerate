package prep

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func ReadSQL(dir string) (map[string]string, error) {

	var mapperlist map[string]string
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	mapperlist = make(map[string]string)
	for _, f := range files {
		fmt.Println(f.Name())
		name := f.Name()
		pos := strings.Index(name, ".sql")
		if len(name)-4 == pos {
			buf, err := ioutil.ReadFile(dir + "/" + name)
			if err != nil {
				return nil, err
			}
			fmt.Println("hier", name[:pos], string(buf))
			mapperlist[name[:pos]] = string(buf)
		}
	}
	return mapperlist, nil
}
