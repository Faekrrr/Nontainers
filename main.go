package main

import (
	"fmt"
	"os"
	"os/exec"
)

// Docker run <container name> cmd args
// go run main.go run cmd args
func main() {
	fmt.Println("Starting Nontainers")
	switch os.Args[1] {
	case "run":
		run()
	default:	
	panic("Unsupported command")
	}
}

func run() {
	fmt.Println("Running Nontainers")
}


func must(err error) {
	if (err != nil) { 
		panic(err)
	}
}