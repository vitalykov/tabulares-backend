package cache

import "testing"

func listToSlice[T any](list *List[T]) []T {
	slice := make([]T, 0)
	node := list.Front()
	for node != list.dummy {
		slice = append(slice, node.val)
		node = node.next
	}
	return slice
}

func TestListInt(t *testing.T) {
	t.Run("PushPopMoveBack", func(t *testing.T) {
		tests := []struct {
			name     string
			input    []int
			expected []int
			front    int
		}{
			{
				name:     "basic",
				input:    []int{1, 2, 3, 4, 5},
				expected: []int{1, 2, 3, 4, 5},
				front:    1,
			},
			{
				name:     "one element",
				input:    []int{1},
				expected: []int{1},
				front:    1,
			},
			{
				name:     "empty",
				input:    []int{},
				expected: []int{},
				front:    0,
			},
		}
		for _, tt := range tests {
			list := NewList[int]()
			for _, num := range tt.input {
				list.PushBack(num)
			}
			if list.Size() != len(tt.expected) {
				t.Errorf("%s: list.Size() = %d. Expected: %d.\n", tt.name, list.Size(), len(tt.expected))
			}
			output := listToSlice(list)
			for i := range output {
				if output[i] != tt.expected[i] {
					t.Errorf("%s: output[%d] = %d. Expected: %d.\n", tt.name, i, output[i], tt.expected[i])
				}
			}
			front := list.Front()
			if front.val != tt.front {
				t.Errorf("%s: list.Front().val = %d. Expected: %d.\n", tt.name, front.val, tt.front)
			}
			list.MoveBack(front)
			back := list.Back()
			if back.val != tt.front {
				t.Errorf("%s: After list.MoveBack(list.Front()): list.Back().val = %d. Expected: %d.\n", tt.name, back.val, tt.front)
			}
			list.MoveFront(list.Back())
			for i := range tt.expected {
				if list.Size() != len(tt.expected)-i {
					t.Errorf("%s: After %d PopBacks: list.Size() = %d. Expected: %d.\n", tt.name, i, list.Size(), len(tt.expected)-i)
				}
				backVal := list.Back().val
				expectedVal := tt.expected[len(tt.expected)-i-1]
				if backVal != expectedVal {
					t.Errorf("%s: After %d PopBacks: list.Back().val = %d. Expected: %d.\n", tt.name, i, backVal, expectedVal)
				}
				list.PopBack()
			}
		}
	})
	t.Run("PushPopMoveFront", func(t *testing.T) {
		tests := []struct {
			name     string
			input    []int
			expected []int
			back     int
		}{
			{
				name:     "basic",
				input:    []int{1, 2, 3, 4, 5},
				expected: []int{5, 4, 3, 2, 1},
				back:     1,
			},
			{
				name:     "one element",
				input:    []int{1},
				expected: []int{1},
				back:     1,
			},
			{
				name:     "empty",
				input:    []int{},
				expected: []int{},
				back:     0,
			},
		}
		for _, tt := range tests {
			list := NewList[int]()
			for _, num := range tt.input {
				list.PushFront(num)
			}
			if list.Size() != len(tt.expected) {
				t.Errorf("%s: list.Size() = %d. Expected: %d.\n", tt.name, list.Size(), len(tt.expected))
			}
			output := listToSlice(list)
			for i := range output {
				if output[i] != tt.expected[i] {
					t.Errorf("%s: output[%d] = %d. Expected: %d.\n", tt.name, i, output[i], tt.expected[i])
				}
			}
			back := list.Back()
			if back.val != tt.back {
				t.Errorf("%s: list.Back().val = %d. Expected: %d.\n", tt.name, back.val, tt.back)
			}
			list.MoveFront(back)
			front := list.Front()
			if front.val != tt.back {
				t.Errorf("%s: After list.MoveBack(list.Back()): list.Front().val = %d. Expected: %d.\n", tt.name, front.val, tt.back)
			}
			list.MoveBack(list.Front())
			for i := range tt.expected {
				if list.Size() != len(tt.expected)-i {
					t.Errorf("%s: After %d PopFronts: list.Size() = %d. Expected: %d.\n", tt.name, i, list.Size(), len(tt.expected)-i)
				}
				frontVal := list.Front().val
				expectedVal := tt.expected[i]
				if frontVal != expectedVal {
					t.Errorf("%s: After %d PopFronts: list.Back().val = %d. Expected: %d.\n", tt.name, i, frontVal, expectedVal)
				}
				list.PopFront()
			}
		}
	})
}

