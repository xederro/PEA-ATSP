package main

import (
	"flag"
	"github.com/xederro/PEA-ATSP/tests"
	"github.com/xederro/PEA-ATSP/utils"
	"os"
	"runtime/pprof"
)

func main() {
	conf := tests.Config{
		RunBruteForce:     flag.Bool("bf", false, "Run bruteforce method"),
		RunBranchAndBound: flag.Bool("bab", false, "Run branch and bound method"),
		Repeat:            flag.Int("rep", -1, "How many times to repeat each iteration"),
		Concurrent:        flag.Bool("con", false, "Run concurrent method"),
	}
	perf := flag.Bool("cpu", false, "Measure performance")
	flag.Parse()
	ts := flag.Args()

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

	args, err := utils.ParseArgs(ts)
	if err != nil {
		os.Exit(1)
	}
	conf.Sizes = args

	if conf.Repeat != nil && *conf.Repeat > 0 {
		conf.Run()
	} else {
		manual()
	}
}
