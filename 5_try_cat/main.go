package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

var rf = flag.Bool("n", false, "行数表示")

func main() {
	flag.Parse()
	for _, fn := range flag.Args() {
		err := readFile(fn)
		if err != nil {
			fmt.Fprintln(os.Stderr, "ファイルの読み込みに失敗しました", err)
		}
	}
}

func readFile(fn string) error {
	f, err := os.Open(fn)
	if err != nil {
		return err
	}
	scanner := bufio.NewScanner(f)
	for i := 1; scanner.Scan(); i++ {
		if *rf {
			fmt.Println(i, ": ", scanner.Text())
		} else {
			fmt.Println(scanner.Text())
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err, "エラー")
	}

	return err
}
