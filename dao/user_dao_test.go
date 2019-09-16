package dao

import (
	"domain"
	"fmt"
	"log"
	"testing"
	"time"
)

func TestDelete(t *testing.T) {
	_, err := Delete(4)
	if err != nil {
		t.Error(err)
	}
}

func TestQueryById(t *testing.T) {
	user, err := QueryById(2, )
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(user.Id, user.Name, user.Password)
}

func TestUpdate(t *testing.T) {
	user := domain.User{
		Id:       5,
		Name:     "hongkong",
		Password: "hongkong123",
	}
	count, err := Update(user)
	if err == nil {
		fmt.Println(count)
	}
}

func TestAdd(t *testing.T) {

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
}
