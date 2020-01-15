package config

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/ghodss/yaml"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/pkg/errors"
	"yunion.io/x/log"
)

func ensureFile(file string) string {
	dir := filepath.Dir(file)
	if _, err := os.Stat(dir); err != nil {
		if os.IsNotExist(err) {
			if err := os.MkdirAll(dir, os.ModePerm); err != nil {
				log.Fatalf("mkdir %q: %v", dir, err)
			}
		}
	}
	if _, err := os.Stat(file); err != nil {
		if os.IsNotExist(err) {
			if err := ioutil.WriteFile(file, []byte(""), 0644); err != nil {
				log.Fatalf("create file %q: %v", file, err)
			}
		}
	}
	return file
}

func getHomeDir() string {
	dir, err := homedir.Dir()
	if err != nil {
		log.Fatalf("get home dir: %v", err)
	}
	return dir
}

func getConfigDir() string {
	return filepath.Join(getHomeDir(), ".ncmbox")
}

func ConfigPath() string {
	return ensureFile(filepath.Join(getConfigDir(), "config.yml"))
}

var globalConfig *Config

type Config struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func GetConfig() (*Config, error) {
	if globalConfig != nil {
		return globalConfig, nil
	}
	cfgPath := ConfigPath()
	content, err := ioutil.ReadFile(cfgPath)
	if err != nil {
		return nil, err
	}
	cfg := new(Config)
	if err := yaml.Unmarshal(content, cfg); err != nil {
		return nil, err
	}
	globalConfig = cfg
	return cfg, nil
}

func EnsureGetConfig() *Config {
	cfg, err := GetConfig()
	if err != nil {
		panic(err)
	}
	return cfg
}

func (c *Config) Save() error {
	content, err := yaml.Marshal(c)
	if err != nil {
		return errors.Wrap(err, "marshal to yaml")
	}
	if err := ioutil.WriteFile(ConfigPath(), content, 0644); err != nil {
		return errors.Wrap(err, "save configfile")
	}
	return nil
}
