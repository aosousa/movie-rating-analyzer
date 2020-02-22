package test

import (
	"testing"

	"github.com/aosousa/movie-rating-analyzer/maths"
)

func TestSum(t *testing.T) {
	ratings := []int{5, 8, 9}
	sum := maths.SliceSum(ratings)
	if sum != 22 {
		t.Errorf("SliceSum was incorrect, got: %d, want %d.", sum, 22)
	}
}

func TestAverage(t *testing.T) {
	ratings := []int{4, 3, 3}
	average := maths.Average(ratings)
	if average != 3.3333333 {
		t.Errorf("Average was incorrect, got: %f, want %f", average, 3.3333333)
	}
}
