package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/Aksh-Bansal-dev/simpledb"
)

var n = flag.Int("n", 10000, "Number of entries")

func main() {
	flag.Parse()
	log.SetFlags(log.Lshortfile)
	db := simpledb.NewDatabase("simple.db")
	defer db.Close()

	runWithIndex(*n, db)
	runWithoutIndex(*n, db)
}

func runWithoutIndex(n int, db *simpledb.SimpleDB) {
	rand.Seed(time.Now().Unix())

	// Putting
	startTime := time.Now()
	for i := 0; i < n; i++ {
		db.Put(fmt.Sprintf("key%d", rand.Int31n(int32(n))), "val")
	}

	fmt.Printf("[no index] Inserted %d entries in %v\n", n, time.Now().Sub(startTime))

	// Getting
	startTime = time.Now()
	for i := 0; i < n; i++ {
		db.GetWithoutIndex(fmt.Sprintf("key%d", i))
	}

	fmt.Printf("[no index] Fetched %d entries in %v\n", n, time.Now().Sub(startTime))

	if err := os.Truncate("simple.db", 0); err != nil {
		log.Printf("Failed to truncate: %v", err)
	}
}

func runWithIndex(n int, db *simpledb.SimpleDB) {
	rand.Seed(time.Now().Unix())

	// Putting
	startTime := time.Now()
	for i := 0; i < n; i++ {
		db.Put(fmt.Sprintf("key%d", rand.Int31n(int32(n))), "val")
	}

	fmt.Printf("[index] Inserted %d entries in %v\n", n, time.Now().Sub(startTime))

	// Getting
	startTime = time.Now()
	for i := 0; i < n; i++ {
		db.Get(fmt.Sprintf("key%d", i))
	}

	fmt.Printf("[index] Fetched %d entries in %v\n", n, time.Now().Sub(startTime))

	if err := os.Truncate("simple.db", 0); err != nil {
		log.Printf("Failed to truncate: %v", err)
	}
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func randStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
