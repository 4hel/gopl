package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go readLoop(conn)
	}
}

func readLoop(c net.Conn) {

	var portAddr string

	defer c.Close()
	fmt.Fprintln(c, "220 FTP Server ready.")
	s := bufio.NewScanner(c)
	for s.Scan() {
		if err := s.Err(); err != nil {
			log.Print(err)
			continue
		}
		cmd := s.Text()
		if strings.ToLower(cmd) == "close" {
			c.Close()
		} else if strings.HasPrefix(cmd, "USER") {
			doUser(cmd, c)
		} else if cmd == "PWD" {
			doPwd(c)
		} else if cmd == "LIST" {
			doList(portAddr, c)
		} else if strings.HasPrefix(cmd, "RETR") {
			doGet(cmd, portAddr, c)
		} else if strings.HasPrefix(cmd, "PORT") {
			portAddr = doPort(cmd, c)
		} else {
			fmt.Fprintln(c, cmd)
		}
	}
}

func doGet(cmd, addr string, c net.Conn) {

	filename := strings.Split(cmd, " ")[1]

	fmt.Fprintln(c, "150 Opening ASCII mode data connection for file list")
	dataConn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer dataConn.Close()

	f, err := os.Open(filename)
	handleFileError(err, c)
	io.Copy(dataConn, f)

	fmt.Fprintln(c, "226 Transfer complete.")
}

func doPort(cmd string, c net.Conn) string {
	//PORT 127,0,0,1,163,142
	byteList := strings.Split(cmd, " ")[1]
	bytes := strings.Split(byteList, ",")
	var port uint16
	byte, _ := strconv.Atoi(bytes[4]) //163
	port += uint16(byte)
	port = port << 8
	byte, _ = strconv.Atoi(bytes[5]) //142
	port += uint16(byte)
	fmt.Fprintln(c, "200 PORT command successful")
	return fmt.Sprintf("%s.%s.%s.%s:%d", bytes[0], bytes[1], bytes[2], bytes[3], port)
}

func doList(addr string, c net.Conn) {
	cwd, err := os.Getwd()
	handleFileError(err, c)
	dir, err := os.Open(cwd)
	handleFileError(err, c)
	infos, err := dir.Readdir(0)
	handleFileError(err, c)
	fmt.Fprintln(c, "150 Opening ASCII mode data connection for file list")
	dataConn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer dataConn.Close()
	for i, fi := range infos {
		if i > 0 {
			fmt.Fprint(dataConn, ",")
		}
		fmt.Fprint(dataConn, fi.Name())
	}
	fmt.Fprintln(c, "226 Transfer complete.")
}

func doUser(cmd string, c net.Conn) {
	user := strings.Split(cmd, " ")[1]
	fmt.Fprintf(c, "230 User is %s.\n", user)
}

func doPwd(c net.Conn) {
	dir, err := os.Getwd()
	handleFileError(err, c)
	fmt.Fprintf(c, "257 \"%s\" is current directory.\n", dir)
}

func handleFileError(err error, c net.Conn) {
	if err != nil {
		log.Fatal(err)
		c.Close()
	}
}
