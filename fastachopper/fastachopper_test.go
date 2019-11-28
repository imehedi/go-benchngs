package fastachopper

import (
	"bytes"
	"io"
	"reflect"
	"testing"
)

func TestChopper(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func Test_buildFasta(t *testing.T) {
	type args struct {
		header string
		seq    bytes.Buffer
	}
	tests := []struct {
		name       string
		args       args
		wantRecord Fasta
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRecord := buildFasta(tt.args.header, tt.args.seq); !reflect.DeepEqual(gotRecord, tt.wantRecord) {
				t.Errorf("buildFasta() = %v, want %v", gotRecord, tt.wantRecord)
			}
		})
	}
}

func Test_check(t *testing.T) {
	type args struct {
		e error
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func Test_parse(t *testing.T) {
	type args struct {
		FastaFh io.Reader
	}
	tests := []struct {
		name string
		args args
		want chan Fasta
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parse(tt.args.FastaFh); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parse() = %v, want %v", got, tt.want)
			}
		})
	}
}
