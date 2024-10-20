package tests

import (
	"github.com/xederro/PEA-ATSP/algo"
	"github.com/xederro/PEA-ATSP/algo/methods"
	"github.com/xederro/PEA-ATSP/algo/methods/branchandbound"
	"github.com/xederro/PEA-ATSP/algo/methods/bruteforce"
	"github.com/xederro/PEA-ATSP/framework"
	"sync"
	"time"
)

type Config struct {
	RunBruteForce     *bool
	RunBranchAndBound *bool
	Sizes             []int
	Repeat            *int
	Concurrent        *bool
}

func (c Config) Run() {
	if c.Concurrent != nil && *c.Concurrent == true {
		wg := &sync.WaitGroup{}
		if c.RunBruteForce != nil && *c.RunBruteForce == true {
			wg.Add(1)
			go c.runBruteForce(wg, *c.Repeat, c.Sizes...)
		}
		if c.RunBranchAndBound != nil && *c.RunBranchAndBound == true {
			wg.Add(1)
			go c.runBranchAndBound(wg, *c.Repeat, c.Sizes...)
		}
		wg.Wait()
	} else {
		if c.RunBruteForce != nil && *c.RunBruteForce == true {
			go c.runBruteForce(nil, *c.Repeat, c.Sizes...)
		}
		if c.RunBranchAndBound != nil && *c.RunBranchAndBound == true {
			go c.runBranchAndBound(nil, *c.Repeat, c.Sizes...)
		}
	}
}

func (c Config) runBruteForce(wg *sync.WaitGroup, count int, sizes ...int) {
	framework.NewTimeTestHarness(count, sizes...).
		AddTest(
			framework.NewTimeTestObject("BruteForce", true, false).
				SetBefore(func(size int) methods.Method {
					return bruteforce.NewBruteforce(
						algo.NewIncidenceMatrix(size).
							Generate(),
					)
				}).
				SetMeasure(func(data methods.Method) *methods.Res {
					return data.Solve()
				}),
		).
		ExecWG(wg)
}

func (c Config) runBranchAndBound(wg *sync.WaitGroup, count int, sizes ...int) {
	framework.NewTimeTestHarness(count, sizes...).
		AddTest(
			framework.NewTimeTestObject("BranchAndBound", true, true).
				SetBefore(func(size int) methods.Method {
					return branchandbound.NewLittle(
						algo.NewIncidenceMatrix(size).
							Generate(),
					)
				}).
				SetMeasure(func(data methods.Method) *methods.Res {
					return data.Solve()
				}).
				SetTimeout(5 * time.Minute),
		).
		ExecWG(wg)
}
