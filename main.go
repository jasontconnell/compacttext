package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/jasontconnell/compacttext/compactor"
	"io"
	"os"
)

func main() {
	in := flag.String("in", "", "input file")
	l := flag.Int("l", 100, "minimum length for each new line")
	flag.Parse()

	if *l < 0 {
		fmt.Println("need positive value")
		os.Exit(1)
	}

	var r io.Reader
	if *in == "" {
		r = os.Stdin
	} else {
		var err error
		r, err = os.Open(*in)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	s := bufio.NewScanner(r)
	var text string
	for s.Scan() {
		text += s.Text() + " "
	}

	res := compactor.Process(text, *l)

	fmt.Print(res)
}
