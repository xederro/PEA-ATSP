package main

import (
	"flag"
	"fmt"
	"os"
	"runtime/pprof"
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
	manual()

	//file := "D:\\projects\\PEA-ATSP\\tests\\tsp_4.txt"
	//framework.NewTimeTestHarness(1, 12).
	//	AddTest(
	//		framework.NewTimeTestObject("BruteForce", true, false).
	//			SetBefore(func(size int) any {
	//				a := algo.NewIncidenceMatrixFromFile(file)
	//				//a := algo.NewIncidenceMatrix(size).Generate()
	//				fmt.Println(a.Stringify())
	//				bf := bruteforce.NewBruteforce(a)
	//				return bf
	//			}).
	//			SetMeasure(func(data any) any {
	//				return data.(methods.Method).Solve()
	//			}).
	//			SetAfter(func(name string, nr int, testSize int, time time.Duration, data any) {
	//				fmt.Println("Results:")
	//				fmt.Println(nr)
	//				fmt.Println(testSize)
	//				fmt.Println(time)
	//				fmt.Println(data.(*methods.Res).Value)
	//				fmt.Println(data.(*methods.Res).Route)
	//			}).SetTimeout(8 * time.Second),
	//	).
	//	AddTest(
	//		framework.NewTimeTestObject("Little", true, false).
	//			SetBefore(func(size int) any {
	//				a := algo.NewIncidenceMatrixFromFile(file)
	//				//a := algo.NewIncidenceMatrix(size).Generate()
	//				fmt.Println(a.Stringify())
	//				l := branchandbound.NewLittle(a)
	//				return l
	//			}).
	//			SetMeasure(func(data any) any {
	//				return data.(methods.Method).Solve()
	//			}).
	//			SetAfter(func(name string, nr int, testSize int, time time.Duration, data any) {
	//				fmt.Println("Results:")
	//				fmt.Println(nr)
	//				fmt.Println(testSize)
	//				fmt.Println(time)
	//				fmt.Println(data.(*methods.Res).Value)
	//				fmt.Println(data.(*methods.Res).Route)
	//			}).SetTimeout(8 * time.Second),
	//	).
	//	Exec()
}
