package main

import (
	"fmt"
	"reflect"
)

type User struct {
	name string `json:"user_name"`
	email string `json:"user_email"`

}

func main()  {
	t := reflect.TypeOf(User{})
	f := t.Field(1).Tag.Get("json")
	fmt.Printf(f)
}
