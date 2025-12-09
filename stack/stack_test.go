package stack

import "testing"

func TestInitializeStackSize(t *testing.T) {
	st := Stack[int]{}
	if st.Size() != 0 {
		t.Error("Incorrect size on the new stack")
	}
}

func TestStackPushAndPeek(t *testing.T) {
	st := Stack[int]{}

	st.Push(9)

	if st.Size() != 1 {
		t.Error("Stach size must be 1")
	}

	got, err := st.Peek()
	if err != nil {
		t.Errorf("Faced an error while trying valid peek: %s", err)
	}

	if got != 9 {
		t.Errorf("Peeked incorrect value, got: %d, want: 9", got)
	}

	if st.Size() != 1 {
		t.Error("Stach size must be 1")
	}
}

func TestStackPushAndPop(t *testing.T) {
	st := Stack[int]{}

	st.Push(9)
	st.Push(8)
	st.Push(7)

	if st.Size() != 3 {
		t.Error("Stach size must be 1")
	}

	got, err := st.Pop()
	if err != nil {
		t.Errorf("Faced an error while trying valid peek: %s", err)
	}

	if got != 7 {
		t.Errorf("Peeked incorrect value, got: %d, want: 9", got)
	}

	if st.Size() != 2 {
		t.Error("Stach size must be 2")
	}

	got2, err := st.Pop()
	if err != nil {
		t.Errorf("Faced an error while trying valid peek: %s", err)
	}

	if got2 != 8 {
		t.Errorf("Peeked incorrect value, got: %d, want: 9", got)
	}

	if st.Size() != 1 {
		t.Error("Stach size must be 2")
	}
}

func TestNumericStackAvg(t *testing.T) {
	st := NumericStack[int]{}

	st.Push(10)
	if st.Avg() != 10.0 {
		t.Errorf("Avg incorrect, got: %f, want: 10.0", st.Avg())
	}

	st.Push(20)
	if st.Avg() != 15.0 {
		t.Errorf("Avg incorrect, got: %f, want: 15.0", st.Avg())
	}

	st.Push(30)
	if st.Avg() != 20.0 {
		t.Errorf("Avg incorrect, got: %f, want: 20.0", st.Avg())
	}

	st.Pop()
	if st.Avg() != 15.0 {
		t.Errorf("Avg incorrect after pop, got: %f, want: 15.0", st.Avg())
	}

	st.Pop()
	if st.Avg() != 10.0 {
		t.Errorf("Avg incorrect after pop, got: %f, want: 10.0", st.Avg())
	}

	// Test Min still works
	st.Push(5)
	min, _ := st.Min()
	if min != 5 {
		t.Errorf("Min incorrect, got: %d, want: 5", min)
	}
	if st.Avg() != 7.5 {
		t.Errorf("Avg incorrect, got: %f, want: 7.5", st.Avg())
	}
}

func TestOrderedStackMin(t *testing.T) {
	st := OrderedStack[int]{}

	st.Push(5)
	min, _ := st.Min()
	if min != 5 {
		t.Errorf("Min incorrect, got: %d, want: 5", min)
	}

	st.Push(3)
	min, _ = st.Min()
	if min != 3 {
		t.Errorf("Min incorrect, got: %d, want: 3", min)
	}

	st.Push(7)
	min, _ = st.Min()
	if min != 3 {
		t.Errorf("Min incorrect after pushing larger value, got: %d, want: 3", min)
	}

	st.Push(2)
	min, _ = st.Min()
	if min != 2 {
		t.Errorf("Min incorrect, got: %d, want: 2", min)
	}

	st.Push(8)
	min, _ = st.Min()
	if min != 2 {
		t.Errorf("Min incorrect, got: %d, want: 2", min)
	}

	// Pop 8, min still 2
	st.Pop()
	min, _ = st.Min()
	if min != 2 {
		t.Errorf("Min incorrect after pop, got: %d, want: 2", min)
	}

	// Pop 2, min should restore to 3
	st.Pop()
	min, _ = st.Min()
	if min != 3 {
		t.Errorf("Min incorrect after popping minimum, got: %d, want: 3", min)
	}

	// Pop 7, min still 3
	st.Pop()
	min, _ = st.Min()
	if min != 3 {
		t.Errorf("Min incorrect, got: %d, want: 3", min)
	}

	// Pop 3, min should be 5
	st.Pop()
	min, _ = st.Min()
	if min != 5 {
		t.Errorf("Min incorrect, got: %d, want: 5", min)
	}
}
