package main

import (
	"fmt"
	"generic/grammar"
)

func main() {
	grammar.F1()
	grammar.F2()
	grammar.F3()
	grammar.F4()

	var lst = make(grammar.List[int], 20)
	var lst1 = make(grammar.List[float64], 20)
	for i := 0; i < 20; i++ {
		lst[i] = i * i
		lst1[i] = float64(i) * float64(i)
	}

	fmt.Println(lst)
	fmt.Println(lst1)

	grammar.F5()
}
