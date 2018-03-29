package main

import (
	"sync"
	"flag"
	"runtime"
	"os"
	"fmt"
)

var (
	wg            sync.WaitGroup
	decrypt       bool
	secret        string
	directoryBase string
)

func init() {
	flag.BoolVar(&decrypt, "decrypt", false, "-decrypt")
	flag.StringVar(&secret, "secret", "", "-secret=your_secret")
	flag.Parse()

	if runtime.GOOS == "windows" {
		directoryBase = "C:/"
	} else {
		directoryBase = "/home/" + os.Getenv("USER")
	}
}

func main() {
	fmt.Println(false)
	fmt.Println(decrypt)
	fmt.Println(secret)
	fmt.Println(directoryBase)
}
