package main

import (
    "crypto/tls"
    "fmt"
    "net"
    "os"
    "strconv"
    "strings"
    "sync"
    "time"
    "io"
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
    results := make(chan []byte, 1024)
    var wg sync.WaitGroup

    for _, port := range ports {
        wg.Add(1)
        go testPort(target, port, results, &wg)
    }

    go func() {
        wg.Wait()
        close(results)
    }()

    for res := range results {
        fmt.Printf("%s\n", string(res))
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

func testPort(target string, port int, results chan<- []byte, wg *sync.WaitGroup) {
    defer wg.Done()
    address := fmt.Sprintf("%s:%d", target, port)

    // First try TLS connection
    if testTLSConnection(address, port, results) {
        return // If TLS is successful, no need to test plain TCP
    }

    testTCPConnection(address, port, results)

    
}

func testTLSConnection(address string, port int, results chan<- []byte) bool {
    
    dialer := net.Dialer{Timeout: 3 * time.Second}
    tlsConn, err := tls.DialWithDialer(&dialer, "tcp", address, &tls.Config{InsecureSkipVerify: true})
    if err != nil {
        return false
    }
    defer tlsConn.Close()


    
    probeServer(tlsConn, port, true, results) // TLS probe
    return true
}

func testTCPConnection(address string, port int, results chan<- []byte) {
    dialer := net.Dialer{Timeout: 3 * time.Second}
    conn, err := dialer.Dial("tcp", address)
    if err != nil {
        results <- []byte(fmt.Sprintf("Port %d closed (TCP)\n", port))
        return
    }
    defer conn.Close()
    probeServer(conn, port, false, results) // TCP probe
}

func probeServer(conn net.Conn, port int, isTLS bool, results chan<- []byte) {
    buffer := make([]byte, 1024)
    conn.SetReadDeadline(time.Now().Add(3 * time.Second))
    n, err := conn.Read(buffer)
    if err == nil || n > 0 {
        if isTLS {
            results <- []byte(fmt.Sprintf("Response from port %d TLS server [server-initiated] : %s\n", port, string(buffer[:n])))
        } else {
            results <- []byte(fmt.Sprintf("Response from port %d TCP server [server-initiated] : %s\n", port, string(buffer[:n])))
        }
        return
    }


    _, err = conn.Write([]byte("GET / HTTP/1.0\r\n\r\n"))
    if err != nil {
        results <- []byte(fmt.Sprintf("Error in probing HTTP server %v\n", err))
    } else {
        conn.SetReadDeadline(time.Now().Add(3 * time.Second))
        n, err = conn.Read(buffer)
        if err == nil && n > 0 {
            if isTLS {
                results <- []byte(fmt.Sprintf("Response from port %d TLS HTTPS server [client-initiated]: %s\n", port, string(buffer[:n])))
            } else {
                results <- []byte(fmt.Sprintf("Response from port %d TCP HTTP server [client-initiated]: %s\n", port, string(buffer[:n])))
            }
            return
        }
    }

    _, err = conn.Write([]byte("\r\n\r\n\r\n\r\n"))
    if err != nil {
        results <- []byte(fmt.Sprintf("Error in probing Generic server %v\n", err))
    } else {
        conn.SetReadDeadline(time.Now().Add(5 * time.Second))
        n, err = conn.Read(buffer)
        fmt.Printf("The read length is %d\n", n)
        if err == nil || err == io.EOF {
            if isTLS {
                results <- []byte(fmt.Sprintf("Response from port %d Generic TLS server [client-initiated]: %s\n", port, string(buffer[:n])))
            } else {
                results <- []byte(fmt.Sprintf("Response from port %d Generic TCP server [client-initiated]: %s\n", port, string(buffer[:n])))
            }
            return
        } else {
            fmt.Printf("The error is %v\n", err)
        }
    }

}
