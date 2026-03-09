package task04_test

import (
	"gocourse_htask_v2/tasks/task_04"
	"testing"
)

func FuzzGroupUsersByAge(f *testing.F) {
	f.Add([]byte{20, 'B', 'o', 'b', 25, 'A', 'l', 'i', 'c', 'e', 20, 'T', 'o', 'm'})

	f.Fuzz(func(t *testing.T, data []byte) {
		const userSize = 4
		if len(data) < userSize {
			return
		}

		var users []task04.User
		for i := 0; i+userSize <= len(data); i += userSize {
			u := task04.User{
				Age:  int(data[i]),
				Name: string(data[i+1 : i+4]),
			}
			users = append(users, u)
		}

		if len(users) == 0 {
			return
		}

		grouped, err := task04.GroupUsersByAge(users)

		for _, u := range users {
			if u.Name == "" {
				if err == nil || err.Error() != "empty name" {
					t.Fatalf("expected empty name error for user %+v, got %v", u, err)
				}
				return
			}
			if u.Age < 0 {
				if err == nil || err.Error() != "invalid age" {
					t.Fatalf("expected invalid age error for user %+v, got %v", u, err)
				}
				return
			}
		}

		seen := make(map[string]struct{})
		for _, u := range users {
			if _, ok := seen[u.Name]; ok {
				if err == nil || err.Error() != "duplicate name" {
					t.Fatalf("expected duplicate name error for user %+v, got %v", u, err)
				}
				return
			}
			seen[u.Name] = struct{}{}
		}

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		for _, u := range users {
			group, ok := grouped[u.Age]
			if !ok {
				t.Fatalf("age %d not found in grouped map", u.Age)
			}

			found := false
			for _, gu := range group {
				if gu == u {
					found = true
					break
				}
			}
			if !found {
				t.Fatalf("user %+v not found in age group %d", u, u.Age)
			}
		}

		for age, group := range grouped {
			for _, gu := range group {
				found := false
				for _, u := range users {
					if u == gu {
						found = true
						break
					}
				}
				if !found {
					t.Fatalf("unexpected user %+v in age group %d", gu, age)
				}
			}
		}
	})
}