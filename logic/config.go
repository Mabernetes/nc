package logic

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
)

type ConfigLogic struct {
	rootDir string
}

func NewConfigLogic() *ConfigLogic {
	dir := os.Getenv("M8S_CONFIG_DIR")
	if dir == "" {
		dir = "~/m8s"
	}
	return &ConfigLogic{
		rootDir: dir,
	}
}

type ConfigFile struct {
	Deployment string `json:"deployment"`
	Pod        string `json:"pod"`
}

type ConfigsTree []ConfigFile

func (l *ConfigLogic) GetTree() (ConfigsTree, error) {
	var out ConfigsTree

	dir, err := os.ReadDir(l.rootDir)
	if err != nil {
		return out, err
	}

	for _, deployment := range dir {
		if deployment.IsDir() {
			deploymentDir, _ := os.ReadDir(filepath.Join(l.rootDir, deployment.Name()))
			for _, file := range deploymentDir {
				re := regexp.MustCompile(`^pod-*`)
				reYaml := regexp.MustCompile(`.yaml`)
				if !file.IsDir() && re.MatchString(file.Name()) {
					out = append(out, ConfigFile{
						Deployment: deployment.Name(),
						Pod:        reYaml.ReplaceAllString(re.ReplaceAllString(file.Name(), ""), ""),
					})
				}
			}
		}
	}

	return out, nil
}

func (l *ConfigLogic) LoadConfigFile(deployment, pod string) (string, error) {
	var path string = filepath.Join(l.rootDir, deployment)
	if pod == "" {
		path = filepath.Join(path, "deployment.yaml")
	} else {
		path = filepath.Join(path, fmt.Sprintf("pod-%s.yaml", pod))
	}

	return path, nil
}
