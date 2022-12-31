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
	val, present := db.Get("some")
	if present {
		fmt.Println(val)
	}
	db.Put("some", "sus")
	val, present = db.Get("some")
	if present {
		fmt.Println(val)
	}
}
