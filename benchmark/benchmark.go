package benchmark

import (
	"fmt"
)

// Simulated short read sequences characteristics chart
type ShortSequence struct {
	seqID             string // Simulated short reads, q(n) ⊂ Q of variant bp lengths n ⊂ {50, 100, 150, 400};
	refSeqID          string
	readStartPosition int8
}

var simulatedReadsParameters ShortSequence
var alignmentResultsParameters ShortSequence

func BenchNGS(toolName string, seqLength int) (fOneScore float64) {
	// ToDo 1: Align reads to model genome
	// ToDo 2: From new alignment results generate R ← {q(n), ref, stp}
	// ToDo 3: Compare R with GSA and produce a set S ← {T P, F P, F N}

	fmt.Printf("%v", toolName)
	fmt.Printf("%d", seqLength)

	myEntry := ShortSequence{
		"read1",
		"refSeq1",
		10,
	}
	myResults := ShortSequence{
		"read1",
		"refSeq1",
		10,
	}
	return fOneCalculator(myEntry, myResults)
}

// Rijsbergen’s accuracy measurement score named here as fOneScore,
// We used true positive rate (r) and positive predictive rate (p), to compute F1 score.
// Sensitivity or true positive rate, alternatively called as Recall.
func fOneCalculator(simulatedReadsParameters, alignmentResultsParameters ShortSequence) (fOneScore float64) {

	var tp, fp, fn, positivePredictiveValue, positiveRate float64
	fmt.Printf("%f,%f,%f", tp, fp, fn)
	positivePredictiveValue = positivePredictiveValueCalculator(5, 5)
	positiveRate = positiveRateCalculator(4, 5)
	return 2 * ((positivePredictiveValue * positiveRate) / (positivePredictiveValue + positiveRate))
}

// Positive predictive value or alternatively called as Precision. Precision
// measures the probability of positively mapped reads by dividing the total
// number of correct alignments by total number of alignments detected by the
// tool: p = T P/(T P + F P).

func positivePredictiveValueCalculator(truePositive, falsePositive float64) (positivePredictiveValue float64) {
	return truePositive / (truePositive + falsePositive)
}

// A recall measures the probability of actually mapped reads. The true positive rate
// is computed by dividing the total number of correct results by the
// number of alignments that were expected: r = T P/(T P + F N).

func positiveRateCalculator(truePositive, falseNegative float64) (positiveRate float64) {
	return truePositive / (truePositive + falseNegative)
}
