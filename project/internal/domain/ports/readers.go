package ports

import (
	"github.com/JuanGQCadavid/projectsUpdater/project/internal/domain/models"
)

type Reader interface {
	GetProjectConfig(path string) (error, *models.ProjectConfig)
}
