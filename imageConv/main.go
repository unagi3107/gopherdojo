package main

import (
	"flag"
	"log"
	"os"
	"strings"
)

func main() {
	d := flag.String("d", "./", "directory flag")
	oldExt := flag.String("o", "jpg", "old extension")
	newExt := flag.String("n", "png", "new extension")
	flag.Parse()
	recDir(*d, *oldExt, *newExt)
}

func recDir(dir string, o string, n string) {
	files, err := os.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}
	filePath := dir + "/"

	for _, file := range files {
		if file.IsDir() {
			recDir(filePath+file.Name(), o, n)
		} else {
			if strings.Index(file.Name(), o) >= 0 {
				newfile := strings.Split(file.Name(), o)[0] + n
				os.Rename(filePath+file.Name(), filePath+newfile)
			}
		}
	}
	return
}
