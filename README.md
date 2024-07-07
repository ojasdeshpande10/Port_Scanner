# PortSleuth



## Introduction

synprobe is a versatile TCP service fingerprinting tool designed for cybersecurity enthusiasts and network administrators. This tool combines the capabilities of a simple TCP SYN scan and service detection to provide detailed information about the services running on a host. It operates similarly to nmap -sS and nmap -sV, identifying open ports and gathering additional service information.


## Features

### 1. Port Scanning
- **TCP SYN Scan**: Quickly identify open TCP ports on the target host.
- **Custom Port Ranges**: Specify a single port or a range of ports (e.g., 80 or 20-100) to scan.
- **Default Common Ports**: If no port range is specified, synprobe scans commonly used TCP ports (21, 22, 23, 25, 80, 110, 143, 443, 587, 853, 993, 3389, 8080).

### 2. Service Fingerprinting
- **Immediate Server Responses**: For open ports, synprobe captures and prints the first 1024 bytes returned by the server.
- **Probe Requests**: If no immediate response is received, synprobe sends predefined probe requests to elicit a response.
  - **GET Request**: `GET / HTTP/1.0\r\n\r\n`
  - **Generic Lines**: `\r\n\r\n\r\n\r\n`
- **TCP and TLS Support**: Distinguish between TCP and TLS services using appropriate probes.



### Steps to run

Build the synprobe program:

go build synprobe.go

Server Command:

./synprobe -p [port range/port number] host address

Tested on Kali Linux

Go should be installed >1.16
