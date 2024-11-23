package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

func GetADConnections() ([]string, error) {
	cmd := exec.Command("netstat", "-an")
	out, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	connections := strings.Split(string(out), "\n")
	return connections, nil
}

func GetADOutgoingConnections() ([]string, error) {
	cmd := exec.Command("netstat", "-an", "|", "findstr", "ESTABLISHED")
	out, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	connections := strings.Split(string(out), "\n")
	return connections, nil
}

func GetADIncomingConnections() ([]string, error) {
	cmd := exec.Command("netstat", "-an", "|", "findstr", "ESTABLISHED")
	out, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	connections := strings.Split(string(out), "\n")
	return connections, nil
}

func GetADIncomingConnectionIPs() ([]string, error) {
	connections, err := GetADIncomingConnections()
	if err != nil {
		return nil, err
	}

	var ipList []string
	for _, conn := range connections {
		if strings.Contains(conn, "ESTABLISHED") {
			parts := strings.Fields(conn)
			if len(parts) > 4 {
				ipList = append(ipList, parts[2])
			}
		}
	}

	return ipList, nil
}

func GetADUniqueIncomingConnectionIPs() ([]string, error) {
	ips, err := GetADIncomingConnectionIPs()
	if err != nil {
		return nil, err
	}

	uniqueIPs := make(map[string]bool)
	for _, ip := range ips {
		uniqueIPs[ip] = true
	}

	var uniqueIPList []string
	for ip := range uniqueIPs {
		uniqueIPList = append(uniqueIPList, ip)
	}

	return uniqueIPList, nil
}

func GetADConnectionDurations() ([]string, error) {
	cmd := exec.Command("netstat", "-an", "|", "findstr", "ESTABLISHED")
	out, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	connections := strings.Split(string(out), "\n")
	var connectionDurations []string
	for _, conn := range connections {
		if strings.Contains(conn, "ESTABLISHED") {
			connectionDurations = append(connectionDurations, fmt.Sprintf("Connection established at %s", time.Now().Format(time.RFC1123)))
		}
	}

	return connectionDurations, nil
}

func GetADTunnel() ([]string, error) {
	cmd := exec.Command("netstat", "-an")
	out, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	tunnels := strings.Split(string(out), "\n")
	var tunnelList []string
	for _, line := range tunnels {
		if strings.Contains(line, "TCP") {
			tunnelList = append(tunnelList, line)
		}
	}
	return tunnelList, nil
}

func GetADTunnelConnections() ([]string, error) {
	return GetADTunnel()
}

func GetADForwardedTunnels() ([]string, error) {
	cmd := exec.Command("netstat", "-an")
	out, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	return strings.Split(string(out), "\n"), nil
}

func GetADFileTransferTo() ([]string, error) {
	return []string{"File Transfer To: /path/to/remote/destination"}, nil
}

func GetADFileTransferFrom() ([]string, error) {
	return []string{"File Transfer From: /path/to/source"}, nil
}

func GetADKeyboardLayout() (string, error) {
	cmd := exec.Command("powershell", "-Command", "Get-WmiObject -Class Win32_Keyboard")
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(out), nil
}

func GetADFileCopy() ([]string, error) {
	return []string{"File copied: /path/to/source to /path/to/destination"}, nil
}

func GetADFilePaste() ([]string, error) {
	return []string{"File pasted: /path/to/source to /path/to/destination"}, nil
}

func main() {
	connFlag := flag.Bool("c", false, "Get AD connections")
	outgoingFlag := flag.Bool("o", false, "Get AD outgoing connections")
	incomingFlag := flag.Bool("i", false, "Get AD incoming connections")
	ipFlag := flag.Bool("ip", false, "Get AD incoming connection IPs")
	uniqueIPFlag := flag.Bool("uip", false, "Get AD unique incoming connection IPs")
	durationFlag := flag.Bool("d", false, "Get AD connection durations")
	tunnelFlag := flag.Bool("t", false, "Get AD tunnel")
	tunnelConnFlag := flag.Bool("tc", false, "Get AD tunnel connections")
	forwardedTunnelFlag := flag.Bool("ft", false, "Get AD forwarded tunnels")
	fileTransferToFlag := flag.Bool("fto", false, "Get AD file transfer to")
	fileTransferFromFlag := flag.Bool("ftf", false, "Get AD file transfer from")
	keyboardLayoutFlag := flag.Bool("kl", false, "Get AD keyboard layout")
	fileCopyFlag := flag.Bool("fc", false, "Get AD file copy")
	filePasteFlag := flag.Bool("fp", false, "Get AD file paste")

	flag.Parse()

	if len(os.Args) < 2 && !(*connFlag || *outgoingFlag || *incomingFlag || *ipFlag || *uniqueIPFlag || *durationFlag || *tunnelFlag || *tunnelConnFlag || *forwardedTunnelFlag || *fileTransferToFlag || *fileTransferFromFlag || *keyboardLayoutFlag || *fileCopyFlag || *filePasteFlag) {
		fmt.Println("Please provide a function name as an argument or use flags for shortcuts.")
		os.Exit(1)
	}

	if len(os.Args) > 1 && os.Args[1][0] != '-' {
		function := os.Args[1]
		switch function {
		case "GetADConnections":
			GetADConnections()
		case "GetADOutgoingConnections":
			GetADOutgoingConnections()
		case "GetADIncomingConnections":
			GetADIncomingConnections()
		case "GetADIncomingConnectionIPs":
			GetADIncomingConnectionIPs()
		case "GetADUniqueIncomingConnectionIPs":
			GetADUniqueIncomingConnectionIPs()
		case "GetADConnectionDurations":
			GetADConnectionDurations()
		case "GetADTunnel":
			GetADTunnel()
		case "GetADTunnelConnections":
			GetADTunnelConnections()
		case "GetADForwardedTunnels":
			GetADForwardedTunnels()
		case "GetADFileTransferTo":
			GetADFileTransferTo()
		case "GetADFileTransferFrom":
			GetADFileTransferFrom()
		case "GetADKeyboardLayout":
			GetADKeyboardLayout()
		case "GetADFileCopy":
			GetADFileCopy()
		case "GetADFilePaste":
			GetADFilePaste()
		default:
			fmt.Printf("Invalid function name: %s\n", function)
			os.Exit(1)
		}
	} else {
		if *connFlag {
			GetADConnections()
		} else if *outgoingFlag {
			GetADOutgoingConnections()
		} else if *incomingFlag {
			GetADIncomingConnections()
		} else if *ipFlag {
			GetADIncomingConnectionIPs()
		} else if *uniqueIPFlag {
			GetADUniqueIncomingConnectionIPs()
		} else if *durationFlag {
			GetADConnectionDurations()
		} else if *tunnelFlag {
			GetADTunnel()
		} else if *tunnelConnFlag {
			GetADTunnelConnections()
		} else if *forwardedTunnelFlag {
			GetADForwardedTunnels()
		} else if *fileTransferToFlag {
			GetADFileTransferTo()
		} else if *fileTransferFromFlag {
			GetADFileTransferFrom()
		} else if *keyboardLayoutFlag {
			GetADKeyboardLayout()
		} else if *fileCopyFlag {
			GetADFileCopy()
		} else if *filePasteFlag {
			GetADFilePaste()
		}
	}
}
