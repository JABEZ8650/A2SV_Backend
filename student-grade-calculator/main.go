package main

import (
	"fmt"
)

// calculateAverage takes a slice of float64 grades and returns the average.
func calculateAverage(grades []float64) float64 {
	var sum float64
	for _, grade := range grades {
		sum += grade
	}
	return sum / float64(len(grades))
}

func main() {
	var studentName string
	var subjectCount int

	// 1. Get student's name
	fmt.Print("Enter your name: ")
	fmt.Scanln(&studentName)

	// 2. Get number of subjects
	fmt.Print("Enter number of subjects: ")
	fmt.Scanln(&subjectCount)

	// Validate subject count
	if subjectCount <= 0 {
		fmt.Println("You must enter at least one subject.")
		return
	}

	// 3. Store subject names and grades
	subjects := make(map[string]float64)
	var grades []float64

	for i := 0; i < subjectCount; i++ {
		var subjectName string
		var grade float64

		fmt.Printf("Enter name for subject #%d: ", i+1)
		fmt.Scanln(&subjectName)

		// Input validation for grade
		for {
			fmt.Printf("Enter grade for %s (0-100): ", subjectName)
			fmt.Scanln(&grade)

			if grade >= 0 && grade <= 100 {
				break
			}
			fmt.Println("Invalid grade. Must be between 0 and 100.")
		}

		subjects[subjectName] = grade
		grades = append(grades, grade)
	}

	// 4. Calculate average
	average := calculateAverage(grades)

	// 5. Display results
	fmt.Printf("\nStudent Name: %s\n", studentName)
	fmt.Println("Grades:")
	for subject, grade := range subjects {
		fmt.Printf(" - %s: %.2f\n", subject, grade)
	}
	fmt.Printf("Average Grade: %.2f\n", average)
}
