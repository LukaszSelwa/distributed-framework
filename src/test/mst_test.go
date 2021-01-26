package test

import (
	"graphs/mst"
	"io/ioutil"
	"lib"
	"log"
	"testing"
)

func TestRunSynchGHSTriangle(t *testing.T) {
	adjencyList := [][][2]int{
		{{2, 10}, {3, 10}},
		{{1, 10}, {3, 20}},
		{{1, 10}, {2, 20}},
	}
	graph, synchronizer := lib.BuildSynchronizedWeightedGraphFromAdjencyList(adjencyList, lib.GetGenerator())
	if mst.RunSynchGHS(graph, synchronizer) == false {
		t.Fatalf("Verification failed\n")
	}
}

func TestRunSynchGHSSimplePath4(t *testing.T) {
	adjencyList := [][][2]int{
		{{2, 10}},
		{{1, 10}, {3, 20}},
		{{2, 20}, {4, 10}},
		{{3, 10}},
	}
	graph, synchronizer := lib.BuildSynchronizedWeightedGraphFromAdjencyList(adjencyList, lib.GetGenerator())
	if mst.RunSynchGHS(graph, synchronizer) == false {
		t.Fatalf("Verification failed\n")
	}
}

func TestRunSynchGHSSimplePath8(t *testing.T) {
	adjencyList := [][][2]int{
		{{2, 10}},
		{{1, 10}, {3, 20}},
		{{2, 20}, {4, 10}},
		{{3, 10}, {5, 30}},
		{{4, 30}, {6, 10}},
		{{5, 10}, {7, 20}},
		{{6, 20}, {8, 10}},
		{{7, 10}},
	}
	graph, synchronizer := lib.BuildSynchronizedWeightedGraphFromAdjencyList(adjencyList, lib.GetGenerator())
	if mst.RunSynchGHS(graph, synchronizer) == false {
		t.Fatalf("Verification failed\n")
	}
}

func TestRunSynchGHSWikipediaExample(t *testing.T) {
	adjencyList := [][][2]int{
		{{2, 4}, {3, 1}, {4, 4}},
		{{1, 4}, {3, 3}, {5, 10}, {10, 18}},
		{{1, 1}, {2, 3}, {4, 5}, {5, 9}},
		{{1, 4}, {3, 5}, {5, 7}, {6, 9}, {7, 9}},
		{{2, 10}, {3, 9}, {4, 7}, {7, 8}, {8, 9}, {10, 8}},
		{{4, 9}, {7, 2}, {8, 4}, {9, 6}},
		{{4, 9}, {5, 8}, {6, 2}, {8, 2}},
		{{5, 9}, {6, 4}, {7, 2}, {9, 3}, {10, 9}},
		{{6, 6}, {8, 3}, {10, 9}},
		{{2, 18}, {5, 8}, {8, 9}, {9, 9}},
	}
	graph, synchronizer := lib.BuildSynchronizedWeightedGraphFromAdjencyList(adjencyList, lib.GetGenerator())
	if mst.RunSynchGHS(graph, synchronizer) == false {
		t.Fatalf("Verification failed\n")
	}
}

func TestRunSynchGHSRandom(t *testing.T) {
	if mst.RunSynchGHSRandom(25, 100, 25) == false {
		t.Fatalf("Verification failed\n")
	}
}

func BenchmarkSynchGHSRandom(b *testing.B) {
	log.SetOutput(ioutil.Discard)
	for iteration := 0; iteration < b.N; iteration++ {
		mst.RunSynchGHSRandom(25, 100, 25)
	}
}