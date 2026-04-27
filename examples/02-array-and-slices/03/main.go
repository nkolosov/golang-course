package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"runtime/pprof"
)

type BigStruct struct {
	Array [1000]int64
}

type SmallStruct struct {
	Array [10]int64
}

// go run -gcflags "-m" main.go
func main() {
	f1, err := os.Create("mem.pprof")
	if err != nil {
		log.Fatal(err)
	}
	defer f1.Close()

	f2, err := os.Create("cpu.pprof")
	if err != nil {
		log.Fatal(err)
	}

	pprof.StartCPUProfile(f2)

	defer f2.Close()
	defer pprof.StopCPUProfile()

	tFoo()
	tBar()
}

func tFoo() {
	var a [10000]BigStruct

	println("tFoo")
	runtime.StartTrace()
	foo(a)
	fmt.Println(a[0].Array[0])
	runtime.StopTrace()
}

func tBar() {
	var b [10]SmallStruct

	println("tBar")
	runtime.StartTrace()
	bar(b)
	runtime.StopTrace()
	fmt.Println(b[0].Array[0])
}

//go:noinline
func foo(a [10000]BigStruct) {
	fmt.Println("test foo")
	a[0].Array[0] = 1
}

//go:noinline
func bar(a [10]SmallStruct) {
	fmt.Println("test bar")
	a[0].Array[0] = 1
}
