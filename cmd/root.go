package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	// Version info set by main
	Version   = "dev"
	Commit    = "none"
	BuildDate = "unknown"

	// Flags
	cfgFile string
	verbose bool
)

// rootCmd represents the base command
var rootCmd = &cobra.Command{
	Use:   "termux-cli",
	Short: "A powerful CLI tool for Termux on Android",
	Long: `Termux CLI - Una herramienta de línea de comandos potente y eficiente
diseñada específicamente para Termux en Android.

Proporciona acceso simplificado a:
  - Información del sistema
  - Estado de batería
  - Gestión de almacenamiento
  - Networking
  - APIs de Android via Termux-API
  - Notificaciones y más

Desarrollado con ❤️ para la comunidad de Termux.`,
	Version: Version,
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)

	// Global flags
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.termux-cli.yaml)")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")

	// Version template
	rootCmd.SetVersionTemplate(fmt.Sprintf(`{{with .Name}}{{printf "%%s " .}}{{end}}{{printf "version %%s" .Version}}
Commit: %s
Built: %s
Platform: android/arm64
`, Commit, BuildDate))
}

func initConfig() {
	// Config initialization will be added here when using Viper
	if verbose {
		fmt.Println("Verbose mode enabled")
	}
}

// GetVerbose returns the verbose flag value
func GetVerbose() bool {
	return verbose
}
