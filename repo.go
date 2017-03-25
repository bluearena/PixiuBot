package main

var users Users

func RepoAddUser(user *User) {
	users = append(users, user)
}
