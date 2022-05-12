package executors

import (
	"log"
	"os/exec"
)

type BashExecutor struct {
}

func NewBashExecutor() *BashExecutor {
	return &BashExecutor{}
}

func (ex *BashExecutor) ExecuteCommand(commands ...string) {

	var inLine string = ""

	if len(commands) == 1 {
		inLine = commands[0]
	} else {
		inLine = ex.generateInLine(commands...)
	}

	log.Println(inLine)

	cmd := exec.Command("/bin/sh", "-c", inLine)
	res, _ := cmd.CombinedOutput()
	log.Println(string(res))
}

func (ex *BashExecutor) generateInLine(commands ...string) string {
	inLine := ""

	for _, command := range commands {
		inLine = inLine + command + ";"
	}
	return inLine

}
