package GoLangMap

import "Pair"

type GoLangMap struct {
	m map[uint64]int
}

func New() GoLangMap {
	return GoLangMap{
		map[uint64]int{},
	}
}

func (hm *GoLangMap) Put(key uint64, value int) {
	hm.m[key] = value
}

func (hm *GoLangMap) Get(key uint64) (int, bool) {
	if val, ok := hm.m[key]; ok {
		return val, true
	} else {
		return 0, false
	}
}

func (hm *GoLangMap) Remove(key uint64) {
	delete(hm.m, key)
}

func (hm *GoLangMap) KeyValuePairs() []Pair.Pair {
	arr := make([]Pair.Pair, len(hm.m))
	index := 0
	for key, val := range hm.m {
		arr[index] = Pair.Pair{key, val}
		index++
	}
	return arr
}