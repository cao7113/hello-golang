package nesting

import "fmt"

type User struct {
	Name  string `json:"name"`
	Email string `json:"email,omitempty"`
}

type BlogUser struct {
	User
	Title string `json:"title"`
}

func Hi(user BlogUser) {
	fmt.Printf("%+v", user)
}
