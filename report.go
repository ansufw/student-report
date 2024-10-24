package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"text/tabwriter"
)

type report struct {
	name     string
	subjects []*subject
	students []*student
}

type subject struct {
	name string
}

type student struct {
	name   string
	scores map[*subject]uint32
}

func newReport(name string) *report {
	rpt := &report{
		name: name,
	}
	rpt.students = make([]*student, 0)
	rpt.subjects = make([]*subject, 0)

	return rpt
}

func newSubject(name string) *subject {
	sub := &subject{
		name: name,
	}

	return sub
}

func newStudent(name string) *student {
	student := &student{
		name:   name,
		scores: make(map[*subject]uint32, 0),
	}

	return student
}

func (r *report) addSubject(sub *subject) {
	r.subjects = append(r.subjects, sub)
}

func (r *report) addStudent(student *student) {
	r.students = append(r.students, student)
}

func (s *student) addScore(sub *subject, score uint32) {
	s.scores[sub] = score
}

func (s *student) getAverage() float32 {
	var total uint32
	for _, score := range s.scores {
		total += score
	}
	return float32(total) / float32(len(s.scores))
}

func (r *report) generateTable() {

	w := tabwriter.NewWriter(os.Stdout, 8, 8, 2, '\t', tabwriter.Debug)

	rowFormat := "%s"
	for range r.subjects {
		rowFormat += "\t%s"
	}
	headerFormat := rowFormat + "\t%s\t%s\n"
	rowFormat += "\t%.2f\t%s\n"

	headers := []interface{}{"student_name"}
	for _, subject := range r.subjects {
		headers = append(headers, subject.name)
	}
	headers = append(headers, "average", "grade")

	fmt.Fprintf(w, headerFormat, headers...)

	for _, student := range r.students {

		row := []interface{}{student.name}

		for _, subject := range r.subjects {
			row = append(row, fmt.Sprintf("%v", student.scores[subject]))
		}

		avg := student.getAverage()
		row = append(row, avg, getGrade(avg))

		fmt.Fprintf(w, rowFormat, row...)
	}

	w.Flush()
}

func (r *report) tranformTo2DSlice() [][]string {
	numRows := len(r.students) + 1
	result := make([][]string, numRows)

	// add header row
	result[0] = make([]string, 0)
	result[0] = append(result[0], "student_name")
	for _, sub := range r.subjects {
		result[0] = append(result[0], sub.name)
	}
	result[0] = append(result[0], "average", "grade")

	// add data rows
	for i, student := range r.students {
		result[i+1] = make([]string, 0)
		result[i+1] = append(result[i+1], student.name)
		for _, sub := range r.subjects {
			result[i+1] = append(result[i+1], fmt.Sprintf("%v", student.scores[sub]))
		}
		average := student.getAverage()
		result[i+1] = append(result[i+1], fmt.Sprintf("%.2f", average), getGrade(average))
	}

	return result
}

func (r *report) saveAsCSV() error {

	data := r.tranformTo2DSlice()
	file, _ := os.Create(fmt.Sprintf("%s.csv", r.name))
	w := csv.NewWriter(file)

	if err := w.WriteAll(data); err != nil {
		return err
	}

	return nil
}
