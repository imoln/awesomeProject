package main

import "fmt"

func gradeFunc(avg float64) rune {
	if avg >= 90 {
		return 'A'
	} else if avg >= 80 {
		return 'B'
	} else if avg >= 70 {
		return 'C'
	} else if avg >= 60 {
		return 'D'
	}
	return 'F'
}

func main() {
	var count int
	fmt.Println("How many grades? (1-5): ")
	fmt.Scan(&count)

	if count < 1 || count > 5 {
		fmt.Println("Please enter a value between 1 and 5")
		return
	}

	var sum int = 0

	for i := 1; i <= count; i++ {
		fmt.Println("Enter grade", i, ": ")
		var grade int
		fmt.Scan(&grade)
		sum += grade
	}
	avg := (float64(sum) / float64(count))

	fmt.Println("Average grade: ", avg)
	fmt.Printf("Letter grade: %c\n", gradeFunc(avg))

}
