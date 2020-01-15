package main

import "collection"

users := []type User struct {
	       username string
	       password string
           }{}

func login(i []js.Value) {
	username := js.Global().Get("document").Call("getElementById", i[0].String()).Get("value").String()
	password := js.Global().Get("document").Call("getElementById", i[1].String()).Get("value").String()
	new_user := User{username, password}

	if collection.Collect(users).Contains(new_user) {
		js.Global().Get("document").Call("getElementById", i[2].String()).Set("value", "You are welcome back.")
	} else {
		js.Global().Get("document").Call("getElementById", i[2].String()).Set("value", "Username or password does not match.")
	}
}

func signin(i []js.Value) {
	username := js.Global().Get("document").Call("getElementById", i[0].String()).Get("value").String()
	password := js.Global().Get("document").Call("getElementById", i[1].String()).Get("value").String()
	new_user := User{username, password}

	users = append(users, new_user)
    js.Global().Get("document").Call("getElementById", i[2].String()).Set("value", "You are registered.")	
}

func registerCallbacks() {
	js.Global().Set("login", js.NewCallback(login))
	js.Global().Set("signin", js.NewCallback(signin))
}

func main() {
	c := make(chan struct{}, 0)

	println("WASM Go Initialized")
	registerCallbacks()
	<-c
}
