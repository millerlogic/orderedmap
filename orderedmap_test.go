package orderedmap

import (
	"math/rand"
	"testing"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func validate(t *testing.T, insertOrder []int16, m *OrderedMap) {
	i := 0
	m.Range(func(key KeyType, value ValueType) bool {
		if key.(int16) != insertOrder[i] {
			t.Fatalf("OrderedMap invalid, %d inserted key should be %d, got %d", i, insertOrder[i], key.(int16))
		}
		if value != ^insertOrder[i] {
			t.Fatal("Mismatched map value")
		}
		i++
		return true
	})
	if i != m.Len() {
		t.Fatalf("OrderedMap length mismatch, expected %d items, got %d items", len(insertOrder), i)
	}
}

func TestOrderedMap(t *testing.T) {
	m := NewOrderedMap()
	var insertOrder []int16

	for REPEAT := 0; REPEAT < 10; REPEAT++ {

		// Insert:
		for i := 0; i < 1000; i++ {
			n := int16(rand.Uint32())
			_, loaded := m.LoadOrStore(n, ^n)
			if !loaded {
				insertOrder = append(insertOrder, n)
			}
		}
		validate(t, insertOrder, m)

		// Delete randomly:
		for di := 0; di < 80; di++ {
			i := rand.Intn(len(insertOrder))
			n := insertOrder[i]
			insertOrder = append(insertOrder[:i], insertOrder[i+1:]...)
			m.Delete(n)
		}
		validate(t, insertOrder, m)

		// Delete from both ends:
		for di := 0; di < 10; di++ {
			i := 0
			n := insertOrder[i]
			insertOrder = insertOrder[1:]
			m.Delete(n)
		}
		for di := 0; di < 10; di++ {
			i := len(insertOrder) - 1
			n := insertOrder[i]
			insertOrder = insertOrder[:i]
			m.Delete(n)
		}
		validate(t, insertOrder, m)

		// Repeat some inserts (update):
		for i := 0; i < 30; i++ {
			n := insertOrder[i]
			m.Store(n, ^n)
		}
		validate(t, insertOrder, m)

		t.Logf("%d items in the map", m.Len())

	}

	// Single load:
	if _, ok := m.Load(insertOrder[0]); !ok {
		t.Fatal("Item not loaded.")
	}

	// Clear, load test:
	m.Clear()
	if m.Len() != 0 {
		t.Fatal("OrderedMap.Clear() failed, non-0 Len")
	}
	if _, ok := m.Load(insertOrder[0]); ok {
		t.Fatal("OrderedMap.Clear() failed, item still in map")
	}
	insertOrder = nil
	t.Logf("%d items in the map", m.Len())

}
