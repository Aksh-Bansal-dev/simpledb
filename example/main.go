package main

import (
	"fmt"
	"log"

	"github.com/Aksh-Bansal-dev/simpledb"
)

func main() {
	log.SetFlags(log.Lshortfile)
	db := simpledb.NewDatabase("simple.db")
	defer db.Close()
	val, present := db.Get("hello")
	if present {
		fmt.Println(val)
	} else {
		fmt.Println("not found")
	}
	db.Put("hello", "world")
	val, present = db.Get("hello")
	if present {
		fmt.Println(val)
	} else {
		fmt.Println("not found")
	}
}
