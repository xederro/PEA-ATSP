// Package framework Description: This file contains the TTO struct and its methods.
package framework

import "time"

// TTO is a struct that represents a single test.
type TTO struct {
	name    string
	before  func(size int) any
	measure func(data any) any
	after   func(name string, nr int, testSize int, time time.Duration, data any)
	time    time.Duration
	print   bool
}

// NewTTO creates a new TTO object.
func NewTTO(name string, print bool) *TTO {
	return &TTO{
		name:    name,
		before:  Before,
		measure: Measure,
		after:   After,
		print:   print,
	}
}

// SetBefore sets the function to be called before each repetition.
func (test *TTO) SetBefore(sb func(size int) any) *TTO {
	test.before = sb
	return test
}

// SetAfter sets the function to be called after each repetition.
func (test *TTO) SetAfter(sa func(name string, nr int, testSize int, time time.Duration, data any)) *TTO {
	test.after = sa
	return test
}

// SetMeasure sets the function that is meant to be measured.
func (test *TTO) SetMeasure(sm func(data any) any) *TTO {
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
