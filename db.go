package simpledb

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type SimpleDB struct {
	file     *os.File
	indexMap map[string]int
}

type Entry struct {
	Key string
	Val string
}

func NewDatabase(dbPath string) *SimpleDB {
	f, err := os.OpenFile(dbPath, os.O_CREATE|os.O_RDWR, 0600)
	if err != nil {
		log.Fatal(err)
	}

	// create index
	indexMap := map[string]int{}

	offset, _ := f.Seek(0, io.SeekCurrent)
	curLine, eofReached := nextLine(f)
	for !eofReached {
		entry, err := UnmarshalEntry(curLine)
		if err != nil {
			log.Fatal(err)
		}
		indexMap[entry.Key] = int(offset)
		offset, _ = f.Seek(0, io.SeekCurrent)
		curLine, eofReached = nextLine(f)
	}
	return &SimpleDB{file: f, indexMap: indexMap}
}

func nextLine(f *os.File) (string, bool) {
	var sb strings.Builder
	end := false
	for {
		char := make([]byte, 1)
		ret, _ := f.Read(char)
		if ret == 0 {
			end = true
			break
		}
		if char[0] == '\n' {
			break
		}
		sb.Write(char)
	}
	return sb.String(), end
}

func (db *SimpleDB) Put(key, value string) {
	entry := Entry{Key: key, Val: value}
	offset, _ := db.file.Seek(0, io.SeekEnd)
	db.indexMap[key] = int(offset)
	if _, err := db.file.WriteString(fmt.Sprintf("%s\n", entry.Marshal())); err != nil {
		log.Println(err)
	}
}

func (db *SimpleDB) Get(key string) (string, bool) {
	if _, present := db.indexMap[key]; !present {
		return "", false
	}
	db.file.Seek(int64(db.indexMap[key]), io.SeekStart)
	entryLine, _ := nextLine(db.file)
	entry, err := UnmarshalEntry(entryLine)
	if err != nil {
		log.Fatal(err)
	}
	return entry.Val, true
}
func (db *SimpleDB) GetWithoutIndex(key string) (string, bool) {
	db.file.Seek(0, io.SeekStart)
	scanner := bufio.NewScanner(db.file)
	res := ""
	found := false
	for scanner.Scan() {
		entry, err := UnmarshalEntry(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		if entry.Key == key {
			found = true
			res = entry.Val
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return res, found
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

func (entry *Entry) Marshal() string {
	return fmt.Sprintf("%d:%s%d:%s", len(entry.Key), entry.Key, len(entry.Val), entry.Val)
}
