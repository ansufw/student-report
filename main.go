package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func createReport(r *bufio.Reader) *report {

	reportName, _ := getInput("Enter report name: ", r)

	return newReport(reportName)

}

func createSubjects(r *bufio.Reader, rpt *report) {
	subjects, _ := getArrayInput("\nNames of subjects (separate with commas ',', e.g. 'Maths,English,Bahasa Indonesia'):", r)
	for _, subject := range subjects {
		rpt.addSubject(newSubject(subject))
	}
}

func createStudents(r *bufio.Reader, rpt *report) {
	students, _ := getArrayInput("\nInsert of student(s). if more than one student, separate with comma ',', (e.g. 'John Doe,Michael Jackson,Jan Emily'):", r)
	for _, student := range students {
		rpt.addStudent(newStudent(student))
	}
}

func createScores(r *bufio.Reader, rpt *report) {
	for _, student := range rpt.students {
	A:
		scores, err := getUint32Input(fmt.Sprintf("\nInput score (0-100) of %s for %s (put in order and separate with comma ',', e.g. 70,55,80):", student.name, subjectsToString(rpt.subjects)), r)
		if err != nil {
			fmt.Println("invalid input:", err)
			fmt.Println("please check your input")
			goto A
		}
		if len(scores) != len(rpt.subjects) {
			fmt.Println("length of scores must be equal to length of subjects")
			fmt.Println("please check your input")
			goto A
		}

		for _, score := range scores {
			if score > 100 {
				fmt.Println("score must be less than 100")
				fmt.Println("please check your input")
				goto A
			}
		}
		for id, subject := range rpt.subjects {
			student.addScore(subject, scores[id])
		}
	}
}

func getOpt(reader *bufio.Reader, rpt *report) {

	menuOpt :=
		`Choose an option (input a number):
		1. Create new report
		2. Add student(s)
		3. Generate table report
		4. Save as csv
		0. Exit`

A:
	n, err := getAnumber(reader, menuOpt, 0, 1, 2, 3, 4)
	if err != nil {
		fmt.Println("invalid input:", err)
		fmt.Println("please try again")
		goto A
	}

	switch n {
	case 1:
		if rpt != nil {
			fmt.Println("a report already exists")
		B:
			yes, err := getYesOrNo(reader, "do you want to overwrite it? (y/n)")
			if err != nil {
				fmt.Println("invalid input:", err)
				fmt.Println("please check your input")
				time.Sleep(1 * time.Second)
				goto B
			}
			if !yes {
				fmt.Println("back to main menu...")
				time.Sleep(2 * time.Second)
				goto A
			}
		}
		// create new report
		report := createReport(reader)
		createSubjects(reader, report)
		createStudents(reader, report)
		createScores(reader, report)
		fmt.Println("report created successfully")
		time.Sleep(2 * time.Second)
		getOpt(reader, report)
	case 2:
		if rpt == nil {
			fmt.Println("no report exists")
			fmt.Println("please create a new report")
			time.Sleep(2 * time.Second)
			goto A
		}
		// add student(s)
		createStudents(reader, rpt)
		createScores(reader, rpt)
		fmt.Println("student(s) added successfully")
		time.Sleep(2 * time.Second)
		getOpt(reader, rpt)
	case 3:
		if rpt == nil {
			fmt.Println("no report exists")
			fmt.Println("please create a new report")
			time.Sleep(2 * time.Second)
			goto A
		}
		// generate table
		rpt.generateTable()
		getOpt(reader, rpt)
	case 4:
		if rpt == nil {
			fmt.Println("no report exists")
			fmt.Println("please create a new report")
			time.Sleep(2 * time.Second)
			goto A
		}
		// save as csv
		err := rpt.saveAsCSV()
		if err != nil {
			fmt.Println("an error occured:", err)
			fmt.Println("please try again")
			time.Sleep(2 * time.Second)
		} else {
			fmt.Println("csv created successfully")
			time.Sleep(2 * time.Second)
			fmt.Println("back to main menu...")
			time.Sleep(2 * time.Second)
		}
		getOpt(reader, rpt)

	case 0:
		fmt.Println("Goodbye!")
		time.Sleep(2 * time.Second)
		os.Exit(0)
	default:
		fmt.Println("invalid input")
	}

}

func main() {

	welcomeScreen()

	reader := bufio.NewReader(os.Stdin)
	// report := createReport(reader)
	// createSubjects(reader, report)
	// createStudents(reader, report)
	// createScores(reader, report)
	getOpt(reader, nil)

}
