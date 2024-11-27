package logic

import (
	"node/db"
	"node/utils"
)

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
	db     *db.DB
	Status Status
	Config Config
}

func NewLogic(db *db.DB) *Logic {
	return &Logic{
		db:     db,
		Status: NewStatusLogic(),
		Config: NewConfigLogic(),
	}
}
