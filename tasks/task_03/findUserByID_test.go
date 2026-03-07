package task03_test

import (
	"gocourse_htask_v2/tasks/task_03"
	"testing"
)

func FuzzFindUserByID(f *testing.F) {
	f.Add([]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30}, 5)

	f.Fuzz(func(t *testing.T, data []byte, id int) {
		const userSize = 5
		var users []task03.User
		for i := 0; i+userSize <= len(data); i += userSize {
			users = append(users, task03.User{
				ID:   int(data[i]),
				Name: string(data[i+1 : i+4]),
				Age:  int(data[i+4]),
			})
		}

		if len(users) == 0 {
			return
		}

		result_user, result_err := task03.FindUserByID(users, id)

		if result_err == nil {
			if result_user.ID != id {
				t.Fatalf("expected ID: %d, got: %d", id, result_user.ID)
			}
		}

		if result_err != nil {
			if result_err.Error() != "user not found" && result_err.Error() != "duplicate user id" {
				t.Fatalf("unexpeted error: %v", result_err)
			}
		}
	})
}