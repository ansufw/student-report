package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func welcomeScreen() {

	welcomeScreen1 := `
     _             _            _                               _
 ___| |_ _   _  __| | ___ _ __ | |_   _ __ ___ _ __   ___  _ __| |_
/ __| __| | | |/ _  |/ _ \  _ \| __| |  __/ _ \  _ \ / _ \|  __| __|
\__ \ |_| |_| | (_| |  __/ | | | |_  | | |  __/ |_) | (_) | |  | |_
|___/\__|\__,_|\__,_|\___|_| |_|\__| |_|  \___| .__/ \___/|_|   \__|
                                              |_|

----------------------------------------------------------
v0.0.1
by @ansuf(w)
----------------------------------------------------------

`

	fmt.Printf("%s", welcomeScreen1)
}

func getAnumber(r *bufio.Reader, prompt string, nums ...int) (int, error) {
	fmt.Println(prompt)
	input, err := r.ReadString('\n')
	if err != nil {
		return -1, err
	}

	n, err := strconv.Atoi(strings.TrimSpace(input))
	if err != nil {
		return -1, err
	}

	for _, num := range nums {
		if n == num {
			return n, nil
		}
	}

	return -1, fmt.Errorf("input must be one of %v", nums)
}

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Println(prompt)
	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err
}

func getArrayInput(prompt string, r *bufio.Reader) ([]string, error) {
	fmt.Println(prompt)
	input, err := r.ReadString('\n')

	return strings.Split(strings.TrimSpace(input), ","), err
}

func subjectsToString(subjects []*subject) string {
	subjectsStr := make([]string, 0)
	for _, sub := range subjects {
		subjectsStr = append(subjectsStr, sub.name)
	}
	return strings.Join(subjectsStr, ", ")
}

func getUint32Input(prompt string, r *bufio.Reader) ([]uint32, error) {
	fmt.Println(prompt)
	scores := make([]uint32, 0)

	input, err := r.ReadString('\n')
	if err != nil {
		return nil, err
	}

	numStrs := strings.Split(strings.TrimSpace(input), ",")

	for _, ns := range numStrs {
		nUint32, err := strconv.ParseUint(ns, 10, 32)
		if err != nil {
			return nil, err
		}

		scores = append(scores, uint32(nUint32))
	}

	return scores, nil
}

func getGrade(score float32) string {
	if score >= 80 {
		return "A"
	} else if score >= 70 {
		return "B"
	} else if score >= 60 {
		return "C"
	} else if score >= 50 {
		return "D"
	} else {
		return "E"
	}
}

func getYesOrNo(r *bufio.Reader, prompt string) (bool, error) {
	fmt.Println(prompt)
	input, err := r.ReadString('\n')
	if err != nil {
		return false, err
	}

	input = strings.TrimSpace(strings.ToLower(input))

	if input != "y" && input != "n" {
		return false, fmt.Errorf("input must be 'y' or 'n'")
	}

	return input == "y", nil
}
