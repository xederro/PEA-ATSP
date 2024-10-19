// Package framework Description: This file contains the implementation of the TimeTestHarness witch is a simple test harness for measuring the time of a function.
package framework

import (
	"context"
	"fmt"
	"github.com/xederro/PEA-ATSP/algo/methods"
	"sync"
	"time"
)

// TimeTestHarness is a struct that represents a test harness.
type TimeTestHarness struct {
	repeat int
	sizes  []int
	tests  []*TimeTestObject
}

// NewTimeTestHarness creates a new TimeTestHarness with the given number of repetitions.
func NewTimeTestHarness(repeat int, testSize ...int) *TimeTestHarness {
	if len(testSize) == 0 {
		testSize = []int{0}
	}

	return &TimeTestHarness{
		repeat: repeat,
		sizes:  testSize,
	}
}

// AddTest adds a test to the test harness.
func (test *TimeTestHarness) AddTest(t *TimeTestObject) *TimeTestHarness {
	test.tests = append(test.tests, t)
	return test
}

// Exec executes the test harness.
func (test *TimeTestHarness) Exec() {
	for _, t := range test.tests {
		for _, size := range test.sizes {
			for i := range test.repeat {
				if t.failOnTimeout {
					d := t.before(size)
					out := make(chan *methods.Res, 1)
					ctxTimeout, cancel := context.WithTimeout(context.TODO(), t.timeout)
					start := time.Now()
					go func(ctx context.Context, ch chan *methods.Res) {
						ch <- t.measure(d)
					}(ctxTimeout, out)
					select {
					case res := <-out:
						m := res
						dur := time.Since(start)
						t.time += dur
						t.after(t.name, i, size, dur, m)
					case <-ctxTimeout.Done():
						t.failed++
					}
					cancel()
				} else {
					d := t.before(size)
					start := time.Now()
					m := t.measure(d)
					dur := time.Since(start)
					t.time += dur
					t.after(t.name, i, size, dur, m)
				}
			}

			if t.print {
				fmt.Printf("%s;%d;%d;%d;%d\n", t.name, size, t.failed, test.repeat, t.time.Nanoseconds()/int64(test.repeat-t.failed))
			}
		}
	}
}

// ExecWG executes the test harness with a wait group.
func (test *TimeTestHarness) ExecWG(wg *sync.WaitGroup) {
	test.Exec()
	if wg != nil {
		wg.Done()
	}
}
