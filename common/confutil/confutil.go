package confutil

import (
	"fmt"
	"github.com/aiagt/aiagt/common/osutil"
	ktcenter "github.com/aiagt/kitextool/conf/center"
	"github.com/bytedance/gopkg/util/logger"
	"gopkg.in/yaml.v2"
	"os"
	"path/filepath"

	ktconf "github.com/aiagt/kitextool/conf"
)

func LoadConf(conf ktconf.Conf, dirs ...string) {
	const (
		confFile        = "conf.yaml"
		confLocalFile   = "conf-local.yaml"
		confReleaseFile = "conf-release.yaml"
	)

	var confFiles []string

	for _, dir := range dirs {
		if IsReleaseEnv() {
			confFiles = append(confFiles, filepath.Join(dir, confReleaseFile))
		} else {
			confFiles = append(confFiles, filepath.Join(dir, confFile))
			confFiles = append(confFiles, filepath.Join(dir, confLocalFile))
		}
	}

	existsOne := false

	for _, file := range confFiles {
		exists := osutil.Exists(file)

		if exists {
			existsOne = true
			ktconf.LoadFiles(conf, file)
		}
	}

	if !existsOne {
		logger.Fatal("config file not found")
	}
}

func IsReleaseEnv() bool {
	e := os.Getenv("GO_ENV")
	return e == "release"
}

func NewConfigCenterWithBackup(dest string) ktcenter.ConfigCenter {
	center := ktcenter.WithConsulConfigCenter(nil)

	center.RegisterCallbacks(func(conf ktconf.Conf) {
		yamlConf, err := yaml.Marshal(conf.GetServerConf())
		if err != nil {
			logger.Warn("yaml marshal config value err:", err)
			return
		}

		err = os.MkdirAll("backup", 0755)
		if err != nil {
			logger.Warn("mkdir backup conf dir err:", err)
			return
		}

		confPath := filepath.Join("backup", fmt.Sprintf("%s.yaml", dest))
		confPrePath := filepath.Join("backup", fmt.Sprintf("%s-pre.yaml", dest))

		if osutil.Exists(confPath) {
			err = osutil.CopyFile(confPath, confPrePath)
			if err != nil {
				logger.Warn("copy config file err:", err)
			}
		}

		err = os.WriteFile(confPath, yamlConf, os.ModePerm)
		if err != nil {
			logger.Warn("backup config file err:", err)
			return
		}
	})

	return center
}
