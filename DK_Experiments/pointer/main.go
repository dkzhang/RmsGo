package main

import "fmt"

type User struct {
	ID         int
	Name       string
	Department string
}

func main() {
	users := []User{
		{
			ID:         1,
			Name:       "z1",
			Department: "d1",
		},
		{
			ID:         2,
			Name:       "z2",
			Department: "d2",
		},
		{
			ID:         3,
			Name:       "z3",
			Department: "d3",
		},
	}

	myFunc2(users)
	myFunc3(users)
}

func myFunc1(users []User) {
	fmt.Println("This is myFun1 -->")
	userInfoByID := make(map[int]*User, len(users))
	userInfoByName := make(map[string]*User, len(users))

	for _, v := range users {
		userInfoByID[v.ID] = &v
		userInfoByName[v.Name] = &v
	}

	showResult(userInfoByID, userInfoByName)
}

func myFunc2(users []User) {
	fmt.Println("This is myFun2 -->")
	userInfoByID := make(map[int]*User, len(users))
	userInfoByName := make(map[string]*User, len(users))

	for _, v := range users {
		user := v
		userInfoByID[v.ID] = &user
		userInfoByName[v.Name] = &user
	}

	showResult(userInfoByID, userInfoByName)
	fmt.Println("modify something")
	userInfoByID[2].Name = "zhang2"
	showResult(userInfoByID, userInfoByName)
}

func myFunc3(users []User) {
	fmt.Println("This is myFun3 -->")
	userInfoByID := make(map[int]*User, len(users))
	userInfoByName := make(map[string]*User, len(users))

	for i, v := range users {
		userInfoByID[v.ID] = &users[i]
		userInfoByName[v.Name] = &users[i]
	}

	showResult(userInfoByID, userInfoByName)
	fmt.Println("modify something")
	userInfoByID[3].Name = "new3"
	showResult(userInfoByID, userInfoByName)
}

func showResult(userInfoByID map[int]*User, userInfoByName map[string]*User) {
	for _, uid := range userInfoByID {
		fmt.Printf("userInfoByID = %v \n", *uid)
	}
	for _, uname := range userInfoByName {
		fmt.Printf("userInfoByName = %v \n", *uname)
	}
}
