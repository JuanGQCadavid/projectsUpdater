package commands

import (
	"log"
	"strings"

	"github.com/JuanGQCadavid/projectsUpdater/project/internal/domain/models"
	"github.com/JuanGQCadavid/projectsUpdater/project/internal/domain/ports"
	"github.com/JuanGQCadavid/projectsUpdater/project/internal/services/executors"
)

type BuildImageCommand struct {
	projectConfig *models.ProjectConfig
	bashExecutor  ports.Executors
}

func NewBuildImageCommand(projectConfig *models.ProjectConfig) *BuildImageCommand {
	return &BuildImageCommand{
		projectConfig: projectConfig,
		bashExecutor:  executors.NewBashExecutor(),
	}
}

func (cmd *BuildImageCommand) Execute() {
	bashCommands := cmd.generateBashCommand()
	log.Println(bashCommands)

	cmd.bashExecutor.ExecuteCommand(bashCommands...)
}

// cd /Users/juan.qc/git/PokemonProject/backend/; docker build -t pokemon-service -f Dockerfile /Users/juan.qc/git/PokemonProject/backend/
func (cmd *BuildImageCommand) generateBashCommand() []string {
	cdCommand := "cd ${DOCKER_DIR}"
	buildCommand := "docker build -t ${APP_NAME}-service -f ${DOCKER_FILE} ."
	cdBackCommand := "cd -"

	dockerPathSplited := strings.Split(cmd.projectConfig.Project.Directories.Dockerfile, "/")

	dockerFileName := ""
	dockerFilePath := ""
	if (len(dockerPathSplited)) == 1 {
		dockerFilePath = cmd.projectConfig.BaseDir
	} else {
		dockerFilePath = strings.Join(dockerPathSplited[:len(dockerPathSplited)-1], "/")
	}

	dockerFileName = dockerPathSplited[len(dockerPathSplited)-1]

	cdCommand = strings.ReplaceAll(cdCommand, "${DOCKER_DIR}", dockerFilePath)

	buildCommand = strings.ReplaceAll(buildCommand, "${DOCKER_FILE}", dockerFileName)
	buildCommand = strings.ReplaceAll(buildCommand, "${APP_NAME}", cmd.projectConfig.Project.Name)

	return []string{cdCommand, buildCommand, cdBackCommand}

}
