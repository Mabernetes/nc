package logic

import "node/src/utils"

type Status interface {
	Server() ServerStatusData
	Runner() (map[string]DeploymentStatusData, error)
}

type Config interface {
	GetTree() (ConfigsTree, error)
	ReadConfigFile(deployment, pod string) (utils.ComposeFile, error)
	SaveConfigFile(deployment, pod string, data utils.ComposeFile) error
	GetFilePath(deployment, pod string) string
}

type Logic struct {
	Status Status
	Config Config
}

func NewLogic() *Logic {
	return &Logic{
		Status: NewStatusLogic(),
		Config: NewConfigLogic(),
	}
}
