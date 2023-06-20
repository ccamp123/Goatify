package goatifyInterface

import (
	"io"
	"os"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func openMessageFile(index string) io.Reader {
	var filePath string = "./messages/" + index + ".txt"
	var messageFile *os.File
	messageFile, err := os.Open(filePath)
	check(err)
	var messageReader = io.Reader(messageFile)
}
