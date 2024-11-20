// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    cLIConfig, err := UnmarshalCLIConfig(bytes)
//    bytes, err = cLIConfig.Marshal()

package schemas_cli_config

import "encoding/json"

func UnmarshalCLIConfig(data []byte) (CLIConfig, error) {
	var r CLIConfig
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *CLIConfig) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type CLIConfig struct {
	Build                                                        *Build    `json:"build,omitempty"`
	// Absolute path to a log file used during the build process.          
	LogFile                                                      *string   `json:"logFile,omitempty"`
	LogLevel                                                     *LogLevel `json:"logLevel,omitempty"`
}

type Build struct {
	// Absolute path to the directory where the ULLD app should be generated.        
	BuildDir                                                                 *string `json:"buildDir,omitempty"`
}

type LogLevel string

const (
	Debug LogLevel = "debug"
	Error LogLevel = "error"
	Info  LogLevel = "info"
	Warn  LogLevel = "warn"
)
