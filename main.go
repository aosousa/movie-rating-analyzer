package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/aosousa/movie-rating-analyzer/maths"
	"github.com/aosousa/movie-rating-analyzer/models"
)

var (
	directors models.Directors
)

const (
	version = "1.0.0"
)

func main() {
	// create slice of Director structs
	createDirectorSlice()

	// check if there were extra arguments passed in the command line
	args := os.Args

	if len(args) != 1 {
		cmd := args[1]

		switch cmd {
		case "-o", "--order":
			handleSortingSlice(os.Args)
		case "-h", "--help":
			printHelp()
		case "-v", "--version":
			printVersion()
		default:
			directors.PrintStatistics()
		}
	} else {
		directors.PrintStatistics()
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

	// iterate over the final directors slice to calculate the average
	for i := range directors {
		directors[i].AverageRating = maths.Average(directors[i].Ratings)
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

/*Handle sorting options
Receives:
	* args ([]string) - Slice of arguments passed in the command line during script execution
*/
func handleSortingSlice(args []string) {
	param, order := args[2], args[3]

	switch param {
	case "name":
		if strings.ToLower(order) == "asc" {
			directors.SortAscendingByName()
		} else if strings.ToLower(order) == "desc" {
			directors.SortDescendingByName()
		} else {
			fmt.Println("Invalid sorting option. Accepted options: asc, desc")
		}
	case "rating":
		if strings.ToLower(order) == "asc" {
			directors.SortAscendingByAverage()
		} else if strings.ToLower(order) == "desc" {
			directors.SortDescendingByAverage()
		} else {
			fmt.Println("Invalid sorting option. Accepted options: asc, desc")
		}
	default:
		fmt.Println("Invalid parameter. Accepted parameters: name, rating")
	}
}

// Prints the list of accepted commands
func printHelp() {
	fmt.Printf(`Movie Rating Analyzer (version %s)
Available commands:
* -h | --help    Prints the list of available commands
* -v | --version Prints the version of the script
* -o | --order [param] [order] Prints the ratings but sorted by [param] (accepted values: name, rating) in [order] (accepted values: ASC, DESC)

If the script is run without additional arguments, it will print the average rating for each director
without as it finds them in the spreadsheet.`, version)
}

// Prints the current version of the script
func printVersion() {
	fmt.Printf("Movie Rating Analyzer version %s\n", version)
}
