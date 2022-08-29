package main

import (
	"testing"
)

func TestSoma(t *testing.T) {
	x := 1 + 1
	if x != 2 {
		t.Error("Opa! 1+1 não é igual a 2, obtive", x)
	}
}
