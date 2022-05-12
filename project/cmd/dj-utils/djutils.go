package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/JuanGQCadavid/projectsUpdater/project/internal/core"
	"github.com/JuanGQCadavid/projectsUpdater/project/internal/domain/ports"
	"github.com/JuanGQCadavid/projectsUpdater/project/internal/services/readers/yml"
	"github.com/JuanGQCadavid/projectsUpdater/project/internal/services/repositories/local"
)

type Command string

const (
	RunCommand   Command = "run"
	BuildCommand Command = "build"
	HelpCommand  Command = "help"
	SetCommand   Command = "set"
	InitCommand  Command = "init"
)

func main() {
	argsWithOutProg := os.Args[1:]

	var configRepo ports.ConfigRepository = local.NewConfigRepository()
	var yamlReader ports.Reader = yml.NewYmlReader()

	var coreService core.DjUtilsCore = *core.NewDjUtilsCore(configRepo, yamlReader)

	switch cmd := getCommand(argsWithOutProg); cmd {
	case SetCommand:
		if len(os.Args) < 3 {
			log.Fatalln("You must provide the file path as dj-utils set <path/to/File>")
		}
		fileArgs := os.Args[2]

		if fileArgs == "." {
			cmd := exec.Command("pwd")
			stdout, err := cmd.Output()

			if err != nil {
				log.Fatal(err.Error())
			}

			fileArgs = fmt.Sprintf("%s/djutils.yaml", strings.Split(string(stdout), "\n")[0])
		}

		coreService.Set(fileArgs)

	case InitCommand:
		if len(os.Args) < 3 {
			log.Fatalln("You must provide the file path as dj-utils init <path/to/dir>")
		}
		dirArg := os.Args[2]

		if dirArg == "." {
			cmd := exec.Command("pwd")
			stdout, err := cmd.Output()

			if err != nil {
				log.Fatal(err.Error())
			}

			dirArg = strings.Split(string(stdout), "\n")[0]
			fmt.Println(dirArg)
		}

		coreService.Init(dirArg)

	case BuildCommand:
		mvn, image := false, false
		if len(os.Args) < 3 {
			log.Println("No parameters, so building mvn and image together.")
			mvn, image = true, true

		} else {
			mvn, image = getBuildArgs(os.Args[2:])
		}
		coreService.Build(mvn, image)

	case RunCommand:

		image, k8s := false, false
		if len(os.Args) < 3 {
			log.Println("No parameters, so running it in docker")
			image, k8s = true, false

		} else {
			image, k8s = getRundArgs(os.Args[2:])
		}
		coreService.Run(image, k8s)
	}

}

func readFile() {
	var yamlReader ports.Reader = yml.NewYmlReader()
	print(yamlReader)

	err, data := yamlReader.GetProjectConfig("/Users/juan.qc/git/projectsUpdater/example_tenplate.project.yaml")

	if err != nil {
		log.Println(err)
	}

	fmt.Printf("%+v\n", data)
}

func getRundArgs(args []string) (bool, bool) {
	var image bool = false
	var k8s bool = false

	for _, arg := range args {
		switch arg {
		case "image":
			image = true
		case "k8s":
			k8s = true
		}
	}

	return image, k8s
}

func getBuildArgs(args []string) (bool, bool) {
	var image bool = false
	var mvn bool = false

	for _, arg := range args {
		switch arg {
		case "image":
			image = true
		case "mvn":
			mvn = true
		}
	}

	return mvn, image
}
func getCommand(args []string) Command {
	if len(args) == 0 {
		log.Fatal("You should specified the argument, run dj-utils help for more info.")
	}
	var cmd Command

	switch command := args[0]; command {
	case "run":
		log.Println("run Command")
		cmd = RunCommand
	case "build":
		log.Println("Build Command")
		cmd = BuildCommand
	case "help":
		log.Println("help Command")
		cmd = HelpCommand
	case "set":
		log.Println("set Command")
		cmd = SetCommand
	case "init":
		log.Println("Init Command")
		cmd = InitCommand
	default:
		log.Fatal("Command no found, run dj-utils help for more info.")
	}

	return cmd
}
