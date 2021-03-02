package CuckooHashMap

import "Pair"

var INITIAL_CAPACITY = 1 << 6
var THRESHOLD_MAX float32 = 0.1
var THRESHOLD_MIN float32 = 0.1

type CuckooHashMap struct {
	entries1 []*Pair.Pair
	entries2 []*Pair.Pair
	Size int
}

/*
	Hash functions taken from https://github.com/rahul1947/SP07-Comparison-of-Hashing-Implementations/blob/master/Cuckoo.java
	MIT licensed

	Bit shifts are extracted from Java HashMap implementation to minimize collisions
 */
func (hm *CuckooHashMap) hashFunc1(key uint64) uint64 {
	return (key ^ (key >> 20) ^ (key >> 12) ^ (key >> 7) ^ (key >> 4)) & uint64(len(hm.entries1) - 1)
}

func (hm *CuckooHashMap) hashFunc2(key uint64) uint64 {
	return (hm.hashFunc1(key) + 2 * (1 + key % 9)) & uint64(len(hm.entries1) - 1)
}

func createByCapacity(capacity int) CuckooHashMap {
	return CuckooHashMap{make([]*Pair.Pair, capacity >> 1), make([]*Pair.Pair, capacity >> 1), 0}
}

func New() CuckooHashMap {
	return createByCapacity(INITIAL_CAPACITY)
}

func (hm *CuckooHashMap) Put(key uint64, value int) {
	if existing := hm.get(key); existing != nil {
		existing.Value = value
		return
	}

	pair := &Pair.Pair{key, value}
	location := hm.hashFunc1(pair.Key)
	entriesTable := &hm.entries1

	for i := 0; i < 5; i++ {
		if (*entriesTable)[location] == nil {
			(*entriesTable)[location] = pair
			hm.Size++
			return
		}

		pair, (*entriesTable)[location] = (*entriesTable)[location], pair

		if entriesTable == &hm.entries1 {
			entriesTable = &hm.entries2
			location = hm.hashFunc2(pair.Key)
		} else {
			entriesTable = &hm.entries1
			location = hm.hashFunc1(pair.Key)
		}
	}

	hm.grow()
	hm.Put(pair.Key, pair.Value)
}

func (hm *CuckooHashMap) Get(key uint64) (int, bool) {
	pair := hm.get(key)
	if pair == nil {
		return 0, false
	} else {
		return pair.Value, true
	}
}

func (hm *CuckooHashMap) Remove(key uint64) {
	location := hm.hashFunc1(key)
	if hm.entries1[location] != nil && hm.entries1[location].Key == key {
		hm.entries1[location] = nil
		hm.Size--
	}
	location = hm.hashFunc2(key)
	if hm.entries2[location] != nil && hm.entries2[location].Key == key {
		hm.entries2[location] = nil
		hm.Size--
	}
}

func (hm *CuckooHashMap) KeyValuePairs() ([]Pair.Pair) {
	pairs := make([]Pair.Pair, hm.Size)
	index := 0

	for i := 0; i < len(hm.entries1); i++ {
		if hm.entries1[i] != nil {
			pairs[index] = Pair.Pair{Key: hm.entries1[i].Key, Value: hm.entries1[i].Value}
			index += 1
		}
	}
	for i := 0; i < len(hm.entries2); i++ {
		if hm.entries2[i] != nil {
			pairs[index] = Pair.Pair{Key: hm.entries2[i].Key, Value: hm.entries2[i].Value}
			index += 1
		}
	}

	return pairs
}

func (hm *CuckooHashMap) get(key uint64) *Pair.Pair {
	location := hm.hashFunc1(key)
	if hm.entries1[location] != nil && hm.entries1[location].Key == key {
		return hm.entries1[location]
	}
	location = hm.hashFunc2(key)
	if hm.entries2[location] != nil && hm.entries2[location].Key == key {
		return hm.entries2[location]
	}

	return nil
}

func (hm *CuckooHashMap) grow() {
	newHashMap := createByCapacity(len(hm.entries1) << 2)

	for _, pair := range hm.KeyValuePairs() {
		newHashMap.Put(pair.Key, pair.Value)
	}

	*hm = newHashMap
}
