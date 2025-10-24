// package hashtable

// import (
// 	"fmt"
// 	"hash/fnv"
// )

// // TODO dynamic capacity 75% filling, implement rehash method, write tests

// const (
// 	SLOT_EMPTY   = 0
// 	SLOT_FILLED  = 1
// 	SLOT_DELETED = 2
// )

// type HashTable[K comparable, V any] struct {
// 	buckets  []pair[K, V]
// 	capacity uint
// 	size     uint
// }

// type pair[K comparable, V any] struct {
// 	key   K
// 	value V
// 	state uint8
// }

// func New[K comparable, V any](capacity uint) *HashTable[K, V] {
// 	if capacity <= 0 {
// 		capacity = 8
// 	}

// 	return &HashTable[K, V]{
// 		capacity: capacity,
// 		buckets:  make([]pair[K, V], capacity),
// 	}
// }

// func (ht *HashTable[K, V]) Put(key K, value V) {
// 	for i := range int(ht.capacity) {
// 		hash := ht.hashIndex(key, i)
// 		slot := &ht.buckets[hash]

// 		switch slot.state {
// 		case SLOT_EMPTY, SLOT_DELETED:
// 			slot.key = key
// 			slot.value = value
// 			slot.state = SLOT_FILLED
// 			ht.size += 1
// 			return

// 		case SLOT_FILLED:
// 			if slot.key == key {
// 				slot.value = value
// 				return
// 			}
// 		}

// 		// if slot.state == SLOT_FILLED && slot.key != key {
// 		// 	continue
// 		// }

// 		// slot.key = key
// 		// slot.value = value
// 		// slot.state = SLOT_FILLED
// 		// ht.size += 1
// 		// return
// 	}

// 	panic("table is full") // todo dynamic capacity
// }

// func (ht *HashTable[K, V]) Get(key K) (V, bool) {
// 	var zero V
// 	for i := range int(ht.capacity) {
// 		hash := ht.hashIndex(key, i)
// 		slot := &ht.buckets[hash]

// 		if slot.state == SLOT_FILLED && slot.key != key || slot.state == SLOT_DELETED {
// 			continue
// 		}

// 		return slot.value, true
// 	}

// 	return zero, false
// }

// func (ht *HashTable[K, V]) Remove(key K) (V, bool) {
// 	var emptyKey K
// 	var emptyValue V

// 	for i := range int(ht.capacity) {
// 		hash := ht.hashIndex(key, i)
// 		slot := &ht.buckets[hash]

// 		if slot.state == SLOT_FILLED && slot.key != key {
// 			continue
// 		}
// 		value := slot.value

// 		slot.key = emptyKey
// 		slot.value = emptyValue
// 		slot.state = SLOT_DELETED
// 		return value, true
// 	}

// 	return emptyValue, false
// }

// func (ht *HashTable[K, V]) Print() {
// 	fmt.Println(ht.buckets)
// }

// func (ht *HashTable[K, V]) Len() int {
// 	return int(ht.size)
// }

// func (ht *HashTable[K, V]) hashKey(key K) uint64 {
// 	h := fnv.New64a()
// 	h.Write(fmt.Appendf(nil, "%v", key))
// 	return h.Sum64()
// }

// func (ht *HashTable[K, V]) hashStep(key K) uint64 {
// 	h := fnv.New64a()
// 	h.Write(fmt.Appendf(nil, "%v", key))
// 	return (h.Sum64() % uint64(ht.capacity-1)) + 1
// }

// func (ht *HashTable[K, V]) hashIndex(key K, i int) uint64 {
// 	return (ht.hashKey(key) + uint64(i)*ht.hashStep(key)) % uint64(ht.capacity)
// }

package hashtable

import (
	"fmt"
	"hash/fnv"
)

const (
	SLOT_EMPTY   uint8 = 0
	SLOT_FILLED  uint8 = 1
	SLOT_DELETED uint8 = 2
)

type HashTable[K comparable, V any] struct {
	states   []uint8
	keys     []K
	values   []V
	capacity uint
	size     uint
}

func New[K comparable, V any](capacity uint) *HashTable[K, V] {
	if capacity <= 0 {
		capacity = 8
	}

	return &HashTable[K, V]{
		states:   make([]uint8, capacity),
		keys:     make([]K, capacity),
		values:   make([]V, capacity),
		capacity: capacity,
		size:     0,
	}
}

func (ht *HashTable[K, V]) Put(key K, value V) {
	if ht.isLoadFactor() {
		ht.rehash()
	}

	for i := range int(ht.capacity) {
		hash := ht.hashIndex(key, i)

		switch ht.states[hash] {
		case SLOT_EMPTY, SLOT_DELETED:
			ht.keys[hash] = key
			ht.values[hash] = value
			ht.states[hash] = SLOT_FILLED
			ht.size += 1
			return

		case SLOT_FILLED:
			if ht.keys[hash] == key {
				ht.values[hash] = value
				return
			}
		}
	}

}

func (ht *HashTable[K, V]) Get(key K) (V, bool) {
	var emptyValue V

	for i := range int(ht.capacity) {
		hash := ht.hashIndex(key, i)

		if ht.states[hash] == SLOT_FILLED && ht.keys[hash] != key || ht.states[hash] == SLOT_DELETED {
			continue
		}

		return ht.values[hash], true
	}

	return emptyValue, false
}

func (ht *HashTable[K, V]) Remove(key K) (V, bool) {
	var emptyKey K
	var emptyValue V

	for i := range int(ht.capacity) {
		hash := ht.hashIndex(key, i)

		if ht.states[hash] == SLOT_FILLED && ht.keys[hash] != key {
			continue
		}
		value := ht.values[hash]

		ht.keys[hash] = emptyKey
		ht.values[hash] = emptyValue
		ht.states[hash] = SLOT_DELETED
		ht.size -= 1
		return value, true
	}

	return emptyValue, false
}

func (ht *HashTable[K, V]) Print() {
	for i := uint(0); i < ht.capacity; i++ {

		fmt.Printf("Index %d: Key = %v, Value = %v\n", i, ht.keys[i], ht.values[i])
	}
}

func (ht *HashTable[K, V]) Len() int {
	return int(ht.size)
}

func (ht *HashTable[K, V]) rehash() {
	oldCapacity := ht.capacity
	oldStates := ht.states
	oldKeys := ht.keys
	oldValues := ht.values

	ht.capacity *= 2
	ht.states = make([]uint8, ht.capacity)
	ht.keys = make([]K, ht.capacity)
	ht.values = make([]V, ht.capacity)
	ht.size = 0

	for i := 0; i < int(oldCapacity); i++ {
		if oldStates[i] == SLOT_FILLED {
			ht.Put(oldKeys[i], oldValues[i])
		}
	}
}

func (ht *HashTable[K, V]) isLoadFactor() bool {
	return float64(ht.size)/float64(ht.capacity) > 0.75
}

func (ht *HashTable[K, V]) hashKey(key K) uint64 {
	h := fnv.New64a()
	h.Write(fmt.Appendf(nil, "%v", key))
	return h.Sum64()
}

func (ht *HashTable[K, V]) hashStep(key K) uint64 {
	h := fnv.New64a()
	h.Write(fmt.Appendf(nil, "%v", key))
	return (h.Sum64() % uint64(ht.capacity-1)) + 1
}

func (ht *HashTable[K, V]) hashIndex(key K, i int) uint64 {
	return (ht.hashKey(key) + uint64(i)*ht.hashStep(key)) % uint64(ht.capacity)
}
