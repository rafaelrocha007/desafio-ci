package main

import (
    "testing"
)
	
func TestSoma(t *testing.T) {
    ans := Soma(5, 5)
    if ans != 10 {
        t.Errorf("Soma(5, 5) = %d; resultado esperado é 10", ans)
    }
}