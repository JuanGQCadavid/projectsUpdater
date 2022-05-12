package commands

import "github.com/JuanGQCadavid/projectsUpdater/project/internal/domain/models"

type SetCommand struct {
	projectConfig models.ProjectConfig
}

func (cmd *SetCommand) Execute() {

}
