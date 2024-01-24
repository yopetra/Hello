package main

import (
	"fmt"
	"os"
)

func main() {
	switch os.Args[1] {
	case "hash":
		hash(os.Args[2])
	case "compare":
		compare(os.Args[2], os.Args[3])
	default:
		fmt.Println("Invalid command: %v\n", os.Args[1])
	}
}

func hash(password string) {
	fmt.Printf("TODO: Hash the password %q\n", password)
}

func compare(password, hash string) {
	fmt.Printf("TODO: Compare the password %q with the hash %q\n", password, hash)
}
