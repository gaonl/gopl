package main

import (
	"fmt"
	"log"
)

func main() {

		user, err := QueryById(2, )
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(user.Id, user.Name, user.Password)


	/*
		user := domain.User{
			Id:               0,
			Name:             "wskj",
			Password:         "654321",
			RegisterDateTime: time.Now().Unix(),
		}

		id, err := Add(user)
		if err == nil {
			fmt.Println(id)
		}
	*/

	/*
		count, err := Delete(4)
		if err == nil {
			fmt.Println(count)
		}
	*/

	/*
	user := domain.User{
		Id:       5,
		Name:     "hongkong",
		Password: "hongkong123",
	}
	count, err := Update(user)
	if err == nil {
		fmt.Println(count)
	}

	 */
}
