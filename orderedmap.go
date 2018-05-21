// Copyright (c) 2018 Christopher E. Miller
// MIT license, see LICENSE file.

package orderedmap

// KeyType is the map key type, can be generically overridden.
type KeyType interface{}

// ValueType is the map value type, can be generically overridden.
type ValueType interface{}

type link struct {
	next, prev *link
	key        KeyType
	value      ValueType
}

// OrderedMap is an ordered map collection. Not thread safe, but exposes the same interface as sync.Map
type OrderedMap struct {
	lookup      map[KeyType]*link
	first, last *link
}

// NewOrderedMap makes a new OrderedMap
func NewOrderedMap() *OrderedMap {
	return &OrderedMap{lookup: make(map[KeyType]*link)}
}

// Delete deletes this key from the map
func (m *OrderedMap) Delete(key KeyType) {
	if link, exists := m.lookup[key]; exists {
		if link.prev != nil {
			link.prev.next = link.next
		}
		if link.next != nil {
			link.next.prev = link.prev
		}
		if link == m.first {
			m.first = link.next
		}
		if link == m.last {
			m.last = link.prev
		}
		delete(m.lookup, key)
	}
}

// Load loads this key from the map
func (m *OrderedMap) Load(key KeyType) (value ValueType, ok bool) {
	if link, exists := m.lookup[key]; exists {
		value = link.value
		ok = true
	}
	return
}

// LoadOrStore loads or stores
func (m *OrderedMap) LoadOrStore(key KeyType, value ValueType) (actual ValueType, loaded bool) {
	if link, exists := m.lookup[key]; exists {
		actual = link.value
		loaded = true
	} else {
		m.Store(key, value)
		actual = value
	}
	return
}

// Range iterates over the key/value pairs in order
func (m *OrderedMap) Range(f func(key KeyType, value ValueType) bool) {
	for link := m.first; link != nil; link = link.next {
		if !f(link.key, link.value) {
			break
		}
	}
}

// Store stores this key/value pair in the map
func (m *OrderedMap) Store(key KeyType, value ValueType) {
	if link, exists := m.lookup[key]; exists {
		// Exists, update...
		link.value = value
		return
	}
	// Does not exist, add it...
	link := &link{key: key, value: value, prev: m.last}
	if m.last != nil {
		m.last.next = link
	}
	m.last = link
	if m.first == nil {
		m.first = link
	}
	m.lookup[key] = link
}

// Clear clears the map back to an empty state of no keys.
func (m *OrderedMap) Clear() {
	m.first = nil
	m.last = nil
	m.lookup = make(map[KeyType]*link)
}

// Len returns the number of items in the map.
func (m *OrderedMap) Len() int {
	return len(m.lookup)
}
