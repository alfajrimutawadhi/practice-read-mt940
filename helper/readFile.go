package helper

import (
	"io"
	"log"
	"os"
)

func ReadFile(path string) (string, error) {
	var result string
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	buffer := make([]byte, 1)
	for {
		num, err := file.Read(buffer)
		if err == io.EOF {
			break
		}
		result += string(buffer[:num])
	}
	return result, err
}
