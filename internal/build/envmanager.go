package build

import (
	"os"
)


type UlldEnv struct {
	ULLD_ADDITIONAL_SOURCES string
}


func (e *UlldEnv) Init() {
	ulldAdditionalSources, _ := os.LookupEnv("ULLD_ADDITIONAL_SOURCES")
	e.ULLD_ADDITIONAL_SOURCES = ulldAdditionalSources
}
