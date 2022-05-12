package local

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/JuanGQCadavid/projectsUpdater/project/internal/domain/models"
)

const (
	path       string = "/tmp"
	folderName string = "dj-utils"
	fileName   string = "configstate.json"
)

type ConfigRepository struct {
	fileSytemPath string
	filename      string
}

func NewConfigRepository() *ConfigRepository {
	configRepo := &ConfigRepository{
		fileSytemPath: fmt.Sprintf("%s/%s", path, folderName),
		filename:      fileName,
	}
	configRepo.SetUp()

	return configRepo
}

func (conf *ConfigRepository) GetActualConfig() (error, models.ConfigFile) {
	path := fmt.Sprintf("%s/%s", conf.fileSytemPath, conf.filename)

	fileData, _ := os.ReadFile(path)
	configdata := &models.ConfigFile{}

	json.Unmarshal(fileData, configdata)
	return nil, *configdata
}

func (conf *ConfigRepository) SetActualConfig(jsonData models.ConfigFile) {
	path := fmt.Sprintf("%s/%s", conf.fileSytemPath, conf.filename)
	data, _ := json.Marshal(jsonData)
	os.WriteFile(path, data, os.ModePerm)

}

func (conf *ConfigRepository) SetUp() {

	if err := os.MkdirAll(conf.fileSytemPath, os.ModePerm); err != nil {
		log.Fatal(err)
	}

	path := fmt.Sprintf("%s/%s", conf.fileSytemPath, conf.filename)

	if _, err := os.Stat(path); err == nil {
		log.Println("The config file exist!")
	} else if errors.Is(err, os.ErrNotExist) {
		log.Println("The config file DOES NOT exist!")

		_, err := os.Create(path)

		if err != nil {
			log.Panic(err)
		}

		log.Println("File created!")

		// Create
	} else {
		log.Panic(err.Error())
	}
}
