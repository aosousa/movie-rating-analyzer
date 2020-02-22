package maths

/*SliceSum adds the ratings in a director's ratings slice

Receives:
	* ratings ([]int) - Ratings for a given director's movies

Returns:
	* int - Sum of the ratings for a given director's movies
*/
func SliceSum(ratings []int) int {
	var total int

	for _, rating := range ratings {
		total += rating
	}

	return total
}

/*Average returns the average rating in a slice of ratings for a given director

Receives:
	* ratings ([]int) - Ratings for a given director's movies

Returns:
	* float32 - Average rating for a given director's movies
*/
func Average(ratings []int) float32 {
	sum := SliceSum(ratings)

	return float32(sum) / float32(len(ratings))
}
