package leveldb

import (
	"fmt"
	"testing"
)

func TestNewLevelDB(t *testing.T) {
	//dir := t.TempDir()
	dbs, _ := NewLevelDB("./1")
	//dbs.Put([]byte("a"), []byte("b"))
	has, err := dbs.Has([]byte("a"))
	if err != nil {
		return
	}
	value, err := dbs.Get([]byte("c"))
	fmt.Println("err", err)
	fmt.Println("value", value)

	fmt.Println("has", has)
}
