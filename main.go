package main

import (
	"fmt"
	"go-design-pattern/pkg/solid"
)

func main() {
	arr := [] string{"a","b"}
	entry := solid.NewJounal(arr)
	entry.AddEntry("add")
	entry.AddEntry("second")
	result := entry.String()
	fmt.Println(result)
}

