package config

import (
	"os/user"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

var Config ClusterConfig

// ClusterConfig ...
type ClusterConfig struct {
	Username string `toml:"dbusername"`
	Password string `toml:"dbpassword"`

	Vip   Vip
	Atlas map[string]Atlas
	DB    map[string]Database `toml:"database"`
}

// Database ...
type Database struct {
	IP   string
	Port int
}

// Atlas ...
type Atlas struct {
	IP   string
	Port int
}

// Vip ...
type Vip struct {
	IP   string
	Port int
}

// ParseConfig return ClusterConfig instance
func ParseConfig(data string) error {
	if _, err := toml.Decode(data, &Config); err != nil {
		return err
	}
	return nil
}

// GetAppPath return the working path for the app
// as root user, the result is `/root/.zj-db-cluster`
func GetAppPath() (string, error) {
	osUsr, err := user.Current()
	if err != nil {
		return "", err
	}
	appPath := filepath.Join(osUsr.HomeDir, ".zj-db-cluster")
	return appPath, nil
}

// GetCfgFilePath ...
func GetCfgFilePath(clusterName string) (string, error) {
	cfgFile := clusterName + ".toml"
	appPath, err := GetAppPath()
	if err != nil {
		return "", err
	}
	return filepath.Join(appPath, cfgFile), nil
}
