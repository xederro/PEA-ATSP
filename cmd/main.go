package main

import (
	"flag"
	"github.com/xederro/PEA-ATSP/tests"
	"github.com/xederro/PEA-ATSP/utils"
	"os"
	"runtime/pprof"
)

// main is the entry point of the program.
func main() {
	conf := tests.Config{
		RunBruteForce:     flag.Bool("bf", false, "Run bruteforce method"),
		RunBranchAndBound: flag.Bool("bab", false, "Run branch and bound method"),
		RunMemoization:    flag.Bool("m", false, "Run memoization method"),
		Repeat:            flag.Int("rep", -1, "How many times to repeat each iteration"),
		Concurrent:        flag.Bool("con", false, "Run concurrent method"),
	}
	perf := flag.Bool("cpu", false, "Measure performance")
	flag.Parse()
	ts := flag.Args()

	// If the -cpu flag is set, start the CPU profiler.
	if *perf {
		f, err := os.Create("cpu_profile.prof")
		if err != nil {
			panic(err)
		}
		defer f.Close()

		if err := pprof.StartCPUProfile(f); err != nil {
			panic(err)
		}
		defer pprof.StopCPUProfile()
	}

	// Parse the arguments.
	args, err := utils.ParseArgs(ts)
	if err != nil {
		os.Exit(1)
	}
	conf.Sizes = args

	// Run the tests.
	if conf.Repeat != nil && *conf.Repeat > 0 {
		conf.Run()
	} else {
		manual()
	}
}
