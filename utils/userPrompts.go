package utils

import (
	"errors"
	"fmt"
	"github.com/charmbracelet/huh"
	"log"
	"os"
	"strconv"
)

// GetPath is a function that allows user to choose file with values
func GetPath() string {
	var path string

	// insert path to file with values, and choose type of values
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Insert path to file with values").
				Prompt("> ").
				Value(&path).
				Validate(func(str string) error {
					fileInfo, err := os.Stat(str)
					if err != nil {
						return errors.New(fmt.Sprintf("%s is not a file", str))
					}

					if fileInfo.IsDir() {
						return errors.New(fmt.Sprintf("%s is not a file", str))
					}

					return nil
				}),
		),
	).WithTheme(huh.ThemeCharm())

	err := form.Run()
	if err != nil {
		log.Fatalln("Error with form")
	}
	return path
}

// GetSize is a function that allows user to choose size
func GetSize() int {
	var size string

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Insert size of an matrix").
				Prompt("> ").
				Value(&size).
				Validate(func(str string) error {
					_, err := strconv.Atoi(str)
					if err != nil {
						return errors.New("Cannot be represented as int")
					}
					return nil
				}),
		),
	).WithTheme(huh.ThemeCharm())

	err := form.Run()
	if err != nil {
		log.Fatalln("Error with form")
	}

	atoi, err := strconv.Atoi(size)
	if err != nil {
		log.Fatalln("Error with form")
	}

	return atoi
}

// ParseArgs is a function that parses string arguments to int
func ParseArgs(args *[]string) ([]int, error) {
	var result []int
	for _, arg := range *args {
		i, err := strconv.Atoi(arg)
		if err != nil {
			return nil, err
		}
		result = append(result, i)
	}
	return result, nil
}
