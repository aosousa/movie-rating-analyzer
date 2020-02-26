package models

import (
	"fmt"
	"sort"
)

// Director struct contains information about a film director and my rating
// associated to their movies
type Director struct {
	Name          string  // Director's name
	Ratings       []int   // Slice of ratings associated with this director's movies
	AverageRating float32 // Average rating for a director's movies
}

// Directors represents a slice of Director structs
type Directors []Director

// PrintStatistics prints the statistics for all directors found
func (directors Directors) PrintStatistics() {
	for _, director := range directors {
		fmt.Printf("%s: %.2f\n", director.Name, director.AverageRating)
	}
}

// SortAscendingByName sorts the directors slice by alphabetical ascending order and print the statistics in that order
func (directors Directors) SortAscendingByName() {
	sort.Slice(directors, func(i, j int) bool {
		return directors[i].Name < directors[j].Name
	})

	directors.PrintStatistics()
}

// SortDescendingByName sorts the directors slice by alphabetical descending order and print the statistics in that order
func (directors Directors) SortDescendingByName() {
	sort.Slice(directors, func(i, j int) bool {
		return directors[i].Name > directors[j].Name
	})

	directors.PrintStatistics()
}

// SortAscendingByAverage sorts the directors slice by average rating in ascending order and print the statistics in that order
func (directors Directors) SortAscendingByAverage() {
	sort.Slice(directors, func(i, j int) bool {
		return directors[i].AverageRating < directors[j].AverageRating
	})

	directors.PrintStatistics()
}

// SortDescendingByAverage sorts the directors slice by average rating in descending order and print the statistics in that order
func (directors Directors) SortDescendingByAverage() {
	sort.Slice(directors, func(i, j int) bool {
		return directors[i].AverageRating > directors[j].AverageRating
	})

	directors.PrintStatistics()
}
