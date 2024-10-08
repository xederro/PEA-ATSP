// Package framework Description: This file contains the implementation of the TimeTestHarness witch is a simple test harness for measuring the time of a function.
package framework

import (
	"fmt"
	"sync"
	"time"
)

// TTH is a struct that represents a test harness.
type TTH struct {
	repeat int
	sizes  []int
	tests  []*TTO
}

// NewTimeTestHarness creates a new TimeTestHarness with the given number of repetitions.
func NewTimeTestHarness(repeat int, testSize ...int) *TTH {
	if len(testSize) == 0 {
		testSize = []int{0}
	}

	return &TTH{
		repeat: repeat,
		sizes:  testSize,
	}
}

// AddTest adds a test to the test harness.
func (test *TTH) AddTest(t *TTO) *TTH {
	test.tests = append(test.tests, t)
	return test
}

// Exec executes the test harness.
func (test *TTH) Exec() {
	for _, t := range test.tests {
		for _, size := range test.sizes {
			for i := range test.repeat {
				d := t.before(size)
				start := time.Now()
				m := t.measure(d)
				dur := time.Since(start)
				t.time += dur
				t.after(t.name, i, size, dur, m)
			}

			if t.print {
				fmt.Printf("%s;%d;%d\n", t.name, size, t.time.Nanoseconds()/int64(test.repeat))
			}
		}
	}
}

// ExecWG executes the test harness with a wait group.
func (test *TTH) ExecWG(wg *sync.WaitGroup) {
	test.Exec()
	if wg != nil {
		wg.Done()
	}
}
