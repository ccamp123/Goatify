package goatifyInterface

import (
	"os",
	"fmt",
	"io"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func openMessageFile(index string) io.TeeReader {
	var filePath string = "./messages/" + index + ".txt"
	var messageFile os.File* 
	messageFile,err := os.Open(filePath)
	check(err)
	 
}
