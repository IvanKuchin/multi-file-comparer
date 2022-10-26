package config_reader

type GroupConfig struct {
	Rootdir    string
	Subfolders []string
	Files      []string
}

type Config interface {
	GetGroupConfig(idx int) (*GroupConfig, error)
}
