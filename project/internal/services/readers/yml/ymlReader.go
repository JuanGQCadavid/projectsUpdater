package yml

import (
	"io/ioutil"
	"log"
	"strings"

	"github.com/JuanGQCadavid/projectsUpdater/project/internal/domain/models"
	"gopkg.in/yaml.v3"
)

type YmlReader struct {
}

func NewYmlReader() *YmlReader {
	return &YmlReader{}
}

func (reader *YmlReader) GetProjectConfig(path string) (error, *models.ProjectConfig) {
	var projectConfig *models.ProjectConfig = &models.ProjectConfig{}

	ymlFileBuffer, err := ioutil.ReadFile(path)

	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
		return err, &models.ProjectConfig{}
	}

	err = yaml.Unmarshal(ymlFileBuffer, projectConfig)

	if err != nil {
		log.Printf("yaml.Unmarshal err   #%v ", err)
		return err, &models.ProjectConfig{}
	}
	newPath := strings.Split(path, "/")
	otherPath := strings.Join(newPath[:len(newPath)-1], "/")
	projectConfig.BaseDir = otherPath

	return nil, projectConfig

}
