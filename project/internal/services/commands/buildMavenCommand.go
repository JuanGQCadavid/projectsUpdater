package commands

import (
	"fmt"
	"log"
	"strings"

	"github.com/JuanGQCadavid/projectsUpdater/project/internal/domain/models"
)

type BuildCommand struct {
	projectConfig *models.ProjectConfig
}

func NewBuildCommand(projectConfig *models.ProjectConfig) *BuildCommand {
	return &BuildCommand{
		projectConfig: projectConfig,
	}
}

func (cmd *BuildCommand) Execute() {
	bashCommand := cmd.generateBashCommand()

	log.Println(bashCommand)
}

// docker run -it --rm  \
//     -v ${PWD}/pokemon:/usr/src/pokemon \ #${PWD}/pokemon -> homedir
//     -v ${HOME}/.m2:/root/.m2 \ # Optional
//     -v ${PWD}/pokemon/target:/usr/src/pokemon/target \  # ${PWD}/pokemon/target -> targedir
//     -w /usr/src/pokemon maven:3.8.5-jdk-8-slim mvn package

func (cmd *BuildCommand) generateBashCommand() string {
	command := `
		docker run -it --rm  \
			-v ${HOME_DIR}:/usr/src/${APP_NAME} \
			-v ${HOME}/.m2:/root/.m2 \
			-v ${TARGET_DIR}:/usr/src/${APP_NAME}/target \
			-w /usr/src/${APP_NAME} maven:3.8.5-jdk-8-slim mvn package 
	`
	command = strings.ReplaceAll(command, "${HOME_DIR}", fmt.Sprintf("%s/%s", cmd.projectConfig.BaseDir, cmd.projectConfig.Project.Directories.Home))
	command = strings.ReplaceAll(command, "${TARGET_DIR}", fmt.Sprintf("%s/%s", cmd.projectConfig.BaseDir, cmd.projectConfig.Project.Directories.Target))
	command = strings.ReplaceAll(command, "${APP_NAME}", cmd.projectConfig.Project.Name)

	return command

}
