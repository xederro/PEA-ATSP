package tests

import (
	"github.com/xederro/PEA-ATSP/algo"
	"github.com/xederro/PEA-ATSP/algo/methods"
	"github.com/xederro/PEA-ATSP/algo/methods/branchandbound"
	"github.com/xederro/PEA-ATSP/algo/methods/bruteforce"
	"github.com/xederro/PEA-ATSP/algo/methods/memoization"
	"github.com/xederro/PEA-ATSP/framework"
	"sync"
	"time"
)

// Config is a struct that holds the configuration for the tests
type Config struct {
	RunBruteForce     *bool
	RunBranchAndBound *bool
	RunMemoization    *bool
	Sizes             []int
	Repeat            *int
	Concurrent        *bool
}

// Run runs the tests based on the configuration
func (c Config) Run() {
	if c.Concurrent != nil && *c.Concurrent == true {
		// tests are run concurrently
		wg := &sync.WaitGroup{}
		if c.RunBruteForce != nil && *c.RunBruteForce == true {
			wg.Add(1)
			go c.runBruteForce(wg, *c.Repeat, c.Sizes...)
		}
		if c.RunBranchAndBound != nil && *c.RunBranchAndBound == true {
			wg.Add(1)
			go c.runBranchAndBound(wg, *c.Repeat, c.Sizes...)
		}
		if c.RunMemoization != nil && *c.RunMemoization == true {
			wg.Add(1)
			go c.runMemoization(wg, *c.Repeat, c.Sizes...)
		}
		wg.Wait()
	} else {
		// tests are run sequentially
		if c.RunBruteForce != nil && *c.RunBruteForce == true {
			c.runBruteForce(nil, *c.Repeat, c.Sizes...)
		}
		if c.RunBranchAndBound != nil && *c.RunBranchAndBound == true {
			c.runBranchAndBound(nil, *c.Repeat, c.Sizes...)
		}
		if c.RunMemoization != nil && *c.RunMemoization == true {
			c.runMemoization(nil, *c.Repeat, c.Sizes...)
		}
	}
}

// runBruteForce runs the BruteForce tests
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
				}).
				SetTimeout(5 * time.Minute),
		).
		ExecWG(wg)
}

// runBranchAndBound runs the BranchAndBound tests
func (c Config) runBranchAndBound(wg *sync.WaitGroup, count int, sizes ...int) {
	framework.NewTimeTestHarness(count, sizes...).
		AddTest(
			framework.NewTimeTestObject("BranchAndBound", true, false).
				SetBefore(func(size int) methods.Method {
					return branchandbound.NewBranchAndBound(
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

// runMemoization runs the Memoization tests
func (c Config) runMemoization(wg *sync.WaitGroup, count int, sizes ...int) {
	framework.NewTimeTestHarness(count, sizes...).
		AddTest(
			framework.NewTimeTestObject("Memoization", true, false).
				SetBefore(func(size int) methods.Method {
					return memoization.NewMemoization(
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
