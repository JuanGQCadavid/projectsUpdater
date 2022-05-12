package executors

import (
	"log"
	"os"
)

type FileWritterExecutor struct {
}

func NewFileWritterExecutor() *FileWritterExecutor {
	return &FileWritterExecutor{}
}

func (ex *FileWritterExecutor) ExecuteCommand(filePath string, data string) {
	err := os.WriteFile(filePath, []byte(data), 0644)

	log.Println(filePath)
	log.Println(data)

	if err != nil {
		log.Fatalln(err.Error())
	}
}
