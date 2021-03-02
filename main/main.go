package main

import (
	"BasicHashMap"
	"BinaryTreeHashMap"
	"CuckooHashMap"
	"GoLangMap"
	"Pair"
	"flag"
	"fmt"
	"runtime"
	"time"
)

type HashMap interface {
	Put(key uint64, value int)
	Get(key uint64) (int, bool)
	Remove(key uint64)
	KeyValuePairs() []Pair.Pair
}

func main() {
	hashmapFlag := flag.String("hashmap", "all", "Hashmap to profile (all by default)")
	callsMultiplier := flag.Int("multiplier", 100, "Multiplier for get calls")
	elementsMultiplier := flag.Int("elements", 1048575, "Elements to insert (MUST BE PRIME NUMBER!)")
	flag.Parse()

	switch *hashmapFlag {
	case "all":
		benchAll(*elementsMultiplier, *callsMultiplier)
	case "BasicHashMap":
		hm := BasicHashMap.New()
		bench(&hm, *elementsMultiplier, *callsMultiplier, true)
	case "CuckooHashMap":
		hm := CuckooHashMap.New()
		bench(&hm, *elementsMultiplier, *callsMultiplier, true)
	case "BinaryTreeHashMap":
		hm := BinaryTreeHashMap.New()
		bench(&hm, *elementsMultiplier, *callsMultiplier, true)
	case "GoLangMap":
		hm := GoLangMap.New()
		bench(&hm, *elementsMultiplier, *callsMultiplier, true)
	default:
		benchAll(*elementsMultiplier, *callsMultiplier)
	}
}

func benchAll(elements int, multiplier int) {
	bhm := BasicHashMap.New()
	chm := CuckooHashMap.New()
	bth := BinaryTreeHashMap.New()
	mth := GoLangMap.New()
	testx := [...]HashMap{&chm, &bhm, &bth, &mth}
	for _, el := range testx {
		bench(el, elements, multiplier, false)
	}
}

func bench(hm HashMap, elements int, multiplier int, printMemory bool) {
	fmt.Printf("Benchmarking %T\n", hm)
	fmt.Printf("--------\n")

	/* MUST BE PRIME NUMBER!! */
	var primeMultiplier int
	if elements > 9000 {
		primeMultiplier = 6701
	} else if elements > 1001000 {
		primeMultiplier = 1000537
	} else {
		fmt.Printf("Please choose elements as > 9000")
		return
	}


	fmt.Printf("Config:\n")
	fmt.Printf("\telements: %d\n", elements)
	fmt.Printf("\tget calls: %d\n\n", elements * multiplier)

	start := time.Now()
	for i := 0; i < elements; i++ {
		key := uint64((primeMultiplier * i) % elements)
		hm.Put(key, i)
	}
	fmt.Printf("Insert: %s\n", time.Since(start))

	start = time.Now()
	for x := 0; x < multiplier; x++ {
		for y := 0; y < elements; y++ {
			key := uint64((primeMultiplier * y) % elements)
			hm.Get(key)
		}
	}
	fmt.Printf("Get: %s\n", time.Since(start))

	start = time.Now()
	for _, _ = range hm.KeyValuePairs() {}
	fmt.Printf("KeyValuePairs: %s\n", time.Since(start))

	if printMemory {
		var m runtime.MemStats
		runtime.ReadMemStats(&m)
		fmt.Printf("Alloc = %v MiB\n", bToMb(m.Alloc))
	}

	fmt.Printf("--------\n\n")
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}