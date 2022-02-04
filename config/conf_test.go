package config

import (
	"os"
	"path"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const (
	//configFolderName = "conf"
	Env = "dev"
)

// Conf for the service
type Conf struct {
	// App is the service
	App *struct {
		// Name is the service name
		Name string `yaml:"name"`
		// Ports is the service ports
		Ports map[string]int `yaml:"ports"`
		// Register self to consul
		Register bool `yaml:"register"`
		// RegisterDelay after serve
		RegisterDelay time.Duration `yaml:"registerDelay"`
		// ControlPort is local control port
		ControlPort int `yaml:"controlPort"`
	} `yaml:"app"`
}

func TestLoadConfig(t *testing.T) {
	c := Conf{}
	configDir := "config"
	for i := 0; i < 3; i++ {
		if info, err := os.Stat(configDir); err == nil && info.IsDir() {
			break
		}
		configDir = filepath.Join("..", configDir)
	}
	configFileName := Env + ".conf.yaml"
	configPath := path.Join(configDir, configFileName)
	err := LoadConfig(configPath, &c, LevelInfo)
	assert.NoError(t, err)

	assert.Equal(t, "test", c.App.Name)
	assert.Equal(t, map[string]int{"http": 7001, "grpc": 7002}, c.App.Ports)
	assert.True(t, c.App.Register)
	assert.Equal(t, time.Second, c.App.RegisterDelay)
	assert.Equal(t, 7009, c.App.ControlPort)
}
