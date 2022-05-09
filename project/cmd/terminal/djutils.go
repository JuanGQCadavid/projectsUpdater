package main

import (
	"github.com/JuanGQCadavid/projectsUpdater/project/internal/domain/ports"
	"github.com/JuanGQCadavid/projectsUpdater/project/internal/services/readers/yml"
)

func main() {
	var yamlReader ports.Reader = yml.NewYmlReader()
	print(yamlReader)
}
