package ports

import "github.com/JuanGQCadavid/projectsUpdater/project/internal/domain/models"

type ConfigRepository interface {
	GetActualConfig() (error, models.ConfigFile)
	SetActualConfig(models.ConfigFile)
}
