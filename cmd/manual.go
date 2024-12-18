package main

import (
	"fmt"
	"github.com/charmbracelet/huh"
	"github.com/xederro/PEA-ATSP/algo"
	"github.com/xederro/PEA-ATSP/algo/methods"
	"github.com/xederro/PEA-ATSP/algo/methods/branchandbound"
	"github.com/xederro/PEA-ATSP/algo/methods/bruteforce"
	"github.com/xederro/PEA-ATSP/algo/methods/memoization"
	"github.com/xederro/PEA-ATSP/framework"
	"github.com/xederro/PEA-ATSP/utils"
	"time"
)

const (
	GENERATE = iota
	READ
	WRITE
	BRUTEFORCE
	BRANCHANDBOUND
	MEMOIZATION
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
				huh.NewOption[int]("Solve by Brute Force", BRUTEFORCE),
				huh.NewOption[int]("Solve by Branch And Bound", BRANCHANDBOUND),
				huh.NewOption[int]("Solve by Dynamic Programming", MEMOIZATION),
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
					framework.NewTimeTestObject("Brute Force", false, false).
						SetBefore(func(size int) methods.Method {
							return bruteforce.NewBruteforce(im)
						}).
						SetMeasure(func(data methods.Method) *methods.Res {
							return data.Solve()
						}).
						SetAfter(func(name string, nr int, testSize int, time time.Duration, data *methods.Res) {
							fmt.Println("Results of Brute Force:")
							fmt.Println(time)
							fmt.Println(data.Value)
							fmt.Println(data.Route)
						}),
				).
				Exec()
			break
		case BRANCHANDBOUND:
			framework.NewTimeTestHarness(1, 0).
				AddTest(
					framework.NewTimeTestObject("Branch And Bound", false, false).
						SetBefore(func(size int) methods.Method {
							return branchandbound.NewBranchAndBound(im)
						}).
						SetMeasure(func(data methods.Method) *methods.Res {
							return data.Solve()
						}).
						SetAfter(func(name string, nr int, testSize int, time time.Duration, data *methods.Res) {
							fmt.Println("Results of Branch And Bound:")
							fmt.Println(time)
							fmt.Println(data.Value)
							fmt.Println(data.Route)
						}),
				).
				Exec()
			break
		case MEMOIZATION:
			framework.NewTimeTestHarness(1, 0).
				AddTest(
					framework.NewTimeTestObject("Dynamic Programming", false, false).
						SetBefore(func(size int) methods.Method {
							return memoization.NewMemoization(im)
						}).
						SetMeasure(func(data methods.Method) *methods.Res {
							return data.Solve()
						}).
						SetAfter(func(name string, nr int, testSize int, time time.Duration, data *methods.Res) {
							fmt.Println("Results of Dynamic Programming:")
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
