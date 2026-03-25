package cache

import (
	"testing"
)

func cacheToSlice[K comparable, V any](cache *LRUCache[K, V]) []V {
	list := cache.data
	node := list.Front()
	slice := make([]V, 0, list.Size())
	for node != list.dummy {
		slice = append(slice, node.val)
		node = node.next
	}
	return slice
}

type IntItem struct {
	key string
	val int
}

func TestLRUCacheInt(t *testing.T) {
	tests := []struct {
		name     string
		capacity int
		input    []IntItem
		expected []int
	}{
		{
			name:     "basic",
			capacity: 10,
			input:    []IntItem{{"a", 5}, {"b", 4}, {"c", 3}, {"d", 2}, {"e", 1}},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "capacity equal len",
			capacity: 5,
			input:    []IntItem{{"a", 5}, {"b", 4}, {"c", 3}, {"d", 2}, {"e", 1}},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "not enough capacity",
			capacity: 5,
			input:    []IntItem{{"a", 10}, {"b", 9}, {"c", 8}, {"d", 7}, {"e", 6}, {"f", 5}, {"g", 4}, {"h", 3}, {"i", 2}, {"j", 1}},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "repeated keys",
			capacity: 5,
			input:    []IntItem{{"a", 10}, {"b", 9}, {"c", 8}, {"d", 7}, {"e", 6}, {"f", 5}, {"g", 4}, {"h", 3}, {"i", 2}, {"j", 1}, {"b", 9}, {"a", 10}},
			expected: []int{10, 9, 1, 2, 3},
		},
		{
			name:     "one repeated key",
			capacity: 5,
			input:    []IntItem{{"a", 1}, {"a", 1}, {"a", 1}, {"a", 1}, {"a", 1}},
			expected: []int{1},
		},
		{
			name:     "many repetitions in the end",
			capacity: 5,
			input:    []IntItem{{"a", 5}, {"b", 4}, {"c", 3}, {"d", 2}, {"e", 1}, {"a", 5}, {"a", 5}, {"a", 5}},
			expected: []int{5, 1, 2, 3, 4},
		},
		{
			name:     "empty",
			capacity: 10,
			input:    []IntItem{},
			expected: []int{},
		},
	}
	for _, tt := range tests {
		cache, _ := NewLRUCache[string, int](tt.capacity)
		for _, item := range tt.input {
			cache.Set(item.key, item.val)
		}
		if cache.Size() != len(tt.expected) {
			t.Errorf("%s: After insertions: cache.Size() = %d. Expected: %d.\n", tt.name, cache.Size(), len(tt.expected))
		}
		slice := cacheToSlice(cache)
		for i := range slice {
			if slice[i] != tt.expected[i] {
				t.Errorf("%s: cache[%d] = %d. Expected: %d.\n", tt.name, i, slice[i], tt.expected[i])
			}
		}
		// for _, num := range tt.expected {
		// 	if _, ok := cache.Get(num); !ok {
		// 		t.Errorf("%s: %d not found in cache.\n", tt.name, num)
		// 	}
		// }
		cache.Clear()
		if cache.Size() != 0 {
			t.Errorf("%s: After cache.Clear(): cache.Size() = %d. Expected: 0.\n", tt.name, cache.Size())
		}
		if cache.data.Size() != 0 {
			t.Errorf("%s: After cache.Clear(): cache.data.Size() = %d. Expected: 0.\n", tt.name, cache.data.Size())
		}
		if len(cache.nodes) != 0 {
			t.Errorf("%s: After cache.Clear(): len(cache.nodes) = %d. Expected: 0.\n", tt.name, len(cache.nodes))
		}
	}
}

func TestLRUCacheString(t *testing.T) {
	tests := []struct {
		name     string
		capacity int
		input    []string
		expected []string
	}{
		{
			name:     "basic",
			capacity: 10,
			input:    []string{"Erlang", "Delphi", "C", "Basic", "Assembler"},
			expected: []string{"Assembler", "Basic", "C", "Delphi", "Erlang"},
		},
		{
			name:     "capacity equal len",
			capacity: 5,
			input:    []string{"Erlang", "Delphi", "C", "Basic", "Assembler"},
			expected: []string{"Assembler", "Basic", "C", "Delphi", "Erlang"},
		},
		{
			name:     "not enough capacity",
			capacity: 5,
			input:    []string{"Go", "Fortran", "Erlang", "Delphi", "C", "Basic", "Assembler"},
			expected: []string{"Assembler", "Basic", "C", "Delphi", "Erlang"},
		},
		{
			name:     "repeated strings",
			capacity: 5,
			input:    []string{"Go", "Fortran", "Erlang", "Delphi", "C", "Basic", "Assembler", "Fortran", "Go"},
			expected: []string{"Go", "Fortran", "Assembler", "Basic", "C"},
		},
		{
			name:     "one repeated element",
			capacity: 5,
			input:    []string{"Go", "Go", "Go", "Go", "Go", "Go"},
			expected: []string{"Go"},
		},
		{
			name:     "many repetitions in the end",
			capacity: 5,
			input:    []string{"Go", "Delphi", "C", "Basic", "Assembler", "Go", "Go", "Go", "Go"},
			expected: []string{"Go", "Assembler", "Basic", "C", "Delphi"},
		},
		{
			name:     "empty",
			capacity: 10,
			input:    []string{},
			expected: []string{},
		},
	}
	for _, tt := range tests {
		cache, _ := NewLRUCache[string, string](tt.capacity)
		for _, s := range tt.input {
			cache.Set(s, s)
		}
		if cache.Size() != len(tt.expected) {
			t.Errorf("%s: After insertions: cache.Size() = %d. Expected: %d.\n", tt.name, cache.Size(), len(tt.expected))
		}
		slice := cacheToSlice(cache)
		for i := range slice {
			if slice[i] != tt.expected[i] {
				t.Errorf("%s: cache[%d] = %s. Expected: %s.\n", tt.name, i, slice[i], tt.expected[i])
			}
		}
		for _, num := range tt.expected {
			if _, ok := cache.Get(num); !ok {
				t.Errorf("%s: %s not found in cache.\n", tt.name, num)
			}
		}
		cache.Clear()
		if cache.Size() != 0 {
			t.Errorf("%s: After cache.Clear(): cache.Size() = %d. Expected: 0.\n", tt.name, cache.Size())
		}
		if cache.data.Size() != 0 {
			t.Errorf("%s: After cache.Clear(): cache.data.Size() = %d. Expected: 0.\n", tt.name, cache.data.Size())
		}
		if len(cache.nodes) != 0 {
			t.Errorf("%s: After cache.Clear(): len(cache.nodes) = %d. Expected: 0.\n", tt.name, len(cache.nodes))
		}
	}
}

// func TestLRUCacheIntConcurrent(t *testing.T) {
// 	const numGoroutines = 100000
// 	cache, _ := NewLRUCache[int](numGoroutines)
// 	var wg sync.WaitGroup
// 	for i := range numGoroutines {
// 		wg.Add(1)
// 		go func(wg *sync.WaitGroup, num int) {
// 			cache.Set(num)
// 			wg.Done()
// 		}(&wg, i)
// 	}
// 	wg.Wait()
// 	if cache.Size() != numGoroutines {
// 		t.Errorf("cache.Size() = %d. Expected: %d.\n", cache.Size(), numGoroutines)
// 	}
// 	for range numGoroutines {
// 		if _, ok := cache.Get(numGoroutines); !ok {
// 			t.Errorf("%d not found in cache.\n", numGoroutines)
// 		}
// 	}
// }
