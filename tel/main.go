package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"time"
)

//gtel www.example.com 80

func main() {
	connTimeout := flag.Duration("timeout", time.Second*10, "timeout for establishing connection")
	flag.Parse()
	if len(flag.Args()) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s [--timeout=<duration>] <host> <port>\n", os.Args[0])
		os.Exit(1)
	}

	serverHost := flag.Arg(0)
	serverPort := flag.Arg(1)
	serverAddr := serverHost + ":" + serverPort

	tcpConn, err := net.DialTimeout("tcp", serverAddr, *connTimeout)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to connect to %s: %v\n", serverAddr, err)
		os.Exit(1)
	}
	defer tcpConn.Close()

	go handleServerOutput(tcpConn)

	if err := handleClientInput(tcpConn); err != nil {
		fmt.Fprintf(os.Stderr, "Error during input processing: %v\n", err)
		os.Exit(1)
	}
}

func handleServerOutput(conn net.Conn) {
	buffer := make([]byte, 2048)
	for {
		bytesRead, err := conn.Read(buffer)
		if err != nil {
			if err != io.EOF {
				fmt.Fprintf(os.Stderr, "Server read error: %v\n", err)
			}
			os.Exit(1)
		}
		os.Stdout.Write(buffer[:bytesRead])
	}
}

func handleClientInput(conn net.Conn) error {
	inputBuffer := make([]byte, 2048)
	for {
		count, err := os.Stdin.Read(inputBuffer)
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return fmt.Errorf("stdin read failed: %v", err)
		}
		if count == 0 {
			return nil
		}
		_, err = conn.Write(inputBuffer[:count])
		if err != nil {
			return fmt.Errorf("failed to send to server: %v", err)
		}
	}
}
