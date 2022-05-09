package yml

import (
	"io/ioutil"
	"log"

	"github.com/JuanGQCadavid/projectsUpdater/project/internal/domain/models"
)

type YmlReader struct {
}

func NewYmlReader() *YmlReader {
	return &YmlReader{}
}

func (reader *YmlReader) GetProjectConfig(path string) (error, models.ProjectConfig) {
	ymlFile, err := ioutil.ReadFile(path)

	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}

}
