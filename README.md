# Viscosity VPN CLI Tool

A command-line interface for controlling Viscosity VPN connections on macOS.

## Features

- **Interactive Mode**: Select VPN connections from a numbered menu when no connection name is provided
- **Direct Connection**: Connect to VPN connections by name
- **Flexible Disconnection**: Disconnect from specific connections or all connections
- **Visual Status**: View status of all VPN connections with color-coded indicators
- **Simple Syntax**: Intuitive command-line interface with helpful aliases

## Prerequisites

- macOS (uses AppleScript to communicate with Viscosity)
- [Viscosity VPN client](https://www.sparklabs.com/viscosity/) installed
- Go 1.21+ (for building from source)

## Installation

### Option 1: Install with Go (Recommended)

```bash
go install github.com/yourusername/viscosity-cli@latest
```

### Option 2: Build from source

1. Clone or download this repository
2. Navigate to the project directory
3. Build the binary:
   ```bash
   go build -o viscosity-cli main.go
   ```
4. (Optional) Move the binary to your PATH:
   ```bash
   sudo mv viscosity-cli /usr/local/bin/
   ```

#### Alias
I recommend setting an alias in your bash or zsh profile.

**For Zsh (default on macOS Catalina+):**
```bash
echo 'alias vpn="viscosity-cli"' >> ~/.zshrc
source ~/.zshrc
```

## Usage

### Connect to a VPN
```bash
# Interactive mode - shows a menu to select from available connections
viscosity-cli connect
viscosity-cli on

# Direct connection to a specific VPN
viscosity-cli connect MyVPN
viscosity-cli on MyVPN
```

### Disconnect from a VPN
```bash
# Disconnect from a specific connection
viscosity-cli disconnect MyVPN
viscosity-cli off MyVPN

# Disconnect from all connections
viscosity-cli disconnect
viscosity-cli off
```

### Check VPN status
```bash
viscosity-cli status
# or
viscosity-cli list
```

This will show all your VPN connections with status indicators:
- 🟢 Connected
- 🟡 Connecting
- 🔴 Disconnected

### Get help
```bash
viscosity-cli help
viscosity-cli -h
viscosity-cli --help
```

## Examples

```bash
# Start interactive connection selection
viscosity-cli connect

# Connect to a VPN named "Work VPN" directly
viscosity-cli connect "Work VPN"

# Check status of all connections
viscosity-cli status

# Disconnect from "Work VPN"
viscosity-cli disconnect "Work VPN"

# Disconnect from all VPNs
viscosity-cli off
```

## Notes

- Connection names are case-sensitive and must match exactly as they appear in Viscosity
- If a connection name contains spaces, wrap it in quotes
- The tool requires Viscosity to be installed and accessible via AppleScript
- You may need to grant terminal applications permission to control Viscosity when first running the tool

## Troubleshooting

### Permission Issues
If you get permission errors, you may need to:
1. Open System Preferences → Security & Privacy → Privacy → Automation
2. Allow your terminal application to control Viscosity