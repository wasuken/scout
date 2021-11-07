package config

import (
	"os/user"
	"strings"

	"github.com/pelletier/go-toml"
)

const CONFIG_PATH = "~/.config/plwc/config.tml"

func configPath() string {
	usr, _ := user.Current()
	return strings.Replace(CONFIG_PATH, "~", usr.HomeDir, 1)
}

type Config struct {
	URL string
}

func ReadConfig() (error, Config) {
	cPath := configPath()
	tree, err := toml.LoadFile(cPath)
	if err != nil {
		return err, Config{}
	}
	url := tree.Get("Base.URL").(string)
	return nil, Config{URL: url}
}
