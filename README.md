# go_course_homework_v2

## структура проекта:

- `tasks/task_01/removeAt.go` – сигнатура RemoveAt

- `tasks/task_01/removeAt_test.go` – fuzz-тесты для RemoveAt

- `tasks/task_02/unique.go` – сигнатура Unique

- `tasks/task_02/unique_test.go` – fuzz-тесты для Unique

- `tasks/task_03/findUserByID.go` – структура User и сигнатура FindUserByID

- `tasks/task_03/findUserByID_test.go` – fuzz-тесты для FindUserByID

- `tasks/task_04/groupUserByAge.go` – структура User и сигнатура GroupUsersByAge

- `tasks/task_04/groupUserByAge_test.go` – fuzz-тесты для GroupUsersByAge

Каждый файл функций пустой, вы должны реализовать код сами. Тесты проверяют правильность работы функций и ловят ошибки на случайных входных данных с помощью fuzz-тестов.

В `task_02, task_03, task_04` лучше всего проверять уникальность и наличие элемента с помощью `map`, так производительнее в отличние от сравнения через `slice`!!!

**Пример:**
```
int main() {
    users := []User{
            {
                ID:   1,
                Name: "john",
            },
            {
                ID:   2,
                Name: "alex",
            },
    }

    //unique values
    usersMap := make(map[int64]User, len(users))
        for _, user := range users {
            if _, ok := usersMap[user.ID]; !ok {
                usersMap[user.ID] = user
            }
        }

    fmt.Printf("%#v\n", usersMap)

    fmt.Println(findInSlice(2, users)) //&main.User{ID:2, Name:"alex"}
    fmt.Println(findInMap(1, usersMap)) //&main.User{ID:1, Name:"john"}
}

//O(n)
func findInSlice(id int64, users []User) *User {
	for _, user := range users {
		if user.ID == id {
			return &user
		}
	}
	return nil
}

//O(1)
func findInMap(id int64, usersmap map[int64]User) *User {
	if user, ok := usersmap[id]; ok {
		return &user
	}
	return nil
}
```

Чтобы выполнить задания, нужно заклонить эту репу:

```
git clone https://github.com/shar1mo/gocourse_hols_task_v2
```

## Tasks
### task_01 — RemoveAt

- Функция: ```RemoveAt(nums []int, index int) ([]int, error)```

- Удаляет элемент из слайса nums по индексу `index`.

- Если индекс некорректный (отрицательный или больше/равен длине слайса) → возвращаем ошибку.

- Если индекс корректный → возвращаем новый слайс без этого элемента.

Примеры:

```
RemoveAt([]int{1,2,3,4,5}, 2) → [1,2,4,5], nil
```

```
RemoveAt([]int{1,2,3}, 10) → nil, ошибка
```

### task_02 - Unique

- Функция: ```Unique(nums []int) []int```

- Убирает дубликаты из слайса `nums`.

- Сохраняет порядок элементов.

- Возвращает новый слайс.

Примеры:

```
Unique([]int{1,2,2,3,1,4}) → [1,2,3,4]
```
```
Unique([]int{5,5,5}) → [5]
```

### task_03 - FindUserByID

- Функция:
```
type User struct {
	ID   int
	Name string
	Age  int
}

FindUserByID(users []User, id int) (*User, error)
```

- Находит пользователя по `ID`.

- Если найден → возвращаем указатель на пользователя и `nil`.

- Если не найден → ошибка `"user not found"`.

- Если в слайсе несколько пользователей с одинаковым ID → ошибка `"duplicate user id"`.

Примеры:
```
users := [
 {ID:1, Name:"Bob", Age:20},
 {ID:2, Name:"Alice", Age:30},
 {ID:3, Name:"Tom", Age:25},
]

FindUserByID(users, 2) → &User{ID:2, Name:"Alice", Age:30}, nil
FindUserByID(users, 10) → nil, "user not found"
```

