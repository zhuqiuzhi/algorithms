package sort

import "testing"

func TestSortStudents(t *testing.T) {
	all := []*student{
		{ID: 89, Name: "Alice", Grade: "L1"},
		{ID: 87, Name: "Blob", Grade: "L1"},
	}
	SelectionSort(students(all))
	if all[0].Name != "Blob" {
		t.Fatalf("got %v", all)
	}
}
