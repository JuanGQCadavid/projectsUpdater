package core

import (
	"fmt"
	"log"
	"strconv"

	"github.com/JuanGQCadavid/projectsUpdater/project/internal/domain/models"
	"github.com/JuanGQCadavid/projectsUpdater/project/internal/domain/ports"
	"github.com/JuanGQCadavid/projectsUpdater/project/internal/services/commands"
)

type DjUtilsCore struct {
	configRepository ports.ConfigRepository
	configReader     ports.Reader
}

func NewDjUtilsCore(configRepository ports.ConfigRepository, configReader ports.Reader) *DjUtilsCore {
	return &DjUtilsCore{
		configRepository: configRepository,
		configReader:     configReader,
	}
}

func (core *DjUtilsCore) Build(maven bool, image bool) {
	log.Println(fmt.Sprintf("Set -> maven: %s, image -> %s", strconv.FormatBool(maven), strconv.FormatBool(image)))

	_, configFile := core.configRepository.GetActualConfig()
	_, projectConfig := core.configReader.GetProjectConfig(configFile.Path)

	if maven {
		command := commands.NewBuildCommand(projectConfig)
		command.Execute()
	}

	if image {
		command := commands.NewBuildImageCommand(projectConfig)
		command.Execute()
	}
}

func (core *DjUtilsCore) Init(path string) {
	log.Println(fmt.Sprintf("Init: path-> %s", path))

	cmd := commands.NewInitDockerCommand(path)
	cmd.Execute()

	cmd2 := commands.NewInitDjUtilsCommand(path)
	cmd2.Execute()

}

func (core *DjUtilsCore) Run(image bool, k8s bool) {
	log.Println(fmt.Sprintf("Run -> image: %s, k8s -> %s", strconv.FormatBool(image), strconv.FormatBool(k8s)))

	_, configFile := core.configRepository.GetActualConfig()
	_, projectConfig := core.configReader.GetProjectConfig(configFile.Path)

	if image {
		command := commands.NewRunDockerCommand(projectConfig)
		command.Execute()
	}

	// if k8s {
	// 	command := commands.NewBuildImageCommand(projectConfig)
	// 	command.Execute()
	// }
}

func (core *DjUtilsCore) Set(path string) {
	log.Println(fmt.Sprintf("Set -> path: %s", path))
	core.configRepository.SetActualConfig(*models.NewConfigFile(path))
}

func (core *DjUtilsCore) check(e error) {
	if e != nil {
		panic(e)
	}
}
