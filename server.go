package main

import (
	"fmt"
	"net"
	"log"
)

func main() {

	go func () {
			ln, err := net.Listen("tcp", ":7")
			if err != nil { log.Println(err) }
			for {
				conn, err := ln.Accept()
				if err != nil { log.Println(err) }
				go handleRequest(conn, "ECHO", 7)
			}
	}()

	go func () {
			ln, err := net.Listen("tcp", ":20")
			if err != nil { log.Println(err) }
			for {
				conn, err := ln.Accept()
				if err != nil { log.Println(err) }
				go handleRequest(conn, "FTP", 20)
			}
	}()

	go func () {
			ln, err := net.Listen("tcp", ":21")
			if err != nil { log.Println(err) }
			for {
				conn, err := ln.Accept()
				if err != nil { log.Println(err) }
				go handleRequest(conn, "FTP", 21)
			}
	}()

	go func () {
			ln, err := net.Listen("tcp", ":22")
			if err != nil { log.Println(err) }
			for {
				conn, err := ln.Accept()
				if err != nil { log.Println(err) }
				go handleRequest(conn, "SSH", 22)
			}
	}()

	go func () {
			ln, err := net.Listen("tcp", ":80")
			if err != nil { log.Println(err) }
			for {
				conn, err := ln.Accept()
				if err != nil { log.Println(err) }
				go handleRequest(conn, "HTTP", 80)
			}
	}()

	go func () {
			ln, err := net.Listen("tcp", ":115")
			if err != nil { log.Println(err) }
			for {
				conn, err := ln.Accept()
				if err != nil { log.Println(err) }
				go handleRequest(conn, "SFTP", 115)
			}
	}()

	go func () {
			ln, err := net.Listen("tcp", ":443")
			if err != nil { log.Println(err) }
			for {
				conn, err := ln.Accept()
				if err != nil { log.Println(err) }
				go handleRequest(conn, "HTTPS", 443)
			}
	}()

	go func () {
			ln, err := net.Listen("tcp", ":8080")
			if err != nil { log.Println(err) }
			for {
				conn, err := ln.Accept()
				if err != nil { log.Println(err) }
				go handleRequest(conn, "HTTP", 8080)
			}
	}()

	ln, err := net.Listen("tcp", ":53")
	if err != nil { log.Println(err) }
	for {
		conn, err := ln.Accept()
		if err != nil { log.Println(err) }
		go handleRequest(conn, "DNS", 53)
	}
}


func handleRequest(conn net.Conn, protocol string, port int) {
		ip_addr := conn.RemoteAddr()
		fmt.Printf("Received %s request from IP Address %s on port %d\n", protocol, ip_addr.String(), port)
		conn.Close();
	}
