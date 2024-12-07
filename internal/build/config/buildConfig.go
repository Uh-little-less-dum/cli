package build_config

import (
	"os"

	utils_error "github.com/Uh-little-less-dum/cli/internal/utils/errorHandling"
	schemas_app_config "github.com/Uh-little-less-dum/go-utils/pkg/schemastructs/ulldAppConfig"
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
