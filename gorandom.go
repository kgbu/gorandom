package gorandom

import (
	"fmt"
)

func Act() int {
        _, err := Mad()
	if err != nil {
		fmt.Println("return 3")
	}
	return 3
}

type MadType struct {
	Str string
}

func Mad()( *MadType, error) {
    m := new(MadType)
    return m, nil
}
