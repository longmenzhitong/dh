package dh

import (
	"errors"
	"os"

	"gopkg.in/yaml.v3"
)

const configFileName = "config.yml"

var Config config

type config struct {
	// TODO 添加数据库配置
}

func (c *config) Init() {
	dir := ProjectDir()
	if _, err := os.Stat(dir); errors.Is(err, os.ErrNotExist) {
		err = os.Mkdir(dir, os.ModePerm)
		if err != nil {
			panic(err)
		}
	}

	configPath := Path(configFileName)
	if _, err := os.Stat(configPath); errors.Is(err, os.ErrNotExist) {
		Config = config{}
		Config.check()
		out, err := yaml.Marshal(Config)
		if err != nil {
			panic(err)
		}
		RewriteLinesToPath(configPath, []string{string(out)})
		return
	}

	f, err := os.ReadFile(configPath)
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(f, &c)
	if err != nil {
		panic(err)
	}
	c.check()
}

func (c *config) check() {
	// TODO 检查配置默认值
}
