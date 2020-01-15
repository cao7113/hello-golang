package nesting

import "fmt"

type User struct {
	Name  string `json:"name"`
	Email string `json:"email,omitempty"`
}

type UserBlog struct {
	User
	Title string `json:"title"`
}

func Hi(user UserBlog) {
	fmt.Printf("%+v", user)
}