```
users := [
 {ID:1, Name:"Bob", Age:20},
 {ID:2, Name:"Alice", Age:30},
 {ID:2, Name:"Eve", Age:28},
]

FindUserByID(users, 2) → nil, "duplicate user id"
```

### task_04 — GroupUsersByAge

- Функция:
```
type User struct {
	Name string
	Age  int
}

GroupUsersByAge(users []User) (map[int][]User, error)
```

- Группирует пользователей по возрасту.

- Возвращает `map: ключ — возраст`, значение — слайс пользователей с этим возрастом.

Ошибки:

пустое имя → `"empty name"`

возраст < 0 → `"invalid age"`

повторяющееся имя → `"duplicate name"`

Примеры:
```
users := [
 {Name:"Bob", Age:20},
 {Name:"Alice", Age:25},
 {Name:"Tom", Age:20},
 {Name:"Kate", Age:25},
]

GroupUsersByAge(users) → {
  20: [{Name:"Bob", Age:20}, {Name:"Tom", Age:20}],
  25: [{Name:"Alice", Age:25}, {Name:"Kate", Age:25}],
}, nil
```

```
users := [
 {Name:"", Age:30},
]

GroupUsersByAge(users) → nil, "empty name"
```

```
users := [
 {Name:"John", Age:-5},
]

GroupUsersByAge(users) → nil, "invalid age"
```

```
users := [
 {Name:"Alice", Age:25},
 {Name:"Alice", Age:28},
]

GroupUsersByAge(users) → nil, "duplicate name"
```

## fuzz-тесты

`fuzz-тесты` проверяют свойства функций на случайных данных, а не конкретные кейсы.

Примеры свойств:

`RemoveAt:` после удаления длина меньше на 1, элементы до и после удалённого индекса не изменяются.

`Unique:` результат не содержит дубликатов и элементы изначального слайса.

`FindUserByID:` возвращаемый пользователь имеет правильный ID.

`GroupUsersByAge:` все пользователи присутствуют в map по возрасту.

`Fuzz-тесты` генерируют случайные входные данные и проверяют, что функция не падает и соблюдает свойства.

## как запускать тесты

`!!!ЛУЧШЕ ВСЕГО ЗАПУСКАТЬ ИМЕННО fuzz-тесты, они более комплексные и полезны для дебага и логирования!!!`

1. Клонируйте репозиторий и перейдите в папку проекта.

2. Убедитесь, что установлена go версии 1.21 или выше: `go version` в терминале

3. Запуск всех unit-тестов командой: `go test ./...`, но нам они менее интересны

4. Запуск конкретного fuzz-теста (например, RemoveAt): `cd tasks/task_0x/` (x - номер задания), затем `go test -fuzz=FuzzRemoveAt -fuzztime=10s` 

**(-fuzz=FuzzRemoveAt, -fuzz=FuzzUnique, -fuzz=FuzzFindUserByID, -fuzz=FuzzGroupUsersByAge)**, для каждого теста использовать свой **-fuzz="название теста"**

5. Можно сохранять кэш найденных интересных кейсов: `go test -fuzz=FuzzGroupUsersByAge -test.fuzzcachedir=./fuzz_cache`.

**Что хранится в кэше**

- go сохраняет входные данные, которые сломали или проверили функцию как «интересные» кейсы.

- Эти кейсы потом можно повторно использовать, чтобы функция всегда проверялась на найденные ошибки.

- Папка кэша обычно содержит бинарные файлы с закодированными входными данными.

**Совет**

- fuzz-тесты не проверяют конкретные значения, а свойства функций.

- Можно использовать их для проверки на пустые слайсы, отрицательные индексы, повторяющиеся элементы и граничные значения.

- При работе со структурами `(User)` удобно конвертировать случайные `[]byte` в нужные поля,потому что fuzz-ер не умеет принимать структуры как `seed`

## полезные источники

- тутор по map-ам - [map tutorial golang](https://metanit.com/go/tutorial/2.14.php)

- статья про fuzz-тесты на habr - [fuzzing-тесты](https://habr.com/ru/companies/selectel/articles/709248/)