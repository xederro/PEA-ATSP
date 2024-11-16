// Package framework Description: This file contains the TimeTestObject struct and its methods.
package framework

import (
	"github.com/xederro/PEA-ATSP/algo/methods"
	"time"
)

// TimeTestObject is a struct that represents a single test.
type TimeTestObject struct {
	name          string
	before        func(size int) methods.Method
	measure       func(data methods.Method) *methods.Res
	after         func(name string, nr int, testSize int, time time.Duration, data *methods.Res)
	print         bool
	timeout       time.Duration
	failOnTimeout bool
	failed        int
}

// NewTimeTestObject creates a new TimeTestObject object.
func NewTimeTestObject(name string, print bool, failOnTimeout bool) *TimeTestObject {
	return &TimeTestObject{
		name:          name,
		before:        Before,
		measure:       Measure,
		after:         After,
		print:         print,
		failOnTimeout: failOnTimeout,
	}
}

// SetTimeout sets the timeout duration
func (test *TimeTestObject) SetTimeout(timeout time.Duration) *TimeTestObject {
	test.timeout = timeout
	return test
}

// SetBefore sets the function to be called before each repetition.
func (test *TimeTestObject) SetBefore(sb func(size int) methods.Method) *TimeTestObject {
	test.before = sb
	return test
}

// SetAfter sets the function to be called after each repetition.
func (test *TimeTestObject) SetAfter(sa func(name string, nr int, testSize int, time time.Duration, data *methods.Res)) *TimeTestObject {
	test.after = sa
	return test
}

// SetMeasure sets the function that is meant to be measured.
func (test *TimeTestObject) SetMeasure(sm func(data methods.Method) *methods.Res) *TimeTestObject {
	test.measure = sm
	return test
}

// Before runs before the test.
func Before(size int) methods.Method {
	// Do nothing before
	return nil
}

// Measure is the function to be measured
func Measure(data methods.Method) *methods.Res {
	// Do nothing
	return nil
}

// After runs after the test.
func After(name string, nr int, testSize int, time time.Duration, data *methods.Res) {
	// Do nothing
	return
}
