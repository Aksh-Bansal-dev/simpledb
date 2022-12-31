package simpledb

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
)

type SimpleDB struct {
	file *os.File
}

type Entry struct {
	Key string
	Val string
}

func NewDatabase(dbPath string) *SimpleDB {
	f, err := os.OpenFile(dbPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	return &SimpleDB{file: f}
}

func (db *SimpleDB) Put(key, value string) {
	if _, err := db.file.WriteString("differetn to append\n"); err != nil {
		log.Println(err)
	}
}

func (db *SimpleDB) Close() {
	db.file.Close()
}

func UnmarshalEntry(s string) (Entry, error) {
	var res Entry
	invalidEntryFormatErr := errors.New("Invalid entry format")

	// key
	pos := 0
	for s[pos] != ':' {
		pos++
	}
	if pos == 0 {
		return res, invalidEntryFormatErr
	}
	keyLen, err := strconv.ParseInt(s[:pos], 10, 0)
	if err != nil {
		return res, invalidEntryFormatErr
	}
	key := s[pos+1 : pos+1+int(keyLen)]

	// val
	pos += 1 + int(keyLen)
	startPos := pos
	for s[pos] != ':' {
		pos++
	}
	valLen, err := strconv.ParseInt(s[startPos:pos], 10, 0)
	if err != nil {
		return res, invalidEntryFormatErr
	}
	val := s[pos+1 : pos+1+int(valLen)]
	res.Key = key
	res.Val = val
	return res, nil
}

func (entry *Entry) MarshalEntry() string {
	return fmt.Sprintf("%d:%s%d:%s", len(entry.Key), entry.Key, len(entry.Val), entry.Val)
}
