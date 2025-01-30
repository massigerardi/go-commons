package commons

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Number interface {
	float64 | int64
}

func TestReduceFloat(t *testing.T) {
	type args[T Number, R Number] struct {
		ss           []T
		reducer      func(R, T) R
		initialValue []R
	}
	type testCase struct {
		name   string
		args   args[float64, float64]
		wantEl float64
	}
	tests := []testCase{
		{
			name: "with initial value",
			args: args[float64, float64]{
				ss:           []float64{1.1, 1.2, 1.3},
				reducer:      func(total, n float64) float64 { return total + n },
				initialValue: []float64{1},
			},
			wantEl: 4.6,
		},
		{
			name: "without initial value",
			args: args[float64, float64]{
				ss:           []float64{1.1, 1.2, 1.3},
				reducer:      func(total, n float64) float64 { return total + n },
				initialValue: []float64{},
			},
			wantEl: 3.6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Reduce(tt.args.ss, tt.args.reducer, tt.args.initialValue...)
			assert.InDelta(t, got, tt.wantEl, 0.01)
		})
	}
}

func TestReduceInt(t *testing.T) {
	type args[T Number, R Number] struct {
		ss           []T
		reducer      func(R, T) R
		initialValue []R
	}
	type testCase struct {
		name   string
		args   args[int64, int64]
		wantEl int64
	}
	tests := []testCase{
		{
			name: "with initial value",
			args: args[int64, int64]{
				ss:           []int64{1, 2, 3},
				reducer:      func(total, n int64) int64 { return total + n },
				initialValue: []int64{1},
			},
			wantEl: 7,
		},
		{
			name: "without initial value",
			args: args[int64, int64]{
				ss:           []int64{1, 2, 3},
				reducer:      func(total, n int64) int64 { return total + n },
				initialValue: []int64{},
			},
			wantEl: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Reduce(tt.args.ss, tt.args.reducer, tt.args.initialValue...)
			assert.InDelta(t, got, tt.wantEl, 0.01)
		})
	}
}
