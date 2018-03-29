package main

import (
	"sync"
	"flag"
	"runtime"
	"os"
)

var (
	wg            sync.WaitGroup
	decrypt       bool
	secret        string
	target        string
	directoryBase string
)

func init() {
	flag.BoolVar(&decrypt, "decrypt", false, "--decrypt")
	flag.StringVar(&secret, "secret", "", "--secret=your_secret")
	flag.StringVar(&target, "target", "", "--target=your_path_target")
	flag.Parse()

	if len(target) == 0 {
		if runtime.GOOS == "windows" {
			directoryBase = "C:/"
		} else {
			directoryBase = "/home/" + os.Getenv("USER")
		}
	}
}

func main() {

	if len(secret) == 0 {
		panic("Invalid secret, try: --secret=your_secret")
	}
}
