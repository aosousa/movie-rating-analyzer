package models

import (
	"fmt"

	"github.com/aosousa/movie-rating-analyzer/maths"
)

// Director struct contains information about a film director and my rating
// associated to their movies
type Director struct {
	Name    string // Director's name
	Ratings []int  // Slice of ratings associated with this director's movies
}

// Directors represents a slice of Director structs
type Directors []Director

// PrintStatistics prints the statistics for a given director
func (director Director) PrintStatistics() {
	ratingAverage := maths.Average(director.Ratings)

	fmt.Printf("%s: %f\n", director.Name, ratingAverage)
}
