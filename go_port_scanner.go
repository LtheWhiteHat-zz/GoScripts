package main

import (
	"fmt"
	"net"
	"sort"
	"log"
	"strconv"
	"os"
)

func worker(ports, results chan int, target string) {
	for p := range ports {
		address := fmt.Sprintf("%s:%d",target, p)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			results <- 0
			continue
		}
		conn.Close()
		results <- p

	}
}

func main() {
	inputLastPort := os.Args[2]
	lastPort, err := strconv.Atoi(inputLastPort)
	if err != nil{
		log.Fatalln(err)
	}
	inputChannels := os.Args[3]
	channels, err := strconv.Atoi(inputChannels)
	if err != nil{
		log.Fatalln(err)
	}
	ports := make(chan int, channels)
	var openports []int
	results := make(chan int)

	for i := 0; i < cap(ports); i++ {
		go worker(ports, results, os.Args[1])
	}

	go func() {
		for i := 1; 1 < lastPort; i++ {
			ports <- i
		}
	}()

	for i := 0; i < lastPort; i++ {
		port := <-results
		if port != 0 {
			openports = append(openports, port)
		}
	}
	close(ports)
	close(results)
	sort.Ints(openports)
	for _, port := range openports {
		fmt.Printf("%d open\n", port)
	}
}
