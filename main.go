package main 

import "fmt"

func main() {
	fmt.Println("Hello World!")

	integers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, integer := range integers {
		isEven := integer % 2

		if isEven == 0 {
			fmt.Printf("Nummber %v is Even \n", integer)
		} else {
			fmt.Printf("Nummber %v is Odd \n", integer)
		}
	}
}