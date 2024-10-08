package utils

import (
	"github.com/charmbracelet/huh"
	"log"
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
				Value(&path),
		),
	).WithTheme(huh.ThemeCharm())

	err := form.Run()
	if err != nil {
		log.Fatalln("Error with form")
	}
	return path
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
