package main

import (
	"fmt"

	"github.com/btwiuse/series/fibonacci"
	"github.com/btwiuse/series/flipper"
	"github.com/btwiuse/series/linear"
	"github.com/btwiuse/series/random"
	"github.com/btwiuse/series/square"
)

func main() {
	for i := 0; i < 100; i++ {
		fmt.Println(
			flipper.Value(), "\t",
			linear.Value(), "\t",
			square.Value(), "\t",
			random.Value(), "\t",
			fibonacci.Value(), "\t",
		)
	}
}
