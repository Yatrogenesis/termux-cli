package cmd

import (
	"fmt"
	"os"
	"runtime"

	"github.com/spf13/cobra"
)

// systemCmd represents the system command
var systemCmd = &cobra.Command{
	Use:   "system",
	Short: "System information commands",
	Long:  `Get information about the Android/Termux system.`,
}

// systemInfoCmd shows general system information
var systemInfoCmd = &cobra.Command{
	Use:   "info",
	Short: "Display system information",
	Long:  `Display general information about the device, kernel, and Termux environment.`,
	Run: func(cmd *cobra.Command, args []string) {
		printSystemInfo()
	},
}

func init() {
	rootCmd.AddCommand(systemCmd)
	systemCmd.AddCommand(systemInfoCmd)
}

func printSystemInfo() {
	fmt.Println("=== System Information ===")
	fmt.Printf("OS: %s\n", runtime.GOOS)
	fmt.Printf("Architecture: %s\n", runtime.GOARCH)
	fmt.Printf("CPUs: %d\n", runtime.NumCPU())
	fmt.Printf("Go Version: %s\n", runtime.Version())

	if hostname, err := os.Hostname(); err == nil {
		fmt.Printf("Hostname: %s\n", hostname)
	}

	// Environment info
	fmt.Printf("\n=== Termux Environment ===\n")
	if home := os.Getenv("HOME"); home != "" {
		fmt.Printf("HOME: %s\n", home)
	}
	if prefix := os.Getenv("PREFIX"); prefix != "" {
		fmt.Printf("PREFIX: %s\n", prefix)
	}
	if tmpdir := os.Getenv("TMPDIR"); tmpdir != "" {
		fmt.Printf("TMPDIR: %s\n", tmpdir)
	}

	if GetVerbose() {
		fmt.Printf("\n=== Go Runtime ===\n")
		var m runtime.MemStats
		runtime.ReadMemStats(&m)
		fmt.Printf("Alloc: %d MB\n", m.Alloc/1024/1024)
		fmt.Printf("TotalAlloc: %d MB\n", m.TotalAlloc/1024/1024)
		fmt.Printf("Sys: %d MB\n", m.Sys/1024/1024)
		fmt.Printf("NumGC: %d\n", m.NumGC)
	}
}
