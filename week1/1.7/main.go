package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func Unique(in io.Reader, out io.Writer) error {
	sc := bufio.NewScanner(in)
	var prev string
	for sc.Scan() {
		txt := sc.Text()

		if txt == prev {
			continue
		}

		if txt < prev {
			return fmt.Errorf("file is not sorted")
		}

		prev = txt

		fmt.Fprintln(out, txt)
	}
	return nil
}

func main() {
	err := Unique(os.Stdin, os.Stdout)
	if err != nil {
		panic(err.Error())
	}
}
