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

### Option 1: Build from source

1. Clone or download this repository
2. Navigate to the project directory
3. Build the binary:
   ```bash
   go build -o vpn main.go
   ```
4. (Optional) Move the binary to your PATH:
   ```bash
   sudo mv vpn /usr/local/bin/
   ```

### Option 2: Use the pre-built binary

If you have the `vpn` binary, you can move it directly to your PATH:
```bash
sudo mv vpn /usr/local/bin/
```

## Usage

### Connect to a VPN
```bash
# Interactive mode - shows a menu to select from available connections
vpn connect
vpn on

# Direct connection to a specific VPN
vpn connect MyVPN
vpn on MyVPN
```

### Disconnect from a VPN
```bash
# Disconnect from a specific connection
vpn disconnect MyVPN
vpn off MyVPN

# Disconnect from all connections
vpn disconnect
vpn off
```

### Check VPN status
```bash
vpn status
# or
vpn list
```

This will show all your VPN connections with status indicators:
- 🟢 Connected
- 🟡 Connecting
- 🔴 Disconnected

### Get help
```bash
vpn help
vpn -h
vpn --help
```

## Examples

```bash
# Start interactive connection selection
vpn connect

# Connect to a VPN named "Work VPN" directly
vpn connect "Work VPN"

# Check status of all connections
vpn status

# Disconnect from "Work VPN"
vpn disconnect "Work VPN"

# Disconnect from all VPNs
vpn off
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

### Connection Not Found
Make sure the connection name exactly matches what's shown in Viscosity, including capitalization and spaces.

## License

MIT License