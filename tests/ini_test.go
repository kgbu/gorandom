package main

import (
	"testing"

	"github.com/kgbu/gorandom"
)

func Test001(t *testing.T) {
	if 3 != gorandom.Act() {
		t.Errorf("Ini Test001 failded")
	}
}
