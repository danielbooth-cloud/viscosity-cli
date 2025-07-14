package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		return
	}

	command := os.Args[1]

	switch command {
	case "connect", "on":
		if len(os.Args) < 3 {
			// No connection name provided, start interactive mode
			connectionName := selectConnectionInteractively()
			if connectionName == "" {
				fmt.Println("No connection selected.")
				return
			}
			connectVPN(connectionName)
		} else {
			connectionName := os.Args[2]
			connectVPN(connectionName)
		}
	case "disconnect", "off":
		if len(os.Args) >= 3 {
			connectionName := os.Args[2]
			disconnectVPN(connectionName)
		} else {
			disconnectAllVPN()
		}
	case "status", "list":
		getVPNStatus()
	case "help", "-h", "--help":
		printUsage()
	default:
		fmt.Printf("Unknown command: %s\n", command)
		printUsage()
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Println("Viscosity VPN CLI Tool")
	fmt.Println("")
	fmt.Println("Usage:")
	fmt.Println("  viscosity-cli connect [connection-name]    Connect to a VPN connection (interactive mode if no name)")
	fmt.Println("  viscosity-cli on [connection-name]         Alias for connect")
	fmt.Println("  viscosity-cli disconnect [connection-name] Disconnect from VPN (all if no name specified)")
	fmt.Println("  viscosity-cli off [connection-name]        Alias for disconnect")
	fmt.Println("  viscosity-cli status                       Show status of all VPN connections")
	fmt.Println("  viscosity-cli list                         Alias for status")
	fmt.Println("  viscosity-cli help                         Show this help message")
	fmt.Println("")
	fmt.Println("Examples:")
	fmt.Println("  viscosity-cli connect                      Start interactive connection selection")
	fmt.Println("  viscosity-cli connect MyVPN                Connect to specific VPN")
	fmt.Println("  viscosity-cli on                           Start interactive connection selection")
	fmt.Println("  viscosity-cli disconnect MyVPN             Disconnect from specific VPN")
	fmt.Println("  viscosity-cli off                          Disconnect from all VPNs")
	fmt.Println("  viscosity-cli status                       Show all connections status")
}

func connectVPN(connectionName string) {
	script := fmt.Sprintf(`
		tell application "Viscosity"
			connect "%s"
		end tell
	`, connectionName)

	err := runAppleScript(script)
	if err != nil {
		fmt.Printf("Error connecting to VPN '%s': %v\n", connectionName, err)
		os.Exit(1)
	}
	fmt.Printf("Connecting to VPN: %s\n", connectionName)
}

func disconnectVPN(connectionName string) {
	script := fmt.Sprintf(`
		tell application "Viscosity"
			disconnect "%s"
		end tell
	`, connectionName)

	err := runAppleScript(script)
	if err != nil {
		fmt.Printf("Error disconnecting from VPN '%s': %v\n", connectionName, err)
		os.Exit(1)
	}
	fmt.Printf("Disconnecting from VPN: %s\n", connectionName)
}

func disconnectAllVPN() {
	script := `
		tell application "Viscosity"
			disconnectall
		end tell
	`

	err := runAppleScript(script)
	if err != nil {
		fmt.Printf("Error disconnecting from all VPNs: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Disconnecting from all VPNs")
}

func selectConnectionInteractively() string {
	connections := getVPNConnections()
	if len(connections) == 0 {
		fmt.Println("No VPN connections found in Viscosity.")
		return ""
	}

	fmt.Println("Available VPN Connections:")
	fmt.Println("==========================")
	for i, conn := range connections {
		status := "🔴"
		if conn.State == "Connected" {
			status = "🟢"
		} else if conn.State == "Connecting" {
			status = "🟡"
		}
		fmt.Printf("%d. %s %s (%s)\n", i+1, status, conn.Name, conn.State)
	}
	fmt.Println("0. Cancel")
	fmt.Print("\nSelect a connection (enter number): ")

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("Error reading input: %v\n", err)
		return ""
	}

	input = strings.TrimSpace(input)
	if input == "0" || input == "" {
		return ""
	}

	selection, err := strconv.Atoi(input)
	if err != nil || selection < 1 || selection > len(connections) {
		fmt.Println("Invalid selection.")
		return ""
	}

	return connections[selection-1].Name
}

type VPNConnection struct {
	Name  string
	State string
}

func getVPNConnections() []VPNConnection {
	script := `
		tell application "Viscosity"
			set connectionList to ""
			repeat with conn in connections
				set connName to name of conn
				set connState to state of conn
				set connectionList to connectionList & connName & "|" & connState & "\n"
			end repeat
			return connectionList
		end tell
	`

	output, err := runAppleScriptWithOutput(script)
	if err != nil {
		return []VPNConnection{}
	}

	var connections []VPNConnection
	lines := strings.Split(strings.TrimSpace(output), "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}
		parts := strings.Split(line, "|")
		if len(parts) == 2 {
			name := strings.TrimSpace(parts[0])
			state := strings.TrimSpace(parts[1])
			connections = append(connections, VPNConnection{
				Name:  name,
				State: state,
			})
		}
	}
	return connections
}

func getVPNStatus() {
	connections := getVPNConnections()
	if len(connections) == 0 {
		fmt.Println("No VPN connections found")
		return
	}

	fmt.Println("VPN Connections:")
	fmt.Println("=================")
	for _, conn := range connections {
		status := "🔴"
		if conn.State == "Connected" {
			status = "🟢"
		} else if conn.State == "Connecting" {
			status = "🟡"
		}
		fmt.Printf("%s %s (%s)\n", status, conn.Name, conn.State)
	}
}

func runAppleScript(script string) error {
	cmd := exec.Command("osascript", "-e", script)
	return cmd.Run()
}

func runAppleScriptWithOutput(script string) (string, error) {
	cmd := exec.Command("osascript", "-e", script)
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(output), nil
}
