package config

import (
	"os/user"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

var Config ClusterConfig

// ClusterConfig ...
type ClusterConfig struct {
	Vip   Vip
	Atlas map[string]Atlas
	DB    map[string]Database `toml:"database"`
}

// Database ...
type Database struct {
	IP         string
	Port       int
	Dbusername string
	Dbpassword string
}

// Atlas ...
type Atlas struct {
	IP         string
	Port       int
	Dbusername string
	Dbpassword string
}

// Vip ...
type Vip struct {
	IP         string
	Port       int
	Dbusername string
	Dbpassword string
}

// ParseConfig return ClusterConfig instance
func ParseConfig(data string) error {
	if _, err := toml.Decode(data, &Config); err != nil {
		return err
	}
	return nil
}

func GetAppPath() (string, error) {
	osUsr, err := user.Current()
	if err != nil {
		return "", err
	}
	appPath := filepath.Join(osUsr.HomeDir, ".zj-db-cluster")
	return appPath, nil
}

func GetCfgFilePath(clusterName string) (string, error) {
	cfgFile := clusterName + ".toml"
	appPath, err := GetAppPath()
	if err != nil {
		return "", err
	}
	return filepath.Join(appPath, cfgFile), nil
}
