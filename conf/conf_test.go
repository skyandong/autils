package conf

import (
	"path"
	"runtime"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// Conf for the service
type Conf struct {
	// App is the service
	App *struct {
		// Name is the service name
		Name string `mapstructure:"name"`
		// Ports is the service ports
		Ports map[string]int `mapstructure:"ports"`
		// Register self to consul
		Register bool `mapstructure:"register"`
		// RegisterDelay after serve
		RegisterDelay time.Duration `mapstructure:"register_delay"`
		// ControlPort is local control port
		ControlPort int `yaml:"controlPort"`
	} `mapstructure:"app"`
}

func TestLoadConfig(t *testing.T) {
	c := Conf{}
	_, file, _, _ := runtime.Caller(0)
	configFileName := "dev.conf.yaml"
	configPath := path.Join(path.Dir(file), configFileName)
	err := LoadConfig(configPath, &c, LevelInfo)
	assert.NoError(t, err)

	assert.Equal(t, "test", c.App.Name)
	assert.Equal(t, map[string]int{"http": 7001, "grpc": 7002}, c.App.Ports)
	assert.True(t, c.App.Register)
	assert.Equal(t, time.Second, c.App.RegisterDelay)
	assert.Equal(t, 7009, c.App.ControlPort)
}
