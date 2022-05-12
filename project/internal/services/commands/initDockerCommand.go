package commands

import (
	"github.com/JuanGQCadavid/projectsUpdater/project/internal/services/commands/scripts"
	"github.com/JuanGQCadavid/projectsUpdater/project/internal/services/executors"
)

type InitDockerCommand struct {
	Path    string
	Writter executors.FileWritterExecutor
}

func NewInitDockerCommand(path string) *InitDockerCommand {
	return &InitDockerCommand{
		Path:    path,
		Writter: *executors.NewFileWritterExecutor(),
	}
}

func (cmd *InitDockerCommand) Execute() {
	fileName := cmd.Path + "/Dockerfile"
	data := scripts.ApisDockerfile
	cmd.Writter.ExecuteCommand(fileName, data)

}
