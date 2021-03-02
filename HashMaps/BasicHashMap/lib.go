package BasicHashMap

import "Pair"

var INITIAL_CAPACITY = 1 << 5
var THRESHOLD_MAX float32 = 1.0
var THRESHOLD_MIN float32 = 0.25

type Entry struct {
	Key uint64
	Value int
	Next *Entry
	Previous *Entry
}

type BasicHashMap struct {
	entries []*Entry
	Size int
}

func createByCapacity(capacity int) BasicHashMap {
	return BasicHashMap{make([]*Entry, capacity), 0}
}

func New() BasicHashMap {
	return createByCapacity(INITIAL_CAPACITY)
}

func (hm *BasicHashMap) Put(key uint64, value int) {
	hm.put(key, value, true)
}

func (hm *BasicHashMap) put(key uint64, value int, resize bool) {
	entry := Entry{Key: key, Value: value}
	bucket := key % uint64(len(hm.entries))
	if hm.entries[bucket] == nil {
		hm.entries[bucket] = &entry
	} else {
		loopEntry := hm.entries[bucket]
		for loopEntry.Next != nil && loopEntry.Key != key {
			loopEntry = loopEntry.Next
		}

		if loopEntry.Key == key {
			loopEntry.Value = value
		} else {
			entry.Previous = loopEntry
			loopEntry.Next = &entry
		}
	}
	hm.changeSizeBy(+1, resize)
}

func (hm *BasicHashMap) Get(key uint64) (int, bool) {
	loopEntry := hm.entries[key % uint64(len(hm.entries))]

	for loopEntry != nil {
		if loopEntry.Key == key {
			return loopEntry.Value, true
		}
		loopEntry = loopEntry.Next
	}

	return 0, false
}

func (hm *BasicHashMap) Remove(key uint64) {
	bucket := hm.entries[key % uint64(len(hm.entries))]

	if bucket == nil { return }
	if bucket.Key == key {
		if bucket.Next == nil {
			hm.entries[key % uint64(len(hm.entries))] = nil
		} else {
			bucket.Next.Previous = nil
			hm.entries[key % uint64(len(hm.entries))] = bucket.Next
		}
		hm.changeSizeBy(-1, true)
	} else {
		for bucket.Next != nil {
			bucket = bucket.Next
			if bucket.Key == key {
				if bucket.Next != nil {
					bucket.Previous.Next = bucket.Next
					bucket.Next.Previous = bucket.Previous
				} else {
					bucket.Previous.Next = nil
				}
				hm.changeSizeBy(-1, true)
			}
		}
	}
}

func (hm *BasicHashMap) KeyValuePairs() []Pair.Pair {
	pairs := make([]Pair.Pair, hm.Size)
	index := 0

	for _, element := range hm.entries {
		if element == nil {
			continue
		}

		loopElement := element
		for true {
			pairs[index] = Pair.Pair{Key: loopElement.Key, Value: loopElement.Value}
			index += 1
			if loopElement.Next != nil {
				loopElement = loopElement.Next
			} else {
				break
			}
		}
	}

	return pairs
}

func (hm *BasicHashMap) changeSizeBy(change int, resize bool) {
	hm.Size += change
	if resize {
		hm.resizeOnThreshold()
	}
}

func (hm *BasicHashMap) resizeOnThreshold() {
	newSize := 0
	if hm.Size < int(float32(len(hm.entries)) * THRESHOLD_MIN) && hm.Size > INITIAL_CAPACITY {
		newSize = len(hm.entries) >> 1
	} else if hm.Size > int(float32(len(hm.entries)) * THRESHOLD_MAX) {
		newSize = len(hm.entries) << 1
	} else {
		return
	}

	newHashMap := createByCapacity(newSize)
	for _, pair := range hm.KeyValuePairs() {
		newHashMap.put(pair.Key, pair.Value, false)
	}

	*hm = newHashMap
}
