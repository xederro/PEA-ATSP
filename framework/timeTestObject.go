// Package framework Description: This file contains the TimeTestObject struct and its methods.
package framework

import (
	"time"
)

// TimeTestObject is a struct that represents a single test.
type TimeTestObject struct {
	name          string
	before        func(size int) any
	measure       func(data any) any
	after         func(name string, nr int, testSize int, time time.Duration, data any)
	time          time.Duration
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
func (test *TimeTestObject) SetBefore(sb func(size int) any) *TimeTestObject {
	test.before = sb
	return test
}

// SetAfter sets the function to be called after each repetition.
func (test *TimeTestObject) SetAfter(sa func(name string, nr int, testSize int, time time.Duration, data any)) *TimeTestObject {
	test.after = sa
	return test
}

// SetMeasure sets the function that is meant to be measured.
func (test *TimeTestObject) SetMeasure(sm func(data any) any) *TimeTestObject {
	test.measure = sm
	return test
}

// Before runs before the test.
func Before(size int) any {
	// Do nothing before
	return nil
}

// Measure is the function to be measured
func Measure(data any) any {
	// Do nothing
	return nil
}

// After runs after the test.
func After(name string, nr int, testSize int, time time.Duration, data any) {
	// Do nothing
	return
}
