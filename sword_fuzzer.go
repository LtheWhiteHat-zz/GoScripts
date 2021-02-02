package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
)

func checkerr(err error) {
	if err != nil {
		panic(err)
	}
}

func fuzz(target, item string, channels chan int, wg *sync.WaitGroup, f2 *os.File) {
	defer wg.Done()
	alvo := fmt.Sprintf("%s/%s", target, item)

	res, err := http.Get(alvo)
	checkerr(err)
	status := res.StatusCode
	toWrite := fmt.Sprintf("%s -> %d \n", alvo, status)
	w := bufio.NewWriter(f2)
	_, err2 := w.WriteString(toWrite)
	checkerr(err2)
	w.Flush()
}

func main() {
	if len(os.Args) != 5 {
		log.Fatalln("Usage: ./sword_fuzzer http://target path/to/wordlist output numChannels")
	}

	target := os.Args[1]
	list := os.Args[2]

	f, err := os.Open(list)
	checkerr(err)
	defer f.Close()

	output := os.Args[3]
	f2, err := os.Create(output)
	checkerr(err)

	fmt.Printf("Target: %s\n", target)
	scanner := bufio.NewScanner(f)
	var wg sync.WaitGroup

	numChannels, err := strconv.Atoi(os.Args[4])
	checkerr(err)

	channels := make(chan int, numChannels)

	for scanner.Scan() {
		wg.Add(1)
		item := scanner.Text()
		go fuzz(target, item, channels, &wg, f2)
	}
	wg.Wait()

}
