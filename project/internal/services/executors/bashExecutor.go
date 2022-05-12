package executors

import (
	"fmt"
	"log"
	"os/exec"
)

type BashExecutor struct {
}

func NewBashExecutor() *BashExecutor {
	return &BashExecutor{}
}

func (ex *BashExecutor) ExecuteCommand(commands ...string) {
	inLine := ""
	for _, command := range commands {
		inLine = inLine + command + ";"

	}

	log.Println(inLine)

	cmd := exec.Command("/bin/sh", "-c", inLine)
	res, _ := cmd.CombinedOutput()
	fmt.Println(string(res))
	// err := cmd.Run()

	// if err != nil {
	// 	log.Panicln(err)
	// }

	// res, _ := cmd.CombinedOutput()
	// fmt.Println(string(res))
}
