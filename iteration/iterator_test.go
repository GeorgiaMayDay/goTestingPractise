package iteration

import "testing"

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 5)
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func TestRepeatWithTable(t *testing.T) {
	tables := []struct {
		num      int
		word     string
		expected string
	}{
		{1, "Ha", "Ha"},
		{2, "Ha", "HaHa"},
		{5, "Maybe", "MaybeMaybeMaybeMaybeMaybe"},
		{5, "Cheedle", "CheedleCheedleCheedleCheedleCheedle"},
	}
	for _, table := range tables {
		result := Repeat(table.word, table.num)
		if result != table.expected {
			t.Errorf("Repeat was incorrect, got: %s, want: %s", result, table.expected)
		}
	}
}
