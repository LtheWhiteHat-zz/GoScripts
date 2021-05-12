package main

import (
	"fmt"
)

func hello(codename string) string {
	reply := fmt.Sprint("Good morning Agent %s your mission will be supplied to you shortly.", codename)
	return (reply)
}
