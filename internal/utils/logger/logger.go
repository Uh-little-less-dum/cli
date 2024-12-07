package logger

import (
	"time"

	cli_styles "github.com/Uh-little-less-dum/go-utils/pkg/styles"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
	"github.com/spf13/viper"
	"gopkg.in/natefinch/lumberjack.v2"
)

func setLoggerStyles() {
	styles := log.DefaultStyles()
	styles.Levels[log.ErrorLevel] = lipgloss.NewStyle().
		SetString("ERROR").
		Padding(0, 1, 0, 1).
		Background(lipgloss.Color("204")).
		Foreground(lipgloss.Color("0"))

	styles.Levels[log.DebugLevel] = lipgloss.NewStyle().
		SetString("DEBUG").
		Padding(0, 1, 0, 1).
		Foreground(lipgloss.Color("204"))
	styles.Levels[log.InfoLevel] = lipgloss.NewStyle().
		SetString("INFO").
		Padding(0, 1, 0, 1).
		Foreground(lipgloss.Color(cli_styles.UlldBlue))

	styles.Levels[log.WarnLevel] = lipgloss.NewStyle().
		SetString("WARN").
		Padding(0, 2, 0, 1).
		Foreground(lipgloss.Color("#ffff00"))

	styles.Levels[log.FatalLevel] = lipgloss.NewStyle().
		SetString("Oh Shit...").
		Padding(0, 1, 0, 1).
		Background(lipgloss.Color("204")).
		Foreground(lipgloss.Color("0"))
	styles.Keys["err"] = lipgloss.NewStyle().Foreground(lipgloss.Color("204"))
	styles.Values["err"] = lipgloss.NewStyle().Bold(true)
	log.SetStyles(styles)
}

// TODO: Need to test this dynamically. This is almost sure to not work as is. Find a way to bind the *singleton* instance of the logger to both stdout as text and to a json file as valid json.
func setStructuredLogger() {
	// handler := log.New(os.Stderr)
	// logger := slog.New(handler)

	// WARN: Turn this off or set this conditionally when ready to release.
	tea.LogToFile("/Users/bigsexy/Desktop/Go/projects/ulld/cli/dev.log", "")
	logFile := viper.GetString("logFile")

	if logFile != "" {
		log.SetOutput(&lumberjack.Logger{
			Filename:   logFile,
			MaxSize:    500,
			MaxBackups: 3,
			MaxAge:     28,
			Compress:   true,
		})
	}
}

func InitLogger(prefix string) {
	// logger := log.New(os.Stderr)
	log.SetTimeFormat(time.Kitchen)
	log.SetPrefix(prefix)
	ll := viper.Get("logLevel")
	if ll != nil {
		logLevel := ll.(log.Level)
		log.SetLevel(logLevel)
	}

	setStructuredLogger()
	setLoggerStyles()

}
