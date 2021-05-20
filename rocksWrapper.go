package main

/*
#cgo CFLAGS: -I./lib/rocksdb/include
#cgo LDFLAGS: -L./lib/rocksdb -lrocksdb
#include <stdlib.h>
#include "rocksdb/c.h"

*/
import "C"

import "fmt"

func dbOpen(dbName string) {

	//rocksdb_open(
	//  const rocksdb_options_t* options, const char* name, char** errptr);
	name := C.CString(dbName)
	err := ""
	errptr := C.CString(err)
	options := C.struct_rocksdb_options_t

	db := C.rocksdb_open(&options, name, &errptr)
}

func main() {
	fmt.Println("vim-go")
}
