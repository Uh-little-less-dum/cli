package build_config

import (
	"os"

	schemas_app_config "github.com/igloo1505/ulldCli/internal/schemastructs/ulldAppConfig"
	utils_error "github.com/igloo1505/ulldCli/internal/utils/errorHandling"
)

type ulldConfig struct {
	path string
	Data schemas_app_config.AppConfig
}

func (c *ulldConfig) readData() {
	content, err := os.ReadFile(c.path)
	utils_error.HandleError(err)
	d, err := schemas_app_config.UnmarshalAppConfig(content)
	utils_error.HandleError(err)
	c.Data = d
}

func GetUlldConfig(filepath string) ulldConfig {
	c := ulldConfig{path: filepath}
	c.readData()
	return c
}
