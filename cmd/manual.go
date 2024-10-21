package main

import (
	"fmt"
	"github.com/charmbracelet/huh"
	"github.com/xederro/PEA-ATSP/algo"
	"github.com/xederro/PEA-ATSP/algo/methods"
	"github.com/xederro/PEA-ATSP/algo/methods/branchandbound"
	"github.com/xederro/PEA-ATSP/algo/methods/bruteforce"
	"github.com/xederro/PEA-ATSP/framework"
	"github.com/xederro/PEA-ATSP/utils"
	"time"
)

const (
	GENERATE = iota
	READ
	WRITE
	BRUTEFORCE
	LITTLE
	EXIT
)

// manual is a func that handles interaction with a user
func manual() {
	var im *algo.IncidenceMatrix
	for {
		action := -1
		// prepare actions
		allowed := []huh.Option[int]{
			huh.NewOption[int]("Generate a matrix", GENERATE),
			huh.NewOption[int]("Read a matrix from file", READ),
		}

		if im != nil {
			allowed = append(
				allowed,
				huh.NewOption[int]("Display matrix", WRITE),
				huh.NewOption[int]("Solve by bruteforce", BRUTEFORCE),
				huh.NewOption[int]("Solve by BranchAndBound's algorithm", LITTLE),
			)
		}

		allowed = append(allowed, huh.NewOption[int]("Exit", EXIT))
		//prompt for action
		err := huh.NewForm(
			huh.NewGroup(
				huh.NewSelect[int]().
					Title("Pick an action.").
					Options(allowed...).
					Value(&action),
			),
		).Run()
		if err != nil {
			panic("This should not happen")
		}
		// react on action
		switch action {
		case GENERATE:
			size := utils.GetSize()
			im = algo.NewIncidenceMatrix(size).Generate()
			break
		case READ:
			path := utils.GetPath()
			im = algo.NewIncidenceMatrixFromFile(path)
			break
		case WRITE:
			fmt.Println(im.Stringify())
			break
		case BRUTEFORCE:
			framework.NewTimeTestHarness(1, 0).
				AddTest(
					framework.NewTimeTestObject("BruteForce", false, false).
						SetBefore(func(size int) methods.Method {
							return bruteforce.NewBruteforce(im)
						}).
						SetMeasure(func(data methods.Method) *methods.Res {
							return data.Solve()
						}).
						SetAfter(func(name string, nr int, testSize int, time time.Duration, data *methods.Res) {
							fmt.Println("Results of bruteforce:")
							fmt.Println(time)
							fmt.Println(data.Value)
							fmt.Println(data.Route)
						}),
				).
				Exec()
			break
		case LITTLE:
			framework.NewTimeTestHarness(1, 0).
				AddTest(
					framework.NewTimeTestObject("BranchAndBound", false, false).
						SetBefore(func(size int) methods.Method {
							return branchandbound.NewBranchAndBound(im)
						}).
						SetMeasure(func(data methods.Method) *methods.Res {
							return data.Solve()
						}).
						SetAfter(func(name string, nr int, testSize int, time time.Duration, data *methods.Res) {
							fmt.Println("Results of BranchAndBound's algorithm:")
							fmt.Println(time)
							fmt.Println(data.Value)
							fmt.Println(data.Route)
						}),
				).
				Exec()
			break
		case EXIT:
			fmt.Println("Bye!")
			return
		default:
			panic("Should not happen")
		}
	}
}