func TestListString(t *testing.T) {
	t.Run("PushPopMoveBack", func(t *testing.T) {
		tests := []struct {
			name     string
			input    []string
			expected []string
			front    string
		}{
			{
				name:     "basic",
				input:    []string{"Assembler", "Basic", "C", "Delphi", "Erlang", "Fortran", "Go"},
				expected: []string{"Assembler", "Basic", "C", "Delphi", "Erlang", "Fortran", "Go"},
				front:    "Assembler",
			},
			{
				name:     "one element",
				input:    []string{"Go"},
				expected: []string{"Go"},
				front:    "Go",
			},
			{
				name:     "empty",
				input:    []string{},
				expected: []string{},
				front:    "",
			},
		}
		for _, tt := range tests {
			list := NewList[string]()
			for _, str := range tt.input {
				list.PushBack(str)
			}
			if list.Size() != len(tt.expected) {
				t.Errorf("%s: list.Size() = %d. Expected: %d.\n", tt.name, list.Size(), len(tt.expected))
			}
			output := listToSlice(list)
			for i := range output {
				if output[i] != tt.expected[i] {
					t.Errorf("%s: output[%d] = %s. Expected: %s.\n", tt.name, i, output[i], tt.expected[i])
				}
			}
			front := list.Front()
			if front.val != tt.front {
				t.Errorf("%s: list.Front().val = %s. Expected: %s.\n", tt.name, front.val, tt.front)
			}
			list.MoveBack(front)
			back := list.Back()
			if back.val != tt.front {
				t.Errorf("%s: After list.MoveBack(list.Front()): list.Back().val = %s. Expected: %s.\n", tt.name, back.val, tt.front)
			}
			list.MoveFront(list.Back())
			for i := range tt.expected {
				if list.Size() != len(tt.expected)-i {
					t.Errorf("%s: After %d PopBacks: list.Size() = %d. Expected: %d.\n", tt.name, i, list.Size(), len(tt.expected)-i)
				}
				backVal := list.Back().val
				expectedVal := tt.expected[len(tt.expected)-i-1]
				if backVal != expectedVal {
					t.Errorf("%s: After %d PopBacks: list.Back().val = %s. Expected: %s.\n", tt.name, i, backVal, expectedVal)
				}
				list.PopBack()
			}
		}
	})
	t.Run("PushPopMoveFront", func(t *testing.T) {
		tests := []struct {
			name     string
			input    []string
			expected []string
			back     string
		}{
			{
				name:     "basic",
				input:    []string{"Assembler", "Basic", "C", "Delphi", "Erlang", "Fortran", "Go"},
				expected: []string{"Go", "Fortran", "Erlang", "Delphi", "C", "Basic", "Assembler"},
				back:     "Assembler",
			},
			{
				name:     "one element",
				input:    []string{"Go"},
				expected: []string{"Go"},
				back:     "Go",
			},
			{
				name:     "empty",
				input:    []string{},
				expected: []string{},
				back:     "",
			},
		}
		for _, tt := range tests {
			list := NewList[string]()
			for _, str := range tt.input {
				list.PushFront(str)
			}
			if list.Size() != len(tt.expected) {
				t.Errorf("%s: list.Size() = %d. Expected: %d.\n", tt.name, list.Size(), len(tt.expected))
			}
			output := listToSlice(list)
			for i := range output {
				if output[i] != tt.expected[i] {
					t.Errorf("%s: output[%d] = %s. Expected: %s.\n", tt.name, i, output[i], tt.expected[i])
				}
			}
			back := list.Back()
			if back.val != tt.back {
				t.Errorf("%s: list.Back().val = %s. Expected: %s.\n", tt.name, back.val, tt.back)
			}
			list.MoveFront(back)
			front := list.Front()
			if front.val != tt.back {
				t.Errorf("%s: After list.MoveBack(list.Back()): list.Front().val = %s. Expected: %s.\n", tt.name, front.val, tt.back)
			}
			list.MoveBack(list.Front())
			for i := range tt.expected {
				if list.Size() != len(tt.expected)-i {
					t.Errorf("%s: After %d PopFronts: list.Size() = %d. Expected: %d.\n", tt.name, i, list.Size(), len(tt.expected)-i)
				}
				frontVal := list.Front().val
				expectedVal := tt.expected[i]
				if frontVal != expectedVal {
					t.Errorf("%s: After %d PopFronts: list.Back().val = %s. Expected: %s.\n", tt.name, i, frontVal, expectedVal)
				}
				list.PopFront()
			}
		}
	})
}
