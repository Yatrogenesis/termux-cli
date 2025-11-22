package cmd

import (
	"encoding/json"
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

// batteryCmd represents the battery command
var batteryCmd = &cobra.Command{
	Use:   "battery",
	Short: "Battery status and information",
	Long:  `Get battery status, health, and charging information using Termux-API.`,
}

// batteryStatusCmd shows battery status
var batteryStatusCmd = &cobra.Command{
	Use:   "status",
	Short: "Display battery status",
	Long:  `Display current battery status including percentage, health, and charging state.`,
	Run: func(cmd *cobra.Command, args []string) {
		printBatteryStatus()
	},
}

type BatteryInfo struct {
	Health      string  `json:"health"`
	Percentage  int     `json:"percentage"`
	Plugged     string  `json:"plugged"`
	Status      string  `json:"status"`
	Temperature float64 `json:"temperature"`
	Current     int64   `json:"current"`
}

func init() {
	rootCmd.AddCommand(batteryCmd)
	batteryCmd.AddCommand(batteryStatusCmd)
}

func printBatteryStatus() {
	// Check if termux-battery-status is available
	output, err := exec.Command("termux-battery-status").Output()
	if err != nil {
		fmt.Printf("Error: Unable to get battery status\n")
		fmt.Printf("Make sure Termux-API is installed: pkg install termux-api\n")
		if GetVerbose() {
			fmt.Printf("Error details: %v\n", err)
		}
		return
	}

	var info BatteryInfo
	if err := json.Unmarshal(output, &info); err != nil {
		fmt.Printf("Error parsing battery info: %v\n", err)
		return
	}

	fmt.Println("=== Battery Status ===")
	fmt.Printf("Percentage: %d%%\n", info.Percentage)
	fmt.Printf("Status: %s\n", info.Status)
	fmt.Printf("Health: %s\n", info.Health)
	fmt.Printf("Plugged: %s\n", info.Plugged)
	fmt.Printf("Temperature: %.1f°C\n", info.Temperature)

	// Battery indicator
	indicator := getBatteryIndicator(info.Percentage, info.Status)
	fmt.Printf("\n%s\n", indicator)

	if GetVerbose() && info.Current != -9223372036854775808 {
		fmt.Printf("\nCurrent: %d mA\n", info.Current)
	}
}

func getBatteryIndicator(percentage int, status string) string {
	var icon string
	switch {
	case percentage >= 90:
		icon = "████████████"
	case percentage >= 75:
		icon = "██████████░░"
	case percentage >= 50:
		icon = "████████░░░░"
	case percentage >= 25:
		icon = "██████░░░░░░"
	case percentage >= 10:
		icon = "████░░░░░░░░"
	default:
		icon = "██░░░░░░░░░░"
	}

	statusIcon := ""
	if status == "CHARGING" {
		statusIcon = " ⚡"
	} else if status == "FULL" {
		statusIcon = " ✓"
	}

	return fmt.Sprintf("[%s] %d%%%s", icon, percentage, statusIcon)
}
