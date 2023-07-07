package main

import (
	"html/template"
	"os"
)

type User struct {
	Name        string
	Bio         string
	Age         int
	AccessCodes []int
	Secret      HiddenData
}

type HiddenData struct {
	Password string
}

func main() {
	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}

	user := User{
		Bio:         `<script>alert("Haha, you have been h4x0r3d");</script>`,
		Age:         123,
		AccessCodes: []int{1, 2, 3, 4},
		Secret: HiddenData{
			Password: "bibon",
		},
	}

	err = t.Execute(os.Stdout, user)
	if err != nil {
		panic(err)
	}
}
