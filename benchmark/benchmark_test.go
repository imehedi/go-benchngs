package benchmark

import "testing"

func TestBenchNGS(t *testing.T) {
	type args struct {
		toolName  string
		seqLength int
	}
	tests := []struct {
		name          string
		args          args
		wantFOneScore float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotFOneScore := BenchNGS(tt.args.toolName, tt.args.seqLength); gotFOneScore != tt.wantFOneScore {
				t.Errorf("BenchNGS() = %v, want %v", gotFOneScore, tt.wantFOneScore)
			}
		})
	}
}

func Test_fOneCalculator(t *testing.T) {
	type args struct {
		simulatedReadsParameters   ShortSequence
		alignmentResultsParameters ShortSequence
	}
	tests := []struct {
		name          string
		args          args
		wantFOneScore float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotFOneScore := fOneCalculator(tt.args.simulatedReadsParameters, tt.args.alignmentResultsParameters); gotFOneScore != tt.wantFOneScore {
				t.Errorf("fOneCalculator() = %v, want %v", gotFOneScore, tt.wantFOneScore)
			}
		})
	}
}

func Test_positivePredictiveValueCalculator(t *testing.T) {
	type args struct {
		truePositive  float64
		falsePositive float64
	}
	tests := []struct {
		name                        string
		args                        args
		wantPositivePredictiveValue float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotPositivePredictiveValue := positivePredictiveValueCalculator(tt.args.truePositive, tt.args.falsePositive); gotPositivePredictiveValue != tt.wantPositivePredictiveValue {
				t.Errorf("positivePredictiveValueCalculator() = %v, want %v", gotPositivePredictiveValue, tt.wantPositivePredictiveValue)
			}
		})
	}
}

func Test_positiveRateCalculator(t *testing.T) {
	type args struct {
		truePositive  float64
		falseNegative float64
	}
	tests := []struct {
		name             string
		args             args
		wantPositiveRate float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotPositiveRate := positiveRateCalculator(tt.args.truePositive, tt.args.falseNegative); gotPositiveRate != tt.wantPositiveRate {
				t.Errorf("positiveRateCalculator() = %v, want %v", gotPositiveRate, tt.wantPositiveRate)
			}
		})
	}
}
