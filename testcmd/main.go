package main

import (
	"bufio"
	"github.com/yuya-takeyama/argf"
	"io"
	"os"
)

func main() {
	reader, err := argf.Argf()
	if err != nil {
		panic(err)
	}

	bufreader := bufio.NewReader(reader)

	for {
		b, _, err := bufreader.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}

			panic(err)
		}

		os.Stdout.Write(append(b, byte('\n')))
	}
}
