package models

type ConfigFile struct {
	Path string `json:"path"`
}

func NewConfigFile(path string) *ConfigFile {
	return &ConfigFile{
		Path: path,
	}
}
