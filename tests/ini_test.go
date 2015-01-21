package main

import (
	"fmt"
	"testing"

	"github.com/kgbu/gorandom"
)

func Test001(t *testing.T) {
	fmt.Println("====Test001: check 3")
	if 3 != gorandom.Act() {
		t.Errorf("Ini Test001 failded")
	}
}
