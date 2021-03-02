package BinaryTreeHashMap

import "Pair"

var INITIAL_CAPACITY = 1 << 5
var THRESHOLD_MAX float32 = 1
var THRESHOLD_MIN float32 = 0.2
var GROW_SHIFT = 2

type BinaryTreeHashMap struct {
	entries []*Node
	Size int
}

type Node struct {
	pair Pair.Pair
	left *Node
	right *Node
}

func createByCapacity(capacity int) BinaryTreeHashMap {
	return BinaryTreeHashMap{make([]*Node, capacity), 0}
}

func New() BinaryTreeHashMap {
	return createByCapacity(INITIAL_CAPACITY)
}

func (hm *BinaryTreeHashMap) Put(key uint64, value int) {
	hm.put(key, value, true)
}

func (hm *BinaryTreeHashMap) put(key uint64, value int, resize bool) {
	bucket := key % uint64(len(hm.entries))
	if hm.entries[bucket] == nil {
		hm.entries[bucket] = &Node{Pair.Pair{key, value}, nil, nil}
		hm.changeSizeBy(1, resize)
	} else {
		increment := hm.entries[bucket].putRec(Pair.Pair{key, value})
		if increment {
			hm.changeSizeBy(1, resize)
		}
	}
}

func (hm *BinaryTreeHashMap) Get(key uint64) (int, bool) {
	bucket := key % uint64(len(hm.entries))

	if hm.entries[bucket] == nil {
		return 0, false
	}

	node := hm.entries[bucket].getRec(key)

	if node != nil {
		return node.pair.Value, true
	} else {
		return 0, false
	}
}

func (hm *BinaryTreeHashMap) Remove(key uint64) {
	bucket := hm.entries[key % uint64(len(hm.entries))]
	if bucket == nil { return }

	if _, ok := hm.Get(key); ok {
		hm.changeSizeBy(-1, true)
	}
	hm.entries[key % uint64(len(hm.entries))] = removeRec(bucket, key)
}

func (hm *BinaryTreeHashMap) KeyValuePairs() []Pair.Pair {
	arr := make([]Pair.Pair, hm.Size)
	index := 0
	for i := 0; i < len(hm.entries); i++ {
		if hm.entries[i] != nil {
			keyValuePairsRec(hm.entries[i], arr, &index)
		}
	}

	return arr
}

func (node *Node) putRec(pair Pair.Pair) bool {
	if pair.Key < node.pair.Key {
		if node.left == nil {
			node.left = &Node{pair, nil, nil}
			return true
		} else {
			return node.left.putRec(pair)
		}
	} else if pair.Key > node.pair.Key {
		if node.right == nil {
			node.right = &Node{pair, nil, nil}
			return true
		} else {
			return node.right.putRec(pair)
		}
	} else {
		node.pair = pair
	}
	return false
}

func (node *Node) getRec(key uint64) *Node {
	if key == node.pair.Key {
		return node
	}
	if key < node.pair.Key {
		if node.left != nil {
			return node.left.getRec(key)
		} else {
			return nil
		}
	} else {
		if node.right != nil {
			return node.right.getRec(key)
		} else {
			return nil
		}
	}
}

func removeRec(node *Node, key uint64) *Node {
	if node == nil {
		return nil
	}

	if key < node.pair.Key {
		node.left = removeRec(node.left, key)
	} else if key > node.pair.Key {
		node.right = removeRec(node.right, key)
	} else {
		if node.left != nil && node.right != nil {
			minNodeRight := minNodeRec(node.right)
			node.pair = minNodeRight.pair
			node.right = removeRec(node.right, minNodeRight.pair.Key)
		} else if node.left != nil {
			node = node.left
		} else if node.right != nil {
			node = node.right
		} else {
			node = nil
		}
	}

	return node
}

func minNodeRec(node *Node) *Node {
	if node.left == nil {
		return node
	} else {
		return minNodeRec(node.left)
	}
}

func keyValuePairsRec(node *Node, arr []Pair.Pair, index *int) {
	if node != nil {
		keyValuePairsRec(node.left, arr, index)
		arr[*index] = node.pair
		*index += 1
		keyValuePairsRec(node.right, arr, index)
	}
}

func (hm *BinaryTreeHashMap) changeSizeBy(change int, resize bool) {
	hm.Size += change
	if resize {
		hm.resizeOnThreshold()
	}
}

func (hm *BinaryTreeHashMap) resizeOnThreshold() {
	newSize := 0
	if hm.Size < int(float32(len(hm.entries)) * THRESHOLD_MIN) && hm.Size > INITIAL_CAPACITY {
		newSize = len(hm.entries) >> 1
	} else if hm.Size > int(float32(len(hm.entries)) * THRESHOLD_MAX) {
		newSize = len(hm.entries) << GROW_SHIFT
	} else {
		return
	}

	newHashMap := createByCapacity(newSize)
	for _, pair := range hm.KeyValuePairs() {
		newHashMap.put(pair.Key, pair.Value, false)
	}

	*hm = newHashMap
}
