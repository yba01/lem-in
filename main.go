package main

import "os"

func main() {
	if len(os.Args) != 2 {
		return
	}
	file:=os.Args[1]
}
