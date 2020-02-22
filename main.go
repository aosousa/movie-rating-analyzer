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
	createDirectorSlice()

	// loop directors and print their statistics
	for _, director := range directors {
		director.PrintStatistics()
	}
}

// Populate the slice of directors with data
func createDirectorSlice() {
	XLSXLocation := `Movies.xlsx`

	// open XLSX document
	xlsx, err := excelize.OpenFile(XLSXLocation)
	handleError(err)

	// get all rows in the movie ratings worksheet
	ratings := xlsx.GetRows("movie ratings")

	// loop ratings
	for _, row := range ratings {
		director := row[1]
		rating, _ := strconv.Atoi(strings.Replace(row[2], "/10", "", -1)) // remove the "/10" section from the ratings column
		handleDirector(director, rating)
	}
}

/*Handle actions related to the creation or update of a director.
If a director already exists in the slice of directors, their slice of ratings will be updated.
Otherwise, a new director will be created and added to the slice of directors

Receives:
	* director (string) - Name of the director
	* rating (int) - Movie rating
*/
func handleDirector(director string, rating int) {
	directorExists := false

	for i := range directors {
		if director == directors[i].Name {
			// director exists - simply add rating to the slice
			directorExists = true
			directors[i].Ratings = append(directors[i].Ratings, rating)
		}
	}

	if !directorExists {
		// director does not exist - create director with a Ratings
		// slice that contains this first value
		newDirector := models.Director{
			Name:    director,
			Ratings: []int{rating},
		}

		directors = append(directors, newDirector)
	}
}
