package commands

import (
	"log"
	"strings"

	"github.com/JuanGQCadavid/projectsUpdater/project/internal/domain/models"
)

type RunDockerCommand struct {
	projectConfig *models.ProjectConfig
}

func NewRunDockerCommand(projectConfig *models.ProjectConfig) *RunDockerCommand {
	return &RunDockerCommand{
		projectConfig: projectConfig,
	}
}

func (cmd *RunDockerCommand) Execute() {
	command := cmd.generateBashCommand()
	log.Println(command)

}

func (cmd *RunDockerCommand) generateBashCommand() string {
	command := `
		CONTAINER_ID=$(docker run \
			--name ${APP_NAME} \
			--rm \
			-p 8080:8080 \
			-d \
			-t ${APP_NAME}-service /bin/bash /app/start.sh)

		docker logs -f $CONTAINER_ID
	`
	command = strings.ReplaceAll(command, "${APP_NAME}", cmd.projectConfig.Project.Name)

	return command

}
