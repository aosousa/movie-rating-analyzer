package main

import "fmt"

// Handle any potential errors in the application
func handleError(err error) {
	if err != nil {
		fmt.Println(err)
		return
	}
}
