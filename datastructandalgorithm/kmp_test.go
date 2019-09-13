package datastructandalgorithm

import (
	"fmt"
	"log"
	"testing"
	"time"
)

func TestNext(t *testing.T) {
	sub := "abcdabd"
	next := next(sub)
	fmt.Println(next)
}

func TestSubString(t *testing.T) {
	str := `aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab`
	sub := "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab"
	//str := "bbc abcdab abcdabcdabde"
	//sub := "abcdabd"

	start := time.Now().UnixNano()
	for i := 0; i < 1000; i++ {
		_, ok := SubString(str, sub)
		if !ok {
			log.Fatalln("SubString match false")
		}
	}
	end := time.Now().UnixNano()

	for i := 0; i < 1000; i++ {
		_, ok := SubStringNormal(str, sub)
		if !ok {
			log.Fatalln("SubStringNormal match false")
		}
	}
	end1 := time.Now().UnixNano()
	fmt.Println(end - start)
	fmt.Println(end1 - end)

}
