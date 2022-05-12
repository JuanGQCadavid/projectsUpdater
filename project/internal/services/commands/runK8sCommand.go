package commands

import (
	"strings"

	"github.com/JuanGQCadavid/projectsUpdater/project/internal/domain/models"
)

type RunK8sCommand struct {
	projectConfig models.ProjectConfig
}

func (cmd *RunK8sCommand) Execute() {

}

func (cmd *RunK8sCommand) generateBashCommand() string {
	command := ""
	command = strings.ReplaceAll(command, "${APP_NAME}", cmd.projectConfig.Project.Name)

	return command

}
