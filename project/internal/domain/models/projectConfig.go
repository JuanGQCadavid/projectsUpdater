package models

type ProjectConfig struct {
	Version string        `yaml:"version"`
	Config  Configuration `yaml:"config"`
	Project ProjectSpec   `yaml:"project"`
	BaseDir string
}

type Configuration struct {
	OutputFolders OutputFolder `yaml:"output_folders"`
}

type OutputFolder struct {
	Scripts string `yaml:"scripts"`
	Setup   string `yaml:"setup"`
}
type ProjectSpec struct {
	Name        string    `yaml:"name"`
	Directories Directory `yaml:"directories"`
}

type Directory struct {
	Home       string `yaml:"home"`
	Target     string `yaml:"target"`
	Dockerfile string `yaml:"dockerfile"`
}

// version: v1
// config: # All of this is optional
//   output_folders:
//     scripts: path/to/folder # By default is the same folder that live the yaml file.
//     setup: path/to/folder # By default is the same folder that live the yaml file.
//   behaviors:
//     create_setup_script: true # It is true by defaul

// project:
//   name: ProjectName
//   directories:
//     pom: path/to/pom
//     target: path/to/targetx
