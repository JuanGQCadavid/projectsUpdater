package commands

import (
	"github.com/JuanGQCadavid/projectsUpdater/project/internal/services/commands/scripts"
	"github.com/JuanGQCadavid/projectsUpdater/project/internal/services/executors"
)

type InitDjUtilsCommand struct {
	Path    string
	Writter executors.FileWritterExecutor
}

func NewInitDjUtilsCommand(path string) *InitDjUtilsCommand {
	return &InitDjUtilsCommand{
		Path:    path,
		Writter: *executors.NewFileWritterExecutor(),
	}
}

func (cmd *InitDjUtilsCommand) Execute() {
	fileName := cmd.Path + "/djutils.yml"
	data := scripts.DjTemplate
	cmd.Writter.ExecuteCommand(fileName, data)
}
