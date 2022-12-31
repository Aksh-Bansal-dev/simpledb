package main

import (
	"github.com/Aksh-Bansal-dev/simpledb"
)

func main() {
	db := simpledb.NewDatabase("simple.db")
	defer db.Close()
	db.Put("rank", "1")
}
