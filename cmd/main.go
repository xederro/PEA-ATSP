package main

import (
	"flag"
	"fmt"
	"github.com/xederro/PEA-ATSP/algo"
	"github.com/xederro/PEA-ATSP/algo/methods"
	"github.com/xederro/PEA-ATSP/framework"
	"os"
	"runtime/pprof"
	"time"
)

func main() {
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

	fmt.Println(ts)

	framework.NewTimeTestHarness(10, 12).AddTest(
		framework.NewTimeTestObject("Test", true, true).
			SetBefore(func(size int) any {
				//a := algo.NewIncidenceMatrixFromFile("D:\\projects\\PEA-ATSP\\tests\\tsp_12.txt")
				a := algo.NewIncidenceMatrix(size).Generate()
				fmt.Println(a.Stringify())
				bf := methods.NewBruteforce(a)
				return bf
			}).
			SetMeasure(func(data any) any {
				return data.(*methods.Bruteforce).Solve()
			}).
			SetAfter(func(name string, nr int, testSize int, time time.Duration, data any) {
				fmt.Println("Results:")
				fmt.Println(nr)
				fmt.Println(testSize)
				fmt.Println(time)
				fmt.Println(data.(*methods.Res).Value)
				fmt.Println(data.(*methods.Res).Route)
			}).SetTimeout(8 * time.Second),
	).Exec()
}
