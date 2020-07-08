package main

import (
	"flag"
	"os"

	"github.com/imehedi/go-benchngs/benchmark"
	"github.com/imehedi/go-benchngs/fastachopper"
)

func main() {

	fileName := flag.String("input", "sample.fasta", "Input Fasta file to split. (Required)")
	//metricPtr := flag.String("metric", "chars", "Metric {chars|words|lines};. (Required)")
	//uniquePtr := flag.Bool("unique", false, "Measure unique values of a metric.")
	flag.Parse()

	if *fileName == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	fastachopper.Chopper(*fileName)
	benchmark.BenchNGS("blast", 444)

}
