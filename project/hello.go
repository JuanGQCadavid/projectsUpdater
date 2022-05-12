package main

// import (
// 	"fmt"
// 	"os"
// 	"path/filepath"
// )

// func main() {
// 	ex, err := os.Executable()
// 	if err != nil {
// 		panic(err)
// 	}
// 	exPath := filepath.Dir(ex)
// 	fmt.Println(exPath)
// }

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	log.Println("HI!")
	app := "pwd"

	cmd := exec.Command(app) // arg0, arg1, arg2, arg3)
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Print the output
	fmt.Println(string(stdout))

	hi("1", "2", "3")
}

func hi(params ...string) {
	log.Println(params)
}
