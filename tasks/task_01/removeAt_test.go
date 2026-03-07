package task01_test

import (
	"gocourse_htask_v2/tasks/task_01"
	"testing"
)

func FuzzRemoveAt(f *testing.F) {
	f.Add([]byte{1, 2, 3, 4, 5}, 2)

	f.Fuzz(func(t *testing.T, data []byte, index int){
		nums := make([]int, len(data))
		for i, v := range data {
			nums[i] = int(v)
		}

		result_slice, result_err := task01.RemoveAt(nums, index)

		if index < 0 || index >= len(nums) {
			if result_err == nil {
				t.Fatal("expected err :(")
			}
			return
		}

		if result_err != nil {
			t.Fatalf("unexpected err: %v", result_err)
		}

		if len(result_slice) != len(nums) - 1 {
			t.Fatalf("wrong length: got: %d want: %d", len(result_slice), len(nums) - 1)
		}

		for i := 0; i < index; i++ {
			if result_slice[i] != nums[i] {
				t.Fatal("elements before index incorrect")
			}
		}

		for i := index; i < len(result_slice); i++ {
			if result_slice[i] != nums[i + 1] {
				t.Fatal("elements after index incorrect")
			}
		}
	})
}