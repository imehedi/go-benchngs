package main

import (
	"os"

	"github.com/imehedi/Golang/go-benchngs/benchmark"
	_ "github.com/imehedi/Golang/go-benchngs/benchmark"
	"github.com/imehedi/Golang/go-benchngs/fastachopper"
)

func main() {
	filename := os.Args[1]

	fastachopper.Chopper(filename)
	benchmark.BenchNGS("blast", 444)

}
