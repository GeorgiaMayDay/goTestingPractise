package main

import "testing"


func TestSum(t *testing.T) {

	tables := []struct {
		num1     int
		num2     int
		expected int
	}{
		{1, 1, 2},
		{5, 2, 7},
		{5, 52, 57},
	}
	for _, table := range tables {
		total := Sum(table.num1, table.num2)
		if total != table.expected {
			t.Errorf("Sum was incorrect, got: %d, want: %d", total, table.expected)
		}
	}
}
