package main

import (
	"fmt"
	"goProjects/goModTry/completeFatRate/newFateRate"
)

func main() {
	name, fatRate := newFateRate.InputfateRate()
	fmt.Printf("hi~,%s的体脂率是%.4f\n", name, fatRate)
}
