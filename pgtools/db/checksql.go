package db

import (
	"strconv"
	"strings"
)

type parsedat struct {
	d     []string
	maxer int
}

func ParseWhere(sql string) []string {
	erg := parsedat{make([]string, 20), 0}

	erg.preperg(sql)

	return erg.d
}

func (erg *parsedat) preperg(sql string) {
	tt1 := r6.ReplaceAllLiteralString(sql, " ")
	tt2 := r7.Split(tt1, 100)

	erg.parser(tt2, true)
	erg.d = erg.d[1 : erg.maxer+1]
	for idx := range erg.d {
		pos := strings.Index(erg.d[idx], " ")
		if pos >= 0 {
			erg.d[idx] = erg.d[idx][:pos]
		}
	}
}

func (erg *parsedat) parser(ts1 []string, flag bool) {

	for _, f := range ts1 {

		var tok = []string{}
		if flag {
			tok = r8.FindStringSubmatch(f)
			if len(tok) != 3 {
				tok = r9.FindStringSubmatch(f)
				if len(tok) == 3 {
					tok[1], tok[2] = tok[2], tok[1]
				} else {
					//		fmt.Println(f, tok)
					continue
				}
			}

		} else {
			tok = r4.FindStringSubmatch(f)
		}

		if len(tok) == 3 {

			tt2, err := strconv.Atoi(tok[2])
			if err != nil {

				panic(err)

			} else {
				if tt2 > erg.maxer {
					erg.maxer = tt2
				}
				erg.d[tt2] = tok[1]
			}

		}

	}

}

func CheckSQLReturn(x string) bool {
	p := r3.FindStringSubmatchIndex(x)
	if len(p) == 4 {

		return true
	} else {
		return false
	}
}

func CheckSQL(x string) []string {

	ins := r1.FindStringSubmatch(x)
	upd := r2.FindStringSubmatch(x)

	if len(ins) == 4 {
		t1 := ins[2]
		t2 := ins[3]

		ts1 := strings.Split(t1, ",")
		for idx := range ts1 {
			ts1[idx] = strings.Trim(ts1[idx], " \n\r\t")
		}
		erg := make([]string, 100)
		maxer := 0

		ts2 := splitter(t2)

		for i := range ts2 {
			tok := r5.FindStringSubmatch(ts2[i])

			if len(tok) == 2 {

				tt2, err := strconv.Atoi(tok[1])
				if err != nil {
					panic(err)
				}
				if tt2 > maxer {
					maxer = tt2
				}
				erg[tt2] = ts1[i]
			}

		}
		erg = erg[1 : maxer+1]
		return erg

	} else if len(upd) == 4 {

		erg := parsedat{make([]string, 100), 0}

		erg.parser(splitter(upd[2]), false)
		erg.preperg(upd[3])
		//		fmt.Println("updere", maxer, erg)
		return erg.d
	}

	return []string{}
}
