package main

import (
	"fmt"
	"github.com/xederro/PEA-ATSP/algo"
)

func main() {
	//perf := flag.Bool("cpu", false, "Measure performance")
	//
	//flag.Parse()
	//ts := flag.Args()
	//
	//if *perf {
	//	f, err := os.Create("cpu_profile.prof")
	//	if err != nil {
	//		panic(err)
	//	}
	//	defer f.Close()
	//
	//	if err := pprof.StartCPUProfile(f); err != nil {
	//		panic(err)
	//	}
	//	defer pprof.StopCPUProfile()
	//}
	//
	//fmt.Println(ts)

	a := algo.NewIncidenceMatrix(10)
	a.Generate()
	fmt.Println(a.Stringify())

	a = algo.NewIncidenceMatrixFromFile("D:\\projects\\PEA-ATSP\\tests\\tsp_10.txt")
	fmt.Println(a.Stringify())
}
