package main

// Different alignment tool produces alignment results in different format.
// Our procedure do not discriminate the tools based on the tool’s own claim
// of accuracy. For example, the E-value reported by the BLAST tool is not
// carried towards the result of our scoring. We collect an information set, R
// from the alignment results containing: (i) read id (r(n)), (ii) reference se130 quence identifier (ref),
// (iii) start position of the read (stp). This information
// is then compared with their counterparts from the GSA.
type Results struct {
	readID            string
	refSeqID          string
	readStartPosition int8
}

func main() {

	print(fOneCalculator(positivePredictiveValueCalculator(5, 4), positiveRateCalculator(3, 5)))

}

// We used true positive rate (r) and positive predictive rate (p), to compute F1 score.
// Sensitivity or true positive rate, alternatively called as Recall.
func fOneCalculator(positivePredictiveValue, positiveRate float64) (fOneScore float64) {
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
