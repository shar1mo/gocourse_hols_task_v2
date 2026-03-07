package task02_test

import (
	"gocourse_htask_v2/tasks/task_02"
	"slices"
	"testing"
)

func FuzzUnique(f *testing.F) {
	f.Add([]byte{1, 2, 2, 3, 6, 6, 8, 10, 12, 12, 15, 15, 231, 31, 55})

	f.Fuzz(func(t *testing.T, data []byte) {
		nums := make([]int, len(data))
		for i, v := range data {
			nums[i] = int(v)
		}

		result_slice := task02.Unique(nums)

		if result_slice == nil {
			t.Fatal("empty slice :(")
		}

		tmp_map := make(map[int]struct{})
		for _, v := range result_slice {
			if _, ok := tmp_map[v]; ok {
				t.Fatal("duplicate found")
			}
			tmp_map[v] = struct{}{}
		}

		if !slices.Equal(result_slice, task02.Unique(result_slice)) {
			t.Fatal("not idempotent")
		}

		if len(result_slice) > len(nums) {
			t.Fatal("len(res) > len(source)")
		}

		source := make(map[int]struct{})
		for _, v := range nums {
			source[v] = struct{}{}
		}

		for _, v := range result_slice {
			if _, ok := source[v]; !ok {
				t.Fatal("element not from source")
			}
		}
	})
}