package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

var commonPorts = []int{21, 22, 23, 25, 80, 110, 143, 443, 587, 853, 993, 3389, 8080}

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("Usage: synprobe [-p port_range] target")
		os.Exit(1)
	}

	var target string
	var ports []int
	if args[0] == "-p" {
		if len(args) < 3 {
			fmt.Println("Usage: synprobe [-p port_range] target")
			os.Exit(1)
		}
		ports = parsePortRange(args[1])
		target = args[2]
	} else {
		ports = commonPorts
		target = args[0]
	}

	for _, port := range ports {
		address := fmt.Sprintf("%s:%d", target, port)
		conn, err := net.DialTimeout("tcp", address, 3*time.Second)
		if err != nil {
			fmt.Printf("Port %d closed\n", port)
			continue
		}
		fmt.Printf("Port %d open\n", port)
		interactWithService(conn, port)
		conn.Close()
	}
}

func parsePortRange(portRange string) []int {
	if strings.Contains(portRange, "-") {
		var ports []int
		parts := strings.Split(portRange, "-")
		startPort, _ := strconv.Atoi(parts[0])
		endPort, _ := strconv.Atoi(parts[1])
		for port := startPort; port <= endPort; port++ {
			ports = append(ports, port)
		}
		return ports
	} else {
		singlePort, _ := strconv.Atoi(portRange)
		return []int{singlePort}
	}
}

func interactWithService(conn net.Conn, port int) {
	// Depending on the port, send a specific message to try to identify the service
	var message string
	switch port {
	case 21: // FTP
		message = "HELP\r\n"
	case 80: // HTTP
		message = "HEAD / HTTP/1.0\r\n\r\n"
	case 22, 23: // SSH or Telnet, no initial message, just read banner
		message = ""
	default:
		message = "HELP\r\n"
	}

	if message != "" {
		_, err := conn.Write([]byte(message))
		if err != nil {
			fmt.Printf("Failed to write to port %d\n", port)
			return
		}
	}

	// Read response
	response, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Printf("Failed to read from port %d\n", port)
	} else {
		fmt.Printf("Response from port %d: %s", port, response)
	}
}
