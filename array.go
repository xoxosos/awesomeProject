package main

import "fmt"

func main() {
	var bookings = [50]string{"aa", "bb", "cc"}
	bookings[3] = "dd"
	fmt.Println(bookings)
}
