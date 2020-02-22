package main

import (
	"strconv"
	"strings"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/aosousa/movie-rating-analyzer/models"
)

var (
	directors models.Directors
)

func main() {
	XLSXLocation := `Movies.xlsx`

	// open XLSX document
	xlsx, err := excelize.OpenFile(XLSXLocation)
	handleError(err)

	// get all rows in the movie ratings worksheet
	ratings := xlsx.GetRows("movie ratings")

	// loop ratings
	for _, row := range ratings {
		director := row[1]
		rating, _ := strconv.Atoi(strings.Replace(row[2], "/10", "", -1))
		findOrAddDirector(director, rating)
	}

	// loop directors and print their statistics
	for _, director := range directors {
		director.PrintStatistics()
	}
}

/*Check if a director already exists in the slice of directors

Receives:
	* director (string) - Name of the director

Returns:
	* models.Director - Director to use in the event that a director with this name already exists
	* bool - Whether or not the director has already been added to the slice of directors
*/
func directorExists(director string) (models.Director, bool) {
	var tempDirector models.Director

	for _, el := range directors {
		if director == el.Name {
			return el, true
		}
	}

	return tempDirector, false
}

func findOrAddDirector(director string, rating int) {
	tempDirector, foundDirector := directorExists(director)
	if foundDirector {
		// director exists - simply add rating to the slice
		tempDirector.Ratings = append(tempDirector.Ratings, rating)
	} else {
		// director does not exist - create director with a Ratings
		// slice that contains this first value
		director := models.Director{
			Name:    director,
			Ratings: []int{rating},
		}

		directors = append(directors, director)
	}
}
