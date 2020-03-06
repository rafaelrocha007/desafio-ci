package main

import (
    "fmt"
    "testing"
)
	

func Soma(a, b int) int {
    return a + b;
}
	

func TestSoma(t *testing.T) {
    ans := Soma(5, 5)
    if ans != 10 {
        t.Errorf("Soma(5, 5) = %d; want 10", ans)
    }
}


func main() {
    fmt.Println(Soma(5,5))
}