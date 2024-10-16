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

func manual() {
	var im *algo.IncidenceMatrix
	for {
		action := -1
		allowed := []huh.Option[int]{
			huh.NewOption[int]("Generate a matrix", GENERATE),
			huh.NewOption[int]("Read a matrix from file", READ),
		}

		if im != nil {
			allowed = append(
				allowed,
				huh.NewOption[int]("Display matrix", WRITE),
				huh.NewOption[int]("Solve by bruteforce", BRUTEFORCE),
				huh.NewOption[int]("Solve by Little's algorithm", LITTLE),
			)
		}

		allowed = append(allowed, huh.NewOption[int]("Exit", EXIT))

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
						SetBefore(func(size int) any {
							return bruteforce.NewBruteforce(im)
						}).
						SetMeasure(func(data any) any {
							return data.(methods.Method).Solve()
						}).
						SetAfter(func(name string, nr int, testSize int, time time.Duration, data any) {
							fmt.Println("Results of bruteforce:")
							fmt.Println(time)
							fmt.Println(data.(*methods.Res).Value)
							fmt.Println(data.(*methods.Res).Route)
						}),
				).
				Exec()
			break
		case LITTLE:
			framework.NewTimeTestHarness(1, 0).
				AddTest(
					framework.NewTimeTestObject("Little", false, false).
						SetBefore(func(size int) any {
							return branchandbound.NewLittle(im)
						}).
						SetMeasure(func(data any) any {
							return data.(methods.Method).Solve()
						}).
						SetAfter(func(name string, nr int, testSize int, time time.Duration, data any) {
							fmt.Println("Results of Little's algorithm:")
							fmt.Println(time)
							fmt.Println(data.(*methods.Res).Value)
							fmt.Println(data.(*methods.Res).Route)
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
