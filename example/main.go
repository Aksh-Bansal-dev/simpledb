package main

import (
	"log"

	"github.com/Aksh-Bansal-dev/simpledb"
)

func main() {
	log.SetFlags(log.Lshortfile)
	db := simpledb.NewDatabase("simple.db")
	defer db.Close()
	db.Put("some", "creatinve")
	db.Get("some")
	// db.Get("rank")
	// hui := simpledb.Entry{"hi", "world"}
	// huis := hui.MarshalEntry()
	// log.Println(huis)
	// log.Println(simpledb.UnmarshalEntry(huis))
}
