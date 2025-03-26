package main

import (
	"bufio"
	"fmt"
	"os"
)

func average(grades map[string]float64) float64 {
	var sum float64
	var count int

	for _, mark := range grades {
		sum += mark
		count++
	}
	if count == 0 {
		return 0
	}

	return sum / float64(count)
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Welcome to Student grade calculator!")

	fmt.Print("Enter your name: ")
	studentName, _ := reader.ReadString('\n')
	studentName = studentName[:len(studentName)-1]

	fmt.Print("How many classes do you take: ")
	var subjectNumber int
	fmt.Scanln(&subjectNumber)

	marklist := make(map[string]float64)
	for i := 1; i <= subjectNumber; i++ {
		var subjectName string
		var grade float64

		fmt.Printf("Enter subject #%d name: ", i)
		fmt.Scanln(&subjectName)

		for {
			fmt.Printf("Enter %s grade(0-100): ", subjectName)
			fmt.Scanln(&grade)

			if grade <= 100 && grade >= 0 {
				break
			} else {
				println("Enter a valid number between 1 and 100")
			}
		}
		marklist[subjectName] = grade
	}

	fmt.Printf("\nGrade Summery\n")
	for subj, mark := range marklist {
		fmt.Printf("%s: %.2f\n", subj, mark)
	}

	ave_grade := average(marklist)
	fmt.Printf("%s's average grade is %.2f\n", studentName, ave_grade)
}
