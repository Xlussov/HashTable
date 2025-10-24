package hashtable

import (
	"fmt"
	"testing"
)

func TestPutAndGet(t *testing.T) {
	testKey := "somekey"
	testValue := 42

	ht := New[string, int](8)

	ht.Put(testKey, testValue)

	val, ok := ht.Get(testKey)

	if !ok {
		t.Fatalf("expected to find key '%s'", testKey)
	}
	if val != testValue {
		t.Errorf("expected value %d, got %v", testValue, val)
	}
}

func TestUpdateValue(t *testing.T) {
	ht := New[string, int](8)

	testKey := "somekey"
	testValue := 10
	testNewValue := 20

	ht.Put(testKey, testValue)
	ht.Put(testKey, testNewValue)

	val, _ := ht.Get(testKey)
	if val != testNewValue {
		t.Errorf("expected updated value %d, got %v", testNewValue, val)
	}
	if ht.Size() != 1 {
		t.Errorf("expected size 1 after overwrite, got %v", ht.Size())
	}
}

func TestRemove(t *testing.T) {
	testKey1 := "key1"
	testkey2 := "key2"

	testVal1 := 10
	testVal2 := 20

	ht := New[string, int](8)

	ht.Put(testKey1, testVal1)
	ht.Put(testkey2, testVal2)

	removed, ok := ht.Remove(testKey1)
	if !ok || removed != testVal1 {
		t.Errorf("expected removed value %d, got %v (ok=%v)", testVal1, removed, ok)
	}

	_, found := ht.Get(testKey1)
	if found {
		t.Errorf("expected key %s to be deleted", testKey1)
	}
}

func TestGetMissing(t *testing.T) {
	unexistKey := "key1"
	ht := New[string, int](8)
	_, ok := ht.Get(unexistKey)
	if ok {
		t.Errorf("expected ok=false for unexist key")
	}
}

func TestRehash(t *testing.T) {
	ht := New[string, int](8)
	initialCap := ht.Cap()

	for i := 0; i < 100; i++ {
		ht.Put(fmt.Sprintf("key%d", i), i)
	}

	if ht.Cap() <= initialCap {
		t.Errorf("expected capacity to increase after rehash")
	}

	for i := 0; i < 100; i++ {
		val, ok := ht.Get(fmt.Sprintf("key%d", i))
		if !ok || val != i {
			t.Errorf("expected key%d = %d, got %v (ok=%v)", i, i, val, ok)
		}
	}
}
