package repository

import "fmt"

var ()

func GetAllUserList() []*User {
	var userList []*User
	if err := db.Limit(10).Find(&userList); err != nil {
		fmt.Println("GetAllUserList failed: ", err)
	}
	return userList
}
