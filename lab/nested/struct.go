package nested

import "github.com/cao7113/hellogolang/lab/nested/apkg"

type Greeter interface {
	Hi() string
}

type ZhGreeter string

func (d ZhGreeter) Hi() string {
	return "你好"
}

type Chinese struct {
	Greeter
	City string
}

type FruitStore struct {
	*apkg.Apple
	Name string
}

type User struct {
	Name  string `json:"name"`
	Email string `json:"email,omitempty"`
}

func (s User) Score() int {
	return 5
}

type Blog struct {
	User
	Title string `json:"title"`
}
