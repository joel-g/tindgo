package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {

}

func readKey() {
	dat, err := ioutil.ReadFile(".golang")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(dat))
}
